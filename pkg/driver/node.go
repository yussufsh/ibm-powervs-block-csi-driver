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
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/IBM/ibm-csi-common/pkg/mountmanager"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/klog/v2"
	"k8s.io/mount-utils"
	mountutils "k8s.io/mount-utils"
	"sigs.k8s.io/ibm-powervs-block-csi-driver/pkg/cloud"
	"sigs.k8s.io/ibm-powervs-block-csi-driver/pkg/fibrechannel"
	"sigs.k8s.io/ibm-powervs-block-csi-driver/pkg/util"
)

const (

	// FSTypeExt2 represents the ext2 filesystem type
	FSTypeExt2 = "ext2"
	// FSTypeExt3 represents the ext3 filesystem type
	FSTypeExt3 = "ext3"
	// FSTypeExt4 represents the ext4 filesystem type
	FSTypeExt4 = "ext4"
	// FSTypeXfs represents te xfs filesystem type
	FSTypeXfs = "xfs"
	// default file system type to be used when it is not provided
	defaultFsType = "ext4"

	// defaultMaxVolumesPerInstance is the limit of volumes can be attached in the PowerVS environment
	// TODO: rightnow 99 is just a placeholder, this needs to be changed post discussion with PowerVS team
	defaultMaxVolumesPerInstance = 127 - 1
)

var (
	// nodeCaps represents the capability of node service.
	nodeCaps = []csi.NodeServiceCapability_RPC_Type{
		csi.NodeServiceCapability_RPC_STAGE_UNSTAGE_VOLUME,
		csi.NodeServiceCapability_RPC_EXPAND_VOLUME,
		csi.NodeServiceCapability_RPC_GET_VOLUME_STATS,
	}
)

// nodeService represents the node service of CSI driver
type nodeService struct {
	cloud         cloud.Cloud
	mounter       mountmanager.Mounter
	driverOptions *Options
	pvmInstanceId string
	volumeLocks   *util.VolumeLocks
	stats         StatsUtils
	// TODO: Instead of map use file storage
	blockVolumeDevices map[string]string
}

// newNodeService creates a new node service
// it panics if failed to create the service
func newNodeService(driverOptions *Options) nodeService {
	klog.V(4).Infof("retrieving node info from metadata service")
	metadata, err := cloud.NewMetadataService(cloud.DefaultKubernetesAPIClient, driverOptions.kubeconfig)
	if err != nil {
		panic(err)
	}

	pvsCloud, err := NewPowerVSCloudFunc(metadata.GetCloudInstanceId(), metadata.GetZone(), driverOptions.debug)
	if err != nil {
		panic(err)
	}

	return nodeService{
		cloud:              pvsCloud,
		mounter:            mountmanager.NewNodeMounter(),
		driverOptions:      driverOptions,
		pvmInstanceId:      metadata.GetPvmInstanceId(),
		volumeLocks:        util.NewVolumeLocks(),
		stats:              &VolumeStatUtils{},
		blockVolumeDevices: make(map[string]string),
	}
}

