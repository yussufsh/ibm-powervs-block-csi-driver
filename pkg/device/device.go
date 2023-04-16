/*
Copyright 2023 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package device

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/IBM/ibm-csi-common/pkg/mountmanager"
	"k8s.io/klog/v2"
)

// Device struct
type Device struct {
	// Major    string `json:"major,omitempty"`
	// Minor    string `json:"minor,omitempty"`
	Pathname string   `json:"pathname,omitempty"`
	Mapper   string   `json:"mapper,omitempty"`
	WWID     string   `json:"wwid,omitempty"`
	WWN      string   `json:"wwn,omitempty"`
	Size     int64    `json:"size,omitempty"` // size in MiB
	Slaves   []string `json:"slaves,omitempty"`
}

// StagingDevice represents the device information that is stored in the staging area.
type StagingDevice struct {
	VolumeID         string  `json:"volume_id"`
	VolumeAccessMode string  `json:"volume_access_mode"` // block or mount
	Device           *Device `json:"device"`
	MountInfo        *Mount  `json:"mount_info,omitempty"`
}

// Mount struct
type Mount struct {
	// ID         string   `json:"id,omitempty"`
	Mountpoint string   `json:"Mountpoint,omitempty"`
	Options    []string `json:"Options,omitempty"`
	Device     *Device  `json:"device,omitempty"`
	FSType     string   `json:"fstype,omitempty"`
}

func GetDeviceFromVolume(wwn string) (*Device, error) {
	klog.V(5).Infof(">>>>> GetDeviceFromVolume for wwn %s", wwn)
	defer klog.V(5).Infof("<<<<< GetDeviceFromVolume")

	devices, err := getLinuxDmDevices(wwn, false)
	if err != nil {
		return nil, err
	}
	if len(devices) == 0 {
		return nil, fmt.Errorf("unable to find device matching wwn %s", wwn)
	}
	return devices[0], nil
}

func getLinuxDmDevices(wwn string, needActivePath bool) ([]*Device, error) {
	klog.V(5).Infof(">>>>> getLinuxDmDdevices for wwn %s", wwn)
	defer klog.V(5).Infof("<<<<< getLinuxDmDdevices")

	args := []string{"ls", "--target", "multipath"}
	out, err := exec.Command(dmsetupcommand, args...).CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve multipath devices")
	}

	var devices []*Device

	r, err := regexp.Compile(majorMinorPattern)
	if err != nil {
		return nil, fmt.Errorf("unable to compile regex with %s", majorMinorPattern)
	}
	listOut := r.FindAllString(string(out), -1)

	for _, line := range listOut {
		result := findStringSubmatchMap(line, r)
		dev := &Device{}
		dev.Pathname = "dm-" + result["Minor"]

		mapName, err := getMpathName(dev.Pathname)
		if err != nil {
			return nil, err
		}
		dev.Mapper = "/dev/mapper/" + mapName

		uuid, err := getUUID(dev.Pathname)
		if err != nil {
			return nil, err
		}
		tmpWWID := strings.TrimPrefix(uuid, "mpath-")
		tmpWWN := tmpWWID[1:] // truncate scsi-id prefix

		if strings.EqualFold(wwn, tmpWWN) {
			klog.V(5).Infof("getLinuxDmDevices FOUND for wwn %s", dev.WWN)
			dev.WWN = tmpWWN
			dev.WWID = tmpWWID
			devSize, err := getSizeOfDeviceInMiB(dev.Pathname)
			if err != nil {
				return nil, err
			}
			dev.Size = devSize
			multipathdShowPaths, err := retryGetPathOfDevice(dev, needActivePath)
			if err != nil {
				err = fmt.Errorf("unable to get scsi slaves for device:%s. Error: %v", dev.WWN, err)
				return nil, err
			}
			var slaves []string
			for _, path := range multipathdShowPaths {
				if path != nil && (!needActivePath || path.ChkState == "ready") {
					slaves = append(slaves, path.Device)
				}
			}
			dev.Slaves = slaves
			if len(dev.Slaves) > 0 {
				devices = append(devices, dev)
			}

		}
	}

	return devices, nil
}

func getMpathName(pathname string) (string, error) {
	fileName := fmt.Sprintf("/sys/block/%s/dm/name", pathname)
	return readFirstLine(fileName)
}

func getUUID(pathname string) (string, error) {
	fileName := fmt.Sprintf("/sys/block/%s/dm/uuid", pathname)
	return readFirstLine(fileName)
}

func getSizeOfDeviceInMiB(pathname string) (int64, error) {
	fileName := fmt.Sprintf("/sys/block/%s/size", pathname)
	size, err := readFirstLine(fileName)
	if err != nil {
		err = fmt.Errorf("unable to get size for device: %s Err: %v", pathname, err)
		return -1, err
	}
	sizeInSector, err := strconv.ParseInt(size, 10, 0)
	if err != nil {
		err = fmt.Errorf("unable to parse size for device: %s Err: %v", pathname, err)
		return -1, err
	}
	return sizeInSector / 2 * 1024, nil
}

// DeleteDevice : delete the multipath device
func DeleteDevice(dev *Device, m mountmanager.Mounter) (err error) {
	klog.V(5).Infof(">>>>> DeleteDevice for wwn %s", dev.WWN)
	defer klog.V(5).Infof("<<<<< DeleteDevice")

	// // perform cleanup of the multipath device
	// if dev.WWN == "" {
	// 	return fmt.Errorf("no WWN of device %v present, failing delete", dev)
	// }

	//check if device is mounted or has holders
	err = checkIfDeviceCanBeDeleted(dev, m)
	if err != nil {
		return err
	}

	err = retryTearDownMultipathDevice(dev)
	if err != nil {
		return err
	}
	return nil
}

// checkIfDeviceCanBeDeleted: check if device is currently mounted
func checkIfDeviceCanBeDeleted(dev *Device, m mountmanager.Mounter) (err error) {
	// first check if the device is already mounted
	ml, err := m.List()
	if err != nil {
		return err
	}
	for _, mount := range ml {
		if strings.EqualFold(dev.Mapper, mount.Device) || strings.EqualFold(dev.Pathname, mount.Device) {
			return fmt.Errorf("%s is currently mounted for wwn %s", dev.Mapper, dev.WWN)
		}
	}
	return nil
}

// retry the tearDown if there is failure
func retryTearDownMultipathDevice(dev *Device) error {
	try := 0
	for {
		err := tearDownMultipathDevice(dev)
		if err != nil {
			if try < maxTries {
				try++
				time.Sleep(time.Duration(try) * time.Second)
				continue
			}
			return err
		}
		return nil
	}
}

// CreateDevice : attached and creates linux devices to host
func CreateDevice(wwn string) (dev *Device, err error) {
	klog.V(5).Infof(">>>>> CreateDevice for wwn %s", wwn)
	defer klog.V(5).Infof("<<<<< CreateDevice")

	device, err := createLinuxDevice(wwn)
	if err != nil {
		klog.Errorf("unable to create device for wwn %v", wwn)
		// If we encounter an error, there may be some devices created and some not.
		// If we fail the operation , then we need to rollback all the created device
		//TODO : cleanup all the devices created
		return nil, err
	}

	return device, nil
}

// createLinuxDevice : attaches and creates a new linux device
// nolint: gocyclo
func createLinuxDevice(wwn string) (dev *Device, err error) {
	err = scsiHostRescan()
	if err != nil {
		return nil, err
	}

	// sleeping for 1 second waiting for device %s to appear after rescan
	time.Sleep(time.Second * 1)

	// find multipath devices after the rescan and login
	// Start a Countdown ticker
	var devices []*Device
	for i := 0; i <= 5; i++ {
		devices, err = getLinuxDmDevices(wwn, true)
		if err != nil {
			return nil, err
		}
		for _, d := range devices {
			// ignore devices with no paths
			if len(d.Slaves) == 0 {
				continue
			}
			// Match wwn
			if strings.EqualFold(d.WWN, wwn) {
				klog.V(5).Infof("createLinuxDevice FOUND for wwn %s and slaves %+v", d.WWN, d.Slaves)
				return d, nil
			}
		}
		// TODO:  handle remapped luns, orphan paths,
		// handle error mappers
		cleanupErrorMultipathMaps()

		// sleeping for 5 seconds waiting for device to appear after rescan
		time.Sleep(time.Second * 5)
	}

	// Reached here signifies the device was not found, throw an error
	return nil, fmt.Errorf("fc device not found for wwn %s", wwn)
}

func scsiHostRescan() error {
	scsiPath := "/sys/class/scsi_host/"
	if dirs, err := os.ReadDir(scsiPath); err == nil {
		for _, f := range dirs {
			name := scsiPath + f.Name() + "/scan"
			data := []byte("- - -")
			err := os.WriteFile(name, data, 0666)
			if err != nil {
				return fmt.Errorf("scsi host rescan failed : error: %v", err)
			}
		}
	}
	return nil
}
