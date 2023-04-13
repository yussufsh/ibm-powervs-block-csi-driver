/*
Copyright 2018 The Kubernetes Authors.

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

package fibrechannel

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog/v2"

	"errors"
	"path"
	"path/filepath"
	"strings"
)

type ioHandler interface {
	ReadDir(dirname string) ([]os.DirEntry, error)
	Lstat(name string) (os.FileInfo, error)
	EvalSymlinks(path string) (string, error)
	WriteFile(filename string, data []byte, perm os.FileMode) error
}

// Connector provides a struct to hold all of the needed parameters to make our Fibre Channel connection
type Connector struct {
	VolumeName string
	TargetWWNs []string
	Lun        string
	WWIDs      []string
}

// OSioHandler is a wrapper that includes all the necessary io functions used for (Should be used as default io handler)
type OSioHandler struct{}

// ReadDir calls the ReadDir function from ioutil package
func (handler *OSioHandler) ReadDir(dirname string) ([]os.DirEntry, error) {
	return os.ReadDir(dirname)
}

// Lstat calls the Lstat function from os package
func (handler *OSioHandler) Lstat(name string) (os.FileInfo, error) {
	return os.Lstat(name)
}

// EvalSymlinks calls EvalSymlinks from filepath package
func (handler *OSioHandler) EvalSymlinks(path string) (string, error) {
	return filepath.EvalSymlinks(path)
}

// WriteFile calls WriteFile from ioutil package
func (handler *OSioHandler) WriteFile(filename string, data []byte, perm os.FileMode) error {
	return os.WriteFile(filename, data, perm)
}

// FindMultipathDeviceForDevice given a device name like /dev/sdx, find the devicemapper parent
func FindMultipathDeviceForDevice(device string, io ioHandler, vol string) (string, error) {
	disk, err := findDeviceForPath(device, io)
	if err != nil {
		return "", err
	}

	if strings.HasPrefix(disk, "dm-") {
		klog.Infof("fb: FindMultipathDeviceForDevice from vol: %s device: %s dstPath: %s itself is dm", vol, device, disk)
		return "/dev/" + disk, nil
	}

	sysPath := "/sys/block/"
	if dirs, err2 := io.ReadDir(sysPath); err2 == nil {
		for _, f := range dirs {
			name := f.Name()
			if strings.HasPrefix(name, "dm-") {
				if _, err1 := io.Lstat(sysPath + name + "/slaves/" + disk); err1 == nil {
					return "/dev/" + name, nil
				}
			}
		}
	} else {
		return "", err2
	}

	return "", nil
}

// findDeviceForPath Find the underlaying disk for a linked path such as /dev/disk/by-path/XXXX or /dev/mapper/XXXX
// will return sdX or hdX etc, if /dev/sdX is passed in then sdX will be returned
func findDeviceForPath(path string, io ioHandler) (string, error) {
	devicePath, err := io.EvalSymlinks(path)
	if err != nil {
		return "", err
	}
	// if path /dev/hdX split into "", "dev", "hdX" then we will
	// return just the last part
	parts := strings.Split(devicePath, "/")
	if len(parts) == 3 && strings.HasPrefix(parts[1], "dev") {
		return parts[2], nil
	}
	return "", errors.New("Illegal path for device " + devicePath)
}

func scsiHostRescan(io ioHandler) error {
	scsiPath := "/sys/class/scsi_host/"
	if dirs, err := io.ReadDir(scsiPath); err == nil {
		for _, f := range dirs {
			name := scsiPath + f.Name() + "/scan"
			data := []byte("- - -")
			err := io.WriteFile(name, data, 0666)
			if err != nil {
				return fmt.Errorf("scsi host rescan failed : error: %v", err)
			}
		}
	}
	return nil
}

//func fcHostIssueLip(io ioHandler) {
//	fcPath := "/sys/class/fc_host/"
//	if dirs, err := io.ReadDir(fcPath); err == nil {
//		for _, f := range dirs {
//			name := fcPath + f.Name() + "/issue_lip"
//			data := []byte("1")
//			io.WriteFile(name, data, 0666)
//		}
//	}
//}

func searchDisk(c Connector, io ioHandler) (string, error) {
	var diskIds []string
	var disk string
	var dm string

	if len(c.TargetWWNs) != 0 {
		diskIds = c.TargetWWNs
		klog.Infof("[DEBUG] TargetWWNs with diskIds: %+v", diskIds)
	} else {
		diskIds = c.WWIDs
		klog.Infof("[DEBUG] TargetWWID with diskIds: %+v", diskIds)
	}

	rescaned := false
	// two-phase search:
	// first phase, search existing device path, if a multipath dm is found, exit loop
	// otherwise, in second phase, rescan scsi bus and search again, return with any findings
	for _, diskID := range diskIds {
		klog.Infof("fb: starting to rescan disk with polling %s", c.VolumeName)
		err := wait.PollImmediate(2*time.Second, 2*time.Minute, func() (bool, error) {
			if !rescaned {
				err := scsiHostRescan(io)
				if err != nil {
					return false, err
				}
				rescaned = true
			}

			if len(c.TargetWWNs) != 0 {
				disk, dm = findDisk(diskID, c.Lun, io)
			} else {
				disk, dm = findDiskWWIDs(diskID, c.VolumeName, io)
			}

			// found the disk before timeout
			if dm != "" {
				klog.Infof("fb: rescan disk with polling vol: %s disk: %s dm: %s", c.VolumeName, disk, dm)
				return true, nil
			} else if disk != "" {
				// FIXME: this is only to test dm path; so we are not returning
				klog.Infof("fb: rescan disk with polling vol: %s disk: %s", c.VolumeName, disk)
			} else {
				// if no disk matches then retry
				klog.Infof("fb: rescan disk with polling vol: %s [EMPTY]", c.VolumeName)
			}

			// wait until timeout
			return false, nil
		})
		if err != nil {
			klog.Errorf("failed within timeout vol: %s err: %v", c.VolumeName, err)
			return "", fmt.Errorf("failed within timeout vol: %s err: %v", c.VolumeName, err)
		}
		if err == nil && disk == "" && dm == "" {
			klog.Errorf("failed within timeout: no fc disk found for vol: %s", c.VolumeName)
		}
		// if multipath device is found, break
		if dm != "" {
			return dm, nil
		}

	}

	// FIXME: this is only to test dm path; so we are not returning for disk only
	return "", fmt.Errorf("no fc disk found : no multipath disk found for vol: %v", c.VolumeName)
}

// given a wwn and lun, find the device and associated devicemapper parent
func findDisk(wwn, lun string, io ioHandler) (string, string) {
	FcPath := "-fc-0x" + wwn + "-lun-" + lun
	DevPath := "/dev/disk/by-path/"
	if dirs, err := io.ReadDir(DevPath); err == nil {
		for _, f := range dirs {
			name := f.Name()
			if strings.Contains(name, FcPath) {
				if disk, err1 := io.EvalSymlinks(DevPath + name); err1 == nil {
					if dm, err2 := FindMultipathDeviceForDevice(disk, io, DevPath+name); err2 == nil {
						return disk, dm
					}
				}
			}
		}
	}
	return "", ""
}

// given a wwid, find the device and associated devicemapper parent
func findDiskWWIDs(wwid, vol string, io ioHandler) (string, string) {
	// Example wwid format:
	//   3600508b400105e210000900000490000
	//   <VENDOR NAME> <IDENTIFIER NUMBER>
	// Example of symlink under by-id:
	//   /dev/by-id/scsi-3600508b400105e210000900000490000
	//   /dev/by-id/scsi-<VENDOR NAME>_<IDENTIFIER NUMBER>
	// The wwid could contain white space and it will be replaced
	// underscore when wwid is exposed under /dev/by-id.

	FcPath := "scsi-" + wwid
	DevID := "/dev/disk/by-id/"
	if dirs, err := io.ReadDir(DevID); err == nil {
		for _, f := range dirs {
			name := f.Name()
			if name == FcPath {
				disk, err := io.EvalSymlinks(DevID + name)
				if err != nil {
					klog.Errorf("fc: failed to find a corresponding disk from vol %s, symlink[%s], error %v", vol, DevID+name, err)
					return "", ""
				}
				dm, err1 := FindMultipathDeviceForDevice(disk, io, DevID+name)
				if err1 != nil {
					klog.Errorf("fc: failed to find a multipath disk from vol %s, symlink[%s], for %s, error %v", vol, DevID+name, disk, err)
					return disk, ""
				}
				return disk, dm
			}
		}
	}
	klog.Errorf("fc: failed to find a disk [%s]", DevID+FcPath)
	return "", ""
}

// Attach attempts to attach a fc volume to a node using the provided Connector info
func Attach(c Connector, io ioHandler) (string, error) {
	if io == nil {
		io = &OSioHandler{}
	}

	klog.Infof("Attaching fibre channel volume")
	devicePath, err := searchDisk(c, io)

	if err != nil {
		klog.Infof("unable to find disk given WWNN or WWIDs")
		return "", err
	}

	return devicePath, nil
}

// Detach performs a detach operation on a volume
func Detach(devicePath string, io ioHandler) error {
	if io == nil {
		io = &OSioHandler{}
	}

	klog.Infof("Detaching fibre channel volume")
	var devices []string
	dstPath, err := io.EvalSymlinks(devicePath)

	if err != nil {
		return err
	}

	if strings.HasPrefix(dstPath, "/dev/dm-") {
		devices = FindSlaveDevicesOnMultipath(dstPath, io)
	} else {
		// Add single devicepath to devices
		devices = append(devices, dstPath)
	}

	klog.Infof("fc: DetachDisk devicePath: %v, dstPath: %v, devices: %v", devicePath, dstPath, devices)

	var lastErr error

	for _, device := range devices {
		err := detachFCDisk(device, io)
		if err != nil {
			klog.Errorf("fc: detachFCDisk failed. device: %v err: %v", device, err)
			lastErr = fmt.Errorf("fc: detachFCDisk failed. device: %v err: %v", device, err)
		}
	}

	if lastErr != nil {
		klog.Errorf("fc: last error occurred during detach disk:\n%v", lastErr)
		return lastErr
	}

	return nil
}

// FindSlaveDevicesOnMultipath returns all slaves on the multipath device given the device path
func FindSlaveDevicesOnMultipath(dm string, io ioHandler) []string {
	var devices []string
	// Split path /dev/dm-1 into "", "dev", "dm-1"
	parts := strings.Split(dm, "/")
	if len(parts) != 3 || !strings.HasPrefix(parts[1], "dev") {
		return devices
	}
	disk := parts[2]
	slavesPath := path.Join("/sys/block/", disk, "/slaves/")
	if files, err := io.ReadDir(slavesPath); err == nil {
		for _, f := range files {
			devices = append(devices, path.Join("/dev/", f.Name()))
		}
	}
	return devices
}

// detachFCDisk removes scsi device file such as /dev/sdX from the node.
func detachFCDisk(devicePath string, io ioHandler) error {
	// Remove scsi device from the node.
	if !strings.HasPrefix(devicePath, "/dev/") {
		return fmt.Errorf("fc detach disk: invalid device name: %s", devicePath)
	}
	// flushDevice(devicePath)
	arr := strings.Split(devicePath, "/")
	dev := arr[len(arr)-1]
	err := removeFromScsiSubsystem(dev, io)
	return err
}

// Flushes any outstanding I/O to the device
func flushDevice(deviceName string) {
	out, err := exec.Command("blockdev", "--flushbufs", deviceName).CombinedOutput()
	if err != nil {
		// Ignore the error and continue deleting the device. There is will be no retry on error.
		klog.Warningf("Failed to flush device %s: %s\n%s", deviceName, err, string(out))
	}
	klog.V(4).Infof("Flushed device %s", deviceName)
}

// Removes a scsi device based upon /dev/sdX name
func removeFromScsiSubsystem(deviceName string, io ioHandler) error {
	fileName := "/sys/block/" + deviceName + "/device/delete"
	klog.Infof("fc: remove device from scsi-subsystem: path: %s", fileName)
	data := []byte("1")
	err := io.WriteFile(fileName, data, 0666)
	if err != nil {
		return fmt.Errorf("failed remove from scsi subsystem : error: %v", err)
	}
	return nil
}

func RemoveMultipathDevice(device, volumeID string) error {
	cmd := exec.Command("dmsetup", "remove", "-f", device)
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed remove multipath device for vol %s: %s err: %v", volumeID, device, err)
	}
	klog.Infof("output of multipath device remove commandfor vol %s: %s: %s", volumeID, device, stdoutStderr)
	return nil
}