func (d *nodeService) NodeStageVolume(ctx context.Context, req *csi.NodeStageVolumeRequest) (*csi.NodeStageVolumeResponse, error) {
	klog.V(4).Infof("NodeStageVolume: called with args %+v", *req)

	volumeID := req.GetVolumeId()
	if len(volumeID) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Volume ID not provided")
	}

	target := req.GetStagingTargetPath()
	if len(target) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Staging target not provided")
	}

	volCap := req.GetVolumeCapability()
	if volCap == nil {
		return nil, status.Error(codes.InvalidArgument, "Volume capability not provided")
	}

	if acquired := d.volumeLocks.TryAcquire(volumeID); !acquired {
		return nil, status.Errorf(codes.Aborted, util.VolumeOperationAlreadyExistsFmt, volumeID)
	}
	defer d.volumeLocks.Release(volumeID)

	if !isValidVolumeCapabilities([]*csi.VolumeCapability{volCap}) {
		return nil, status.Error(codes.InvalidArgument, "Volume capability not supported")
	}

	// If the access type is block, do nothing for stage
	switch volCap.GetAccessType().(type) {
	case *csi.VolumeCapability_Block:
		return &csi.NodeStageVolumeResponse{}, nil
	}

	mounted, err := d.isDirMounted(target)
	if mounted {
		// Already mounted
		klog.V(4).Infof("NodeStageVolume succeeded on volume %v to staging target path %s, mount already exists.", volumeID, target)
		return &csi.NodeStageVolumeResponse{}, nil
	}

	if err != nil {
		if os.IsNotExist(err) {
			klog.V(4).Infof("NodeStageVolume attempting mkdir for path %s", target)
			if err := os.MkdirAll(target, 0750); err != nil {
				return nil, fmt.Errorf("mkdir failed for path %s (%v)", target, err)
			}
		} else {
			return nil, err
		}
	}

	mnt := volCap.GetMount()
	if mnt == nil {
		return nil, status.Error(codes.InvalidArgument, "NodeStageVolume: mnt is nil within volume capability")
	}

	fsType := mnt.GetFsType()
	if len(fsType) == 0 {
		fsType = defaultFsType
	}

	var mountOptions []string
	for _, f := range mnt.MountFlags {
		if !hasMountOption(mountOptions, f) {
			mountOptions = append(mountOptions, f)
		}
	}

	wwn, ok := req.PublishContext[WWNKey]
	if !ok || wwn == "" {
		return nil, status.Error(codes.InvalidArgument, "WWN ID is not provided or empty")
	}

	source, err := d.getDevicePath(wwn, volumeID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to find device path %s. %v", volumeID, err)
	}

	klog.V(4).Infof("NodeStageVolume: find device path for vol %s -> %s", volumeID, source)

	// Check if a device is mounted in target directory
	device, _, err := mountutils.GetDeviceNameFromMount(d.mounter, target)
	if err != nil {
		msg := fmt.Sprintf("failed to check if volume is already mounted: %v for vol %s", err, volumeID)
		return nil, status.Error(codes.Internal, msg)
	}

	// This operation (NodeStageVolume) MUST be idempotent.
	// If the volume corresponding to the volume_id is already staged to the staging_target_path,
	// and is identical to the specified volume_capability the Plugin MUST reply 0 OK.
	sourcePath, err := filepath.EvalSymlinks(source)
	if err == nil && device == sourcePath {
		klog.V(4).Infof("NodeStageVolume: volume=%q already staged", volumeID)
		return &csi.NodeStageVolumeResponse{}, nil
	}

	// FormatAndMount will format only if needed
	klog.V(5).Infof("NodeStageVolume: starting formatting %s and mounting at %s with fstype %s for vol %s", source, target, fsType, volumeID)
	err = d.mounter.GetSafeFormatAndMount().FormatAndMount(source, target, fsType, mountOptions)
	if err != nil {
		msg := fmt.Sprintf("could not format %q and mnt it at %q for vol %s with err %v", source, target, volumeID, err)
		return nil, status.Error(codes.Internal, msg)
	}
	klog.V(5).Infof("NodeStageVolume: completed formatting %s and mounting at %s with fstype %s for vol %s", source, target, fsType, volumeID)

	// FIX: storage the block devices in a map to clean up during UnpublishVolume.
	klog.V(5).Infof("NodeStageVolume: storing device %s into blockVolumeDevices for vol %s", source, volumeID)
	d.blockVolumeDevices[volumeID] = sourcePath

	return &csi.NodeStageVolumeResponse{}, nil
}

