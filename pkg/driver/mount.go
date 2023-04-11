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

// import (
// 	"k8s.io/mount-utils"
// 	"k8s.io/utils/exec"
// 	"sigs.k8s.io/ibm-powervs-block-csi-driver/pkg/fibrechannel"
// )

// type mountInterface = mount.Interface

// // Mounter is the interface implemented by Mounter
// type Mounter interface {
// 	mountInterface

// 	GetSafeFormatAndMount() *mount.SafeFormatAndMount
// 	MakeFile(path string) error
// 	MakeDir(path string) error
// 	PathExists(path string) (bool, error)
// 	Resize(string, string) (bool, error)
// 	GetDevicePath(wwn, volumeID string) (string, error)
// }

// type NodeMounter struct {
// 	*mount.SafeFormatAndMount
// }

// // NewNodeMounter ...
// func NewNodeMounter() Mounter {
// 	// mounter.newSafeMounter returns a SafeFormatAndMount
// 	safeMounter := newSafeMounter()
// 	return &NodeMounter{safeMounter}
// }

// // NewSafeMounter ...
// func newSafeMounter() *mount.SafeFormatAndMount {
// 	realMounter := mount.New("")
// 	realExec := exec.New()
// 	return &mount.SafeFormatAndMount{
// 		Interface: realMounter,
// 		Exec:      realExec,
// 	}
// }

// func (m *NodeMounter) GetDevicePath(wwn, volumeID string) (devicePath string, err error) {
// 	c := fibrechannel.Connector{}
// 	// Prepending the 3 which is missing in the wwn getting it from the PowerVS infra
// 	c.WWIDs = []string{"3" + wwn}
// 	c.VolumeName = volumeID
// 	return fibrechannel.Attach(c, &fibrechannel.OSioHandler{})
// }
