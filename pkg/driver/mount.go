/*
Copyright 2019 The Kubernetes Authors.

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

package driver

import (
	"fmt"
	"os"
	goexec "os/exec"

	"k8s.io/klog/v2"
	"k8s.io/utils/exec"
	"k8s.io/utils/mount"
	"sigs.k8s.io/ibm-powervs-block-csi-driver/pkg/fibrechannel"
)

// Mounter is an interface for mount operations
type Mounter interface {
	mount.Interface
	exec.Interface
	FormatAndMount(source string, target string, fstype string, options []string) error
	GetDeviceName(mountPath string) (string, int, error)
	MakeFile(pathname string) error
	MakeDir(pathname string) error
	ExistsPath(filename string) (bool, error)
	RescanSCSIBus() error
	GetDevicePath(wwn, volumeID string) (string, error)
}

type NodeMounter struct {
	mount.SafeFormatAndMount
	exec.Interface
}

func newNodeMounter() Mounter {
	return &NodeMounter{
		mount.SafeFormatAndMount{
			Interface: mount.New(""),
			Exec:      exec.New(),
		},
		exec.New(),
	}
}

func (m *NodeMounter) RescanSCSIBus() error {
	cmd := goexec.Command("/usr/bin/rescan-scsi-bus.sh")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to rescan-scsi-bus.sh: %v", err)
	}
	klog.V(5).Infof("output of rescan-scsi-bus.sh: %s", stdoutStderr)
	return nil
}

func (m *NodeMounter) GetDevicePath(wwn, volumeID string) (devicePath string, err error) {
	c := fibrechannel.Connector{}
	// Prepending the 3 which is missing in the wwn getting it from the PowerVS infra
	c.WWIDs = []string{"3" + wwn}
	c.VolumeName = volumeID
	return fibrechannel.Attach(c, &fibrechannel.OSioHandler{})
}

func (m *NodeMounter) GetDeviceName(mountPath string) (string, int, error) {
	return mount.GetDeviceNameFromMount(m, mountPath)
}

func (m *NodeMounter) MakeFile(pathname string) error {
	f, err := os.OpenFile(pathname, os.O_CREATE, os.FileMode(0644))
	if err != nil {
		if !os.IsExist(err) {
			return err
		}
	}
	if err = f.Close(); err != nil {
		return err
	}
	return nil
}

func (m *NodeMounter) MakeDir(pathname string) error {
	err := os.MkdirAll(pathname, os.FileMode(0755))
	if err != nil {
		if !os.IsExist(err) {
			return err
		}
	}
	return nil
}

func (m *NodeMounter) ExistsPath(filename string) (bool, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}