func (d *nodeService) NodeUnstageVolume(ctx context.Context, req *csi.NodeUnstageVolumeRequest) (*csi.NodeUnstageVolumeResponse, error) {
	klog.V(4).Infof("NodeUnstageVolume: called with args %+v", *req)
	volumeID := req.GetVolumeId()
	if len(volumeID) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Volume ID not provided")
	}

	if acquired := d.volumeLocks.TryAcquire(volumeID); !acquired {
		return nil, status.Errorf(codes.Aborted, util.VolumeOperationAlreadyExistsFmt, volumeID)
	}
	defer d.volumeLocks.Release(volumeID)

	target := req.GetStagingTargetPath()
	if len(target) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Staging target not provided")
	}

	// Check if target directory is a mount point. GetDeviceNameFromMount
	// given a mnt point, finds the device from /proc/mounts
	// returns the device name, reference count, and error code
	dev, refCount, err := mountutils.GetDeviceNameFromMount(d.mounter, target)
	if err != nil {
		msg := fmt.Sprintf("failed to check if volume is mounted: %v for vol %s", err, volumeID)
		return nil, status.Error(codes.Internal, msg)
	}

	// From the spec: If the volume corresponding to the volume_id
	// is not staged to the staging_target_path, the Plugin MUST
	// reply 0 OK.
	if refCount == 0 {
		klog.V(5).Infof("NodeUnstageVolume: %s target not mounted", target)

		// Fix: clean up block device
		devicePath := d.lookupForBlockVolumeDevice(volumeID)
		if devicePath != "" {
			err = cleanupDevice(devicePath, volumeID)
		}

		// only if cleanupDevice is success remove from cache
		d.clearBlockVolumeDevice(volumeID)

		if err != nil {
			return nil, err
		}

		return &csi.NodeUnstageVolumeResponse{}, nil
	}

	if refCount > 1 {
		klog.Warningf("NodeUnstageVolume: found %d references to device %s mounted at target path %s for vol %s", refCount, dev, target, volumeID)
	}

	klog.V(5).Infof("NodeUnstageVolume: starting unmounting %s for vol %s", target, volumeID)
	err = mountutils.CleanupMountPoint(target, d.mounter, false)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Could not unmount for vol %s target %q: %v", volumeID, target, err)
	}
	klog.V(5).Infof("NodeUnstageVolume: completed unmounting %s for vol %s", target, volumeID)

	deviceFromMap := d.lookupForBlockVolumeDevice(volumeID)
	if deviceFromMap == dev {
		klog.V(5).Infof("found xfs device in map")
	} else {
		klog.V(5).Infof("did not find xfs device in map")
	}
	err = cleanupDevice(dev, volumeID)
	if err != nil {
		return nil, err
	}
	// only if cleanupDevice is success remove from cache
	d.clearBlockVolumeDevice(volumeID)

	return &csi.NodeUnstageVolumeResponse{}, nil
}

func (d *nodeService) NodeExpandVolume(ctx context.Context, req *csi.NodeExpandVolumeRequest) (*csi.NodeExpandVolumeResponse, error) {
	klog.V(4).Infof("NodeExpandVolume: called with args %+v", *req)
	volumeID := req.GetVolumeId()
	if len(volumeID) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Volume ID not provided")
	}
	volumePath := req.GetVolumePath()
	if len(volumePath) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Volume Path not provided")
	}

	if acquired := d.volumeLocks.TryAcquire(volumeID); !acquired {
		return nil, status.Errorf(codes.Aborted, util.VolumeOperationAlreadyExistsFmt, volumeID)
	}
	defer d.volumeLocks.Release(volumeID)

	volumeCapability := req.GetVolumeCapability()
	isBlock := false

	// VolumeCapability is optional, if specified, use that as source of truth.
	if volumeCapability != nil {
		if !isValidVolumeCapabilities([]*csi.VolumeCapability{volumeCapability}) {
			return nil, status.Error(codes.InvalidArgument, "Volume capability not supported")
		}
		isBlock = volumeCapability.GetBlock() != nil
	} else {
		// VolumeCapability is nil, check if volumePath points to a block device.
		var err error
		isBlock, err = d.stats.IsBlockDevice(volumePath)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to determine if volumePath %v is a block device: %v", volumePath, err)
		}
	}
	// Noop for block NodeExpandVolume.
	if isBlock {
		klog.V(4).InfoS("NodeExpandVolume: called, since given volumePath is a block device, ignoring...", "volumeID", volumeID, "volumePath", volumePath)
		return &csi.NodeExpandVolumeResponse{}, nil
	}

	notMounted, err := d.mounter.IsLikelyNotMountPoint(volumePath)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "volume path %s check for mount failed: %v", volumePath, err)
	}

	if notMounted {
		return nil, status.Errorf(codes.Internal, "volume path %s is not mounted", volumePath)
	}

	devicePath, _, err := mount.GetDeviceNameFromMount(d.mounter, volumePath)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get device from volume path %s: %v", volumePath, err)
	}

	if devicePath == "" {
		return nil, status.Errorf(codes.Internal, "failed to get device from volume path %s", volumePath)
	}

	// TODO: lock per volume ID to have some idempotency
	if _, err := d.mounter.Resize(devicePath, volumePath); err != nil {
		return nil, status.Errorf(codes.Internal, "Could not resize volume %q (%q):  %v", volumeID, devicePath, err)
	}

	return &csi.NodeExpandVolumeResponse{}, nil
}

