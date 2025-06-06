/*
Copyright 2021 The Kubernetes Authors.

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

package testsuites

import (
	"context"
	"fmt"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/kubernetes/test/e2e/framework"

	"sigs.k8s.io/ibm-powervs-block-csi-driver/pkg/util"
	"sigs.k8s.io/ibm-powervs-block-csi-driver/tests/e2e/driver"

	. "github.com/onsi/ginkgo/v2"
)

// DynamicallyProvisionedResizeVolumeTest will provision required StorageClass(es), PVC(s) and Pod(s)
// Waiting for the PV provisioner to create a new PV
// Update pvc storage size
// Waiting for new PVC and PV to be ready
// And finally attach pvc to the pod and wait for pod to be ready.
type DynamicallyProvisionedResizeVolumeTest struct {
	CSIDriver driver.DynamicPVTestDriver
	Pod       PodDetails
}

func (t *DynamicallyProvisionedResizeVolumeTest) Run(client clientset.Interface, namespace *v1.Namespace) {
	volume := t.Pod.Volumes[0]
	tpvc, _ := volume.SetupDynamicPersistentVolumeClaim(client, namespace, t.CSIDriver)
	defer tpvc.Cleanup()

	pvcName := tpvc.persistentVolumeClaim.Name
	pvc, err := client.CoreV1().PersistentVolumeClaims(namespace.Name).Get(context.TODO(), pvcName, metav1.GetOptions{})
	if err != nil {
		framework.ExpectNoError(err, fmt.Sprintf("fail to resize pvc(%s): %v", pvcName, err))
	}
	By(fmt.Sprintf("Get pvc name: %v", pvc.Name))
	originalSize := pvc.Spec.Resources.Requests["storage"]
	delta := resource.Quantity{}
	delta.Set(util.GiBToBytes(1))
	originalSize.Add(delta)
	pvc.Spec.Resources.Requests["storage"] = originalSize

	By("resizing the pvc")
	updatedPvc, err := client.CoreV1().PersistentVolumeClaims(namespace.Name).Update(context.TODO(), pvc, metav1.UpdateOptions{})
	if err != nil {
		framework.ExpectNoError(err, fmt.Sprintf("fail to resize pvc(%s): %v", pvcName, err))
	}

	updatedSize := updatedPvc.Spec.Resources.Requests["storage"]
	By("checking the resizing PV result")
	resizeErr := WaitForPvToResize(client, namespace, updatedPvc.Spec.VolumeName, updatedSize, 10*time.Minute, 5*time.Second)
	framework.ExpectNoError(resizeErr)
	By("Validate volume can be attached")
	tpod := NewTestPod(client, namespace, t.Pod.Cmd)
	tpod.SetupVolume(tpvc.persistentVolumeClaim, volume.VolumeMount.NameGenerate+"1", volume.VolumeMount.MountPathGenerate+"1", volume.VolumeMount.ReadOnly)
	By("deploying the pod")
	tpod.Create()
	By("checking that the pods is running")
	tpod.WaitForSuccess()
	defer tpod.Cleanup()
}

// WaitForPvToResize waiting for pvc size to be resized to desired size.
func WaitForPvToResize(c clientset.Interface, ns *v1.Namespace, pvName string, desiredSize resource.Quantity, timeout time.Duration, interval time.Duration) error {
	By(fmt.Sprintf("Waiting up to %v for pv in namespace %q to be complete", timeout, ns.Name))
	for start := time.Now(); time.Since(start) < timeout; time.Sleep(interval) {
		newPv, _ := c.CoreV1().PersistentVolumes().Get(context.TODO(), pvName, metav1.GetOptions{})
		newPvSize := newPv.Spec.Capacity["storage"]
		if desiredSize.Equal(newPvSize) {
			By(fmt.Sprintf("Pv size is updated to %v", newPvSize.String()))
			return nil
		}
	}
	return fmt.Errorf("gave up after waiting %v for pv %q to complete resizing", timeout, pvName)
}
