// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/driver/mount.go
//
// Generated by this command:
//
//	mockgen -package mocks -destination=pkg/driver/mocks/mock_mounter.go -source ./pkg/driver/mount.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	mount "k8s.io/mount-utils"
)

// MockMounter is a mock of Mounter interface.
type MockMounter struct {
	mount.Interface
	ctrl     *gomock.Controller
	recorder *MockMounterMockRecorder
}

// MockMounterMockRecorder is the mock recorder for MockMounter.
type MockMounterMockRecorder struct {
	mock *MockMounter
}

// NewMockMounter creates a new mock instance.
func NewMockMounter(ctrl *gomock.Controller) *MockMounter {
	mock := &MockMounter{ctrl: ctrl}
	mock.recorder = &MockMounterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMounter) EXPECT() *MockMounterMockRecorder {
	return m.recorder
}

// CanSafelySkipMountPointCheck mocks base method.
func (m *MockMounter) CanSafelySkipMountPointCheck() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanSafelySkipMountPointCheck")
	ret0, _ := ret[0].(bool)
	return ret0
}

// CanSafelySkipMountPointCheck indicates an expected call of CanSafelySkipMountPointCheck.
func (mr *MockMounterMockRecorder) CanSafelySkipMountPointCheck() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanSafelySkipMountPointCheck", reflect.TypeOf((*MockMounter)(nil).CanSafelySkipMountPointCheck))
}

// ExistsPath mocks base method.
func (m *MockMounter) ExistsPath(filename string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExistsPath", filename)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExistsPath indicates an expected call of ExistsPath.
func (mr *MockMounterMockRecorder) ExistsPath(filename any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExistsPath", reflect.TypeOf((*MockMounter)(nil).ExistsPath), filename)
}

// FormatAndMount mocks base method.
func (m *MockMounter) FormatAndMount(source, target, fstype string, options []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FormatAndMount", source, target, fstype, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// FormatAndMount indicates an expected call of FormatAndMount.
func (mr *MockMounterMockRecorder) FormatAndMount(source, target, fstype, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FormatAndMount", reflect.TypeOf((*MockMounter)(nil).FormatAndMount), source, target, fstype, options)
}

// GetDeviceName mocks base method.
func (m *MockMounter) GetDeviceName(mountPath string) (string, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeviceName", mountPath)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetDeviceName indicates an expected call of GetDeviceName.
func (mr *MockMounterMockRecorder) GetDeviceName(mountPath any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeviceName", reflect.TypeOf((*MockMounter)(nil).GetDeviceName), mountPath)
}

// GetMountRefs mocks base method.
func (m *MockMounter) GetMountRefs(pathname string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMountRefs", pathname)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMountRefs indicates an expected call of GetMountRefs.
func (mr *MockMounterMockRecorder) GetMountRefs(pathname any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMountRefs", reflect.TypeOf((*MockMounter)(nil).GetMountRefs), pathname)
}

// IsLikelyNotMountPoint mocks base method.
func (m *MockMounter) IsLikelyNotMountPoint(file string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsLikelyNotMountPoint", file)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsLikelyNotMountPoint indicates an expected call of IsLikelyNotMountPoint.
func (mr *MockMounterMockRecorder) IsLikelyNotMountPoint(file any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsLikelyNotMountPoint", reflect.TypeOf((*MockMounter)(nil).IsLikelyNotMountPoint), file)
}

// IsMountPoint mocks base method.
func (m *MockMounter) IsMountPoint(file string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsMountPoint", file)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsMountPoint indicates an expected call of IsMountPoint.
func (mr *MockMounterMockRecorder) IsMountPoint(file any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsMountPoint", reflect.TypeOf((*MockMounter)(nil).IsMountPoint), file)
}

// List mocks base method.
func (m *MockMounter) List() ([]mount.MountPoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]mount.MountPoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockMounterMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockMounter)(nil).List))
}

// MakeDir mocks base method.
func (m *MockMounter) MakeDir(pathname string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeDir", pathname)
	ret0, _ := ret[0].(error)
	return ret0
}

// MakeDir indicates an expected call of MakeDir.
func (mr *MockMounterMockRecorder) MakeDir(pathname any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeDir", reflect.TypeOf((*MockMounter)(nil).MakeDir), pathname)
}

// MakeFile mocks base method.
func (m *MockMounter) MakeFile(pathname string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeFile", pathname)
	ret0, _ := ret[0].(error)
	return ret0
}

// MakeFile indicates an expected call of MakeFile.
func (mr *MockMounterMockRecorder) MakeFile(pathname any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeFile", reflect.TypeOf((*MockMounter)(nil).MakeFile), pathname)
}

// Mount mocks base method.
func (m *MockMounter) Mount(source, target, fstype string, options []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Mount", source, target, fstype, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// Mount indicates an expected call of Mount.
func (mr *MockMounterMockRecorder) Mount(source, target, fstype, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mount", reflect.TypeOf((*MockMounter)(nil).Mount), source, target, fstype, options)
}

// MountSensitive mocks base method.
func (m *MockMounter) MountSensitive(source, target, fstype string, options, sensitiveOptions []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MountSensitive", source, target, fstype, options, sensitiveOptions)
	ret0, _ := ret[0].(error)
	return ret0
}

// MountSensitive indicates an expected call of MountSensitive.
func (mr *MockMounterMockRecorder) MountSensitive(source, target, fstype, options, sensitiveOptions any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MountSensitive", reflect.TypeOf((*MockMounter)(nil).MountSensitive), source, target, fstype, options, sensitiveOptions)
}

// MountSensitiveWithoutSystemd mocks base method.
func (m *MockMounter) MountSensitiveWithoutSystemd(source, target, fstype string, options, sensitiveOptions []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MountSensitiveWithoutSystemd", source, target, fstype, options, sensitiveOptions)
	ret0, _ := ret[0].(error)
	return ret0
}

// MountSensitiveWithoutSystemd indicates an expected call of MountSensitiveWithoutSystemd.
func (mr *MockMounterMockRecorder) MountSensitiveWithoutSystemd(source, target, fstype, options, sensitiveOptions any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MountSensitiveWithoutSystemd", reflect.TypeOf((*MockMounter)(nil).MountSensitiveWithoutSystemd), source, target, fstype, options, sensitiveOptions)
}

// MountSensitiveWithoutSystemdWithMountFlags mocks base method.
func (m *MockMounter) MountSensitiveWithoutSystemdWithMountFlags(source, target, fstype string, options, sensitiveOptions, mountFlags []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MountSensitiveWithoutSystemdWithMountFlags", source, target, fstype, options, sensitiveOptions, mountFlags)
	ret0, _ := ret[0].(error)
	return ret0
}

// MountSensitiveWithoutSystemdWithMountFlags indicates an expected call of MountSensitiveWithoutSystemdWithMountFlags.
func (mr *MockMounterMockRecorder) MountSensitiveWithoutSystemdWithMountFlags(source, target, fstype, options, sensitiveOptions, mountFlags any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MountSensitiveWithoutSystemdWithMountFlags", reflect.TypeOf((*MockMounter)(nil).MountSensitiveWithoutSystemdWithMountFlags), source, target, fstype, options, sensitiveOptions, mountFlags)
}

// Unmount mocks base method.
func (m *MockMounter) Unmount(target string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unmount", target)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unmount indicates an expected call of Unmount.
func (mr *MockMounterMockRecorder) Unmount(target any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unmount", reflect.TypeOf((*MockMounter)(nil).Unmount), target)
}