func (d *nodeService) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	klog.V(4).Infof("NodePublishVolume: called with args %+v", *req)
	volumeID := req.GetVolumeId()
	if len(volumeID) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Volume ID not provided")
	}

	source := req.GetStagingTargetPath()
	if len(source) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Staging target not provided")
	}

	target := req.GetTargetPath()
	if len(target) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Target path not provided")
	}

	// Acquire a lock on the target path instead of volumeID, since we do not want to serialize multiple node publish calls on the same volume.
	if acquired := d.volumeLocks.TryAcquire(volumeID); !acquired {
		return nil, status.Errorf(codes.Aborted, util.VolumeOperationAlreadyExistsFmt, target)
	}
	defer d.volumeLocks.Release(volumeID)

	volCap := req.GetVolumeCapability()
	if volCap == nil {
		return nil, status.Error(codes.InvalidArgument, "Volume capability not provided")
	}

	if !isValidVolumeCapabilities([]*csi.VolumeCapability{volCap}) {
		return nil, status.Error(codes.InvalidArgument, "Volume capability not supported")
	}

	// // Check if targetPath is already mounted. If it already moounted return OK
	// notMounted, err := d.mounter.IsLikelyNotMountPoint(target)
	// if err != nil && !os.IsNotExist(err) {
	// 	//Error other than PathNotExists
	// 	return nil, status.Errorf(codes.InvalidArgument, "Can not validate target mount point %s for vol %s: %v", target, volumeID, err)
	// }

	mountOptions := []string{"bind"}
	if req.GetReadonly() {
		mountOptions = append(mountOptions, "ro")
	}

	switch mode := volCap.GetAccessType().(type) {
	case *csi.VolumeCapability_Block:
		if err := d.nodePublishVolumeForBlock(req, mountOptions); err != nil {
			return nil, err
		}
	case *csi.VolumeCapability_Mount:
		if err := d.nodePublishVolumeForFileSystem(req, mountOptions, mode); err != nil {
			return nil, err
		}
	}

	return &csi.NodePublishVolumeResponse{}, nil
}

func (d *nodeService) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	klog.V(4).Infof("NodeUnpublishVolume: called with args %+v", *req)
	volumeID := req.GetVolumeId()
	if len(volumeID) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Volume ID not provided")
	}

	target := req.GetTargetPath()
	if len(target) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Target path not provided")
	}

	// Acquire a lock on the target path instead of volumeID, since we do not want to serialize multiple node publish calls on the same volume.
	if acquired := d.volumeLocks.TryAcquire(volumeID); !acquired {
		return nil, status.Errorf(codes.Aborted, util.VolumeOperationAlreadyExistsFmt, target)
	}
	defer d.volumeLocks.Release(volumeID)

	// Check if target directory is a mount point
	// returns the device name, reference count, and error code
	// dev, _, err := d.mounter.GetDeviceName(target)
	// if err != nil {
	// 	msg := fmt.Sprintf("failed to check if volume is mounted: %v", err)
	// 	return nil, status.Error(codes.Internal, msg)
	// }
	// klog.V(5).Infof("NodeUnpublishVolume device name is: %s", dev)

	// isBlock, err := hostutil.NewHostUtil().PathIsDevice(target)
	// if err != nil {
	// 	return nil, status.Errorf(codes.NotFound, "failed to determine device type for path [%s]: %v", target, err)
	// }

	klog.V(5).Infof("NodeUnpublishVolume: starting unmounting %s for vol %s", target, volumeID)
	err := mountutils.CleanupMountPoint(target, d.mounter, false)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Could not unmount %q for vol %s volumeID: %v", target, volumeID, err)
	}
	klog.V(5).Infof("NodeUnpublishVolume: completed unmounting %s for vol %s", target, volumeID)

	return &csi.NodeUnpublishVolumeResponse{}, nil
}

func (d *nodeService) lookupForBlockVolumeDevice(vol string) string {
	dev, ok := d.blockVolumeDevices[vol]
	if ok {
		klog.V(4).Infof("lookupForBlockVolumeDevice: found volume %s path %s in blockVolumeDevices", vol, dev)
		// delete(d.blockVolumeDevices, vol)
		return dev
	}
	klog.V(4).Infof("lookupForBlockVolumeDevice: did not find volume %s path %s in blockVolumeDevices", vol, dev)
	return dev
}
func (d *nodeService) clearBlockVolumeDevice(vol string) bool {
	dev, ok := d.blockVolumeDevices[vol]
	if ok {
		klog.V(4).Infof("clearBlockVolumeDevice: found volume %s path %s in blockVolumeDevices", vol, dev)
		delete(d.blockVolumeDevices, vol)
		return true
	}
	klog.V(4).Infof("clearBlockVolumeDevice: did not find volume %s path %s in blockVolumeDevices", vol, dev)
	return false
}

func (d *nodeService) NodeGetVolumeStats(ctx context.Context, req *csi.NodeGetVolumeStatsRequest) (*csi.NodeGetVolumeStatsResponse, error) {
	var resp *csi.NodeGetVolumeStatsResponse
	if req != nil {
		klog.V(4).Infof("NodeGetVolumeStats: called with args %+v", *req)
	}

	if req == nil || req.VolumeId == "" {
		return nil, status.Error(codes.InvalidArgument, "Volume ID not provided")
	}

	if req.VolumePath == "" {
		return nil, status.Error(codes.InvalidArgument, "VolumePath not provided")
	}

	volumePath := req.VolumePath
	// return if path does not exist
	if d.stats.IsPathNotExist(volumePath) {
		return nil, status.Error(codes.NotFound, "VolumePath not exist")
	}

	// check if volume mode is raw volume mode
	isBlock, err := d.stats.IsBlockDevice(volumePath)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Failed to check volume %s is block device or not: %v", req.VolumeId, err))
	}
	// if block device, get deviceStats
	if isBlock {
		capacity, err := d.stats.DeviceInfo(volumePath)
		if err != nil {
			return nil, status.Error(codes.Internal, fmt.Sprintf("Failed to collect block device info: %v", err))
		}

		resp = &csi.NodeGetVolumeStatsResponse{
			Usage: []*csi.VolumeUsage{
				{
					Total: capacity,
					Unit:  csi.VolumeUsage_BYTES,
				},
			},
		}

		klog.V(4).Infof("Block Device Volume stats collected: %+v\n", resp)
		return resp, nil
	}

	// else get the file system stats
	available, capacity, usage, inodes, inodesFree, inodesUsed, err := d.stats.FSInfo(volumePath)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Failed to collect FSInfo: %v", err))
	}
	resp = &csi.NodeGetVolumeStatsResponse{
		Usage: []*csi.VolumeUsage{
			{
				Available: available,
				Total:     capacity,
				Used:      usage,
				Unit:      csi.VolumeUsage_BYTES,
			},
			{
				Available: inodesFree,
				Total:     inodes,
				Used:      inodesUsed,
				Unit:      csi.VolumeUsage_INODES,
			},
		},
	}

	klog.V(4).Infof("FS Volume stats collected: %+v\n", resp)
	return resp, nil
}

func (d *nodeService) NodeGetCapabilities(ctx context.Context, req *csi.NodeGetCapabilitiesRequest) (*csi.NodeGetCapabilitiesResponse, error) {
	klog.V(4).Infof("NodeGetCapabilities: called with args %+v", *req)
	var caps []*csi.NodeServiceCapability
	for _, cap := range nodeCaps {
		c := &csi.NodeServiceCapability{
			Type: &csi.NodeServiceCapability_Rpc{
				Rpc: &csi.NodeServiceCapability_RPC{
					Type: cap,
				},
			},
		}
		caps = append(caps, c)
	}
	return &csi.NodeGetCapabilitiesResponse{Capabilities: caps}, nil
}

func (d *nodeService) NodeGetInfo(ctx context.Context, req *csi.NodeGetInfoRequest) (*csi.NodeGetInfoResponse, error) {
	klog.V(4).Infof("NodeGetInfo: called with args %+v", *req)

	in, err := d.cloud.GetPVMInstanceByID(d.pvmInstanceId)
	if err != nil {
		klog.Errorf("failed to get the instance for pvmInstanceId %s, err: %s", d.pvmInstanceId, err)
		return nil, fmt.Errorf("failed to get the instance for pvmInstanceId %s, err: %s", d.pvmInstanceId, err)
	}
	image, err := d.cloud.GetImageByID(in.ImageID)
	if err != nil {
		return nil, fmt.Errorf("failed to get the image details for %s, err: %s", in.ImageID, err)
	}

	segments := map[string]string{
		DiskTypeKey: image.DiskType,
	}

	topology := &csi.Topology{Segments: segments}

	return &csi.NodeGetInfoResponse{
		NodeId:             d.pvmInstanceId,
		MaxVolumesPerNode:  d.getVolumesLimit(),
		AccessibleTopology: topology,
	}, nil
}

func (d *nodeService) nodePublishVolumeForBlock(req *csi.NodePublishVolumeRequest, mountOptions []string) error {
	target := req.GetTargetPath()
	volumeID := req.GetVolumeId()

	wwn, exists := req.PublishContext[WWNKey]
	if !exists {
		return status.Error(codes.InvalidArgument, "WWN ID not provided")
	}
	source, err := d.getDevicePath(wwn, volumeID)
	if err != nil {
		return status.Errorf(codes.Internal, "Failed to find device path for vol %s. %v", volumeID, err)
	}

	klog.V(4).Infof("NodePublishVolume [block]: find device path for vol %s -> %s", volumeID, source)

	globalMountPath := filepath.Dir(target)

	// create the global mount path if it is missing
	// Path in the form of /var/lib/kubelet/plugins/kubernetes.io/csi/volumeDevices/publish/{volumeName}
	exists, err = d.mounter.PathExists(globalMountPath)
	if err != nil {
		return status.Errorf(codes.Internal, "Could not check if path exists %q for vol %s: %v", globalMountPath, volumeID, err)
	}

	if !exists {
		if err = d.mounter.MakeDir(globalMountPath); err != nil {
			return status.Errorf(codes.Internal, "Could not create dir %q for vol %s: %v", globalMountPath, volumeID, err)
		}
	}

	// Create the mount point as a file since bind mount device node requires it to be a file
	klog.V(5).Infof("NodePublishVolume [block]: making target file %s for vol %s", target, volumeID)
	err = d.mounter.MakeFile(target)
	if err != nil {
		if removeErr := os.Remove(target); removeErr != nil {
			return status.Errorf(codes.Internal, "Could not remove mount target %q for vol %s: %v", target, volumeID, removeErr)
		}
		return status.Errorf(codes.Internal, "Could not create file %q for vol %s: %v", target, volumeID, err)
	}

	klog.V(5).Infof("NodePublishVolume [block]: mounting %s at %s for vol %s", source, target, volumeID)
	if err := d.mounter.Mount(source, target, "", mountOptions); err != nil {
		if removeErr := os.Remove(target); removeErr != nil {
			return status.Errorf(codes.Internal, "Could not remove mount target %q for vol %s: %v", target, volumeID, removeErr)
		}
		return status.Errorf(codes.Internal, "Could not mount %q at %q for vol %s: %v", source, target, volumeID, err)
	}

	// FIX: storage the block devices in a map to clean up during UnpublishVolume.
	klog.V(5).Infof("NodePublishVolume [block]: storing device %s into blockVolumeDevices for vol %s", source, volumeID)
	d.blockVolumeDevices[volumeID] = source

	return nil
}

func (d *nodeService) nodePublishVolumeForFileSystem(req *csi.NodePublishVolumeRequest, mountOptions []string, mode *csi.VolumeCapability_Mount) error {
	target := req.GetTargetPath()
	source := req.GetStagingTargetPath()
	volumeID := req.VolumeId

	// if m := mode.Mount; m != nil {
	// 	for _, f := range m.MountFlags {
	// 		if !hasMountOption(mountOptions, f) {
	// 			mountOptions = append(mountOptions, f)
	// 		}
	// 	}
	// }

	klog.V(5).Infof("NodePublishVolume: creating dir %s for vol %s", target, volumeID)
	if err := d.mounter.MakeDir(target); err != nil {
		return status.Errorf(codes.Internal, "Could not create dir %q for vol %s: %v", target, volumeID, err)
	}

	fsType := "" // Let the fsType be derived from global mount(NodeStageVolume)
	// fsType :=  mode.Mount.GetFsType()
	// if len(fsType) == 0 {
	// 	fsType = defaultFsType
	// }

	klog.V(5).Infof("NodePublishVolume: started mounting %s at %s with option %s as fstype %s for vol %s", source, target, mountOptions, fsType, volumeID)
	if err := d.mounter.Mount(source, target, fsType, mountOptions); err != nil {
		notMnt, mntErr := d.mounter.IsLikelyNotMountPoint(target)
		if mntErr != nil {
			return status.Errorf(codes.Internal, "error when validating mount path %s: %v", target, mntErr)
		}
		if !notMnt {
			if mntErr = d.mounter.Unmount(target); mntErr != nil {
				return status.Errorf(codes.Internal, "failed to unmount path %s: %v", target, mntErr)
			}
			notMnt, mntErr = d.mounter.IsLikelyNotMountPoint(target)
			if mntErr != nil {
				return status.Errorf(codes.Internal, "error when validating mount path %s: %v", target, mntErr)
			}
			if !notMnt {
				// This is very odd, we don't expect it.  We'll try again next sync loop.
				return status.Errorf(codes.Internal, "failed to unmount path %s: %v", target, err)
			}
		}
		if removeErr := os.Remove(target); removeErr != nil {
			return status.Errorf(codes.Internal, "Could not remove mount target %q for vol %s: %v", target, volumeID, err)
		}
		return status.Errorf(codes.Internal, "Could not mount %q at %q for vol %s: %v", source, target, volumeID, err)
	}
	klog.V(5).Infof("NodePublishVolume: completed mounting %s at %s with option %s as fstype %s for vol %s", source, target, mountOptions, fsType, volumeID)

	return nil
}

// getVolumesLimit returns the limit of volumes that the node supports
func (d *nodeService) getVolumesLimit() int64 {
	if d.driverOptions.volumeAttachLimit >= 0 {
		return d.driverOptions.volumeAttachLimit
	}
	return defaultMaxVolumesPerInstance
}

// hasMountOption returns a boolean indicating whether the given
// slice already contains a mount option. This is used to prevent
// passing duplicate option to the mount command.
func hasMountOption(options []string, opt string) bool {
	for _, o := range options {
		if o == opt {
			return true
		}
	}
	return false
}

// isDirMounted checks if the path is already a mount point
func (d *nodeService) isDirMounted(target string) (bool, error) {
	// Check if mount already exists
	// TODO(msau): check why in-tree uses IsNotMountPoint
	// something related to squash and not having permissions to lstat
	notMnt, err := d.mounter.IsLikelyNotMountPoint(target)
	if err != nil {
		return false, err
	}
	if !notMnt {
		// Already mounted
		return true, nil
	}
	return false, nil
}

// cleanupDevice remove the device path
func cleanupDevice(dev, volumeID string) error {
	klog.V(5).Infof("running cleanupDevice with device for vol %s: %s", volumeID, dev)
	handler := &fibrechannel.OSioHandler{}

	klog.Infof("Detaching for volume %s: %s", volumeID, dev)
	err := fibrechannel.Detach(dev, handler, volumeID)
	if err != nil {
		return status.Errorf(codes.Internal, "failed to detach %s vol %s: %v", dev, volumeID, err)
	}

	return nil
}
func (m *nodeService) getDevicePath(wwn, volumeID string) (devicePath string, err error) {
	c := fibrechannel.Connector{}
	// Prepending the 3 which is missing in the wwn getting it from the PowerVS infra
	c.WWIDs = []string{"3" + wwn}
	c.VolumeName = volumeID
	return fibrechannel.Attach(c, &fibrechannel.OSioHandler{})
}
