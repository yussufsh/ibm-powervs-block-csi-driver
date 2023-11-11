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

package remote

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog/v2"
)

var (
	piSession *ibmpisession.IBMPISession
	piID      string
)

func (r *Remote) createPVSResources() (err error) {
	var image, network string

	piSession, err = geIBMPISession()
	if err != nil {
		return fmt.Errorf("error while creating ibm pi session: %v", err)
	}

	ic := instance.NewIBMPIInstanceClient(context.Background(), piSession, piID)
	imgc := instance.NewIBMPIImageClient(context.Background(), piSession, piID)
	nc := instance.NewIBMPINetworkClient(context.Background(), piSession, piID)
	kc := instance.NewIBMPIKeyClient(context.Background(), piSession, piID)

	if image, err = getEnvVar("POWERVS_IMAGE"); err != nil {
		image = "Linux-CentOS-8-3"
		if err = createImage(imgc, image); err != nil {
			return fmt.Errorf("error while creating ibm pi image: %v", err)
		}
	}

	if network, err = getEnvVar("POWERVS_NETWORK"); err != nil {
		network = "pub-network"
		if err = createNetwork(nc, network); err != nil {
			return fmt.Errorf("error while creating ibm pi network: %v", err)
		}
	}
	if err = createSSHKey(kc, r.resName); err != nil {
		return fmt.Errorf("error while creating ibm pi ssh key: %v", err)
	}

	r.insID, r.publicIP, err = createInstance(ic, r.resName, image, network)
	if err != nil {
		return err
	}

	return nil
}

func createImage(imgc *instance.IBMPIImageClient, image string) error {
	if _, err := imgc.Get(image); err != nil {
		// If no image then copy one
		stockImages, err := imgc.GetAllStockImages(false, false)
		if err != nil {
			return err
		}
		var stockImageID string
		for _, sI := range stockImages.Images {
			if *sI.Name == image {
				stockImageID = *sI.ImageID
			}
		}
		if stockImageID == "" {
			return fmt.Errorf("cannot find image: %s", image)
		}
		if _, err := imgc.Create(&models.CreateImage{ImageID: stockImageID}); err != nil {
			return err
		}
	}
	return nil
}

func createNetwork(nc *instance.IBMPINetworkClient, network string) error {
	if _, err := nc.Get(network); err != nil {
		// If no public network then create one
		netType := "pub-vlan"
		if _, err := nc.Create(&models.NetworkCreate{Name: network, Type: &netType}); err != nil {
			return err
		}
		_, err := waitForNetworkVLAN(network, nc)
		if err != nil {
			return err
		}
	}
	return nil
}

func createSSHKey(kc *instance.IBMPIKeyClient, sshKey string) error {
	if _, err := kc.Get(sshKey); err != nil {
		// If no ssh key then create one
		// Create SSH key pair
		runCommand("ssh-keygen", "-t", "rsa", "-f", sshDefaultKey, "-N", "")
		publicKey, err := os.ReadFile(sshDefaultKey + ".pub")
		if err != nil {
			return fmt.Errorf("error while creating and reading SSH key files: %v", err)
		}
		sshPubKey := string(publicKey)

		if _, err := kc.Create(&models.SSHKey{Name: &sshKey, SSHKey: &sshPubKey}); err != nil {
			return err
		}
	}
	return nil
}

func createInstance(ic *instance.IBMPIInstanceClient, name, image, network string) (string, string, error) {
	memory := 4.0
	processors := 0.5
	procType := "shared"
	sysType := "e980"
	storageType := "tier1"

	nets := []*models.PVMInstanceAddNetwork{{NetworkID: &network}}
	r := &models.PVMInstanceCreate{
		ImageID:     &image,
		KeyPairName: name,
		Networks:    nets,
		ServerName:  &name,
		Memory:      &memory,
		Processors:  &processors,
		ProcType:    &procType,
		SysType:     sysType,
		StorageType: storageType,
	}
	resp, err := ic.Create(r)
	if err != nil {
		return "", "", fmt.Errorf("error while creating pvm instance: %v", err)
	}

	insID := ""
	for _, in := range *resp {
		insID = *in.PvmInstanceID
	}

	if insID == "" {
		return "", "", fmt.Errorf("error while fetching pvm instance: %v", err)
	}

	in, err := waitForInstanceHealth(insID, ic)
	if err != nil {
		return "", "", fmt.Errorf("error while waiting for pvm instance status: %v", err)
	}

	publicIP := ""
	insNets := in.Networks
	for _, net := range insNets {
		publicIP = net.ExternalIP
	}

	if publicIP == "" {
		return "", "", fmt.Errorf("error while getting pvm instance public IP")
	}

	err = waitForInstanceSSH(publicIP)
	if err != nil {
		return "", "", fmt.Errorf("error while waiting for pvm instance ssh connection: %v", err)
	}

	return insID, publicIP, err
}

func waitForNetworkVLAN(netID string, nc *instance.IBMPINetworkClient) (*models.Network, error) {
	var network *models.Network
	err := wait.PollUntilContextTimeout(context.Background(), 15*time.Second, 5*time.Minute, true, func(context.Context) (bool, error) {
		var err error
		network, err = nc.Get(netID)
		if err != nil || network == nil {
			return false, err
		}
		if network.VlanID != nil {
			return true, nil
		}
		return false, nil
	})

	if err != nil || network == nil {
		return nil, fmt.Errorf("failed to get target instance status: %v", err)
	}
	return network, err
}

func waitForInstanceHealth(insID string, ic *instance.IBMPIInstanceClient) (*models.PVMInstance, error) {
	var pvm *models.PVMInstance
	err := wait.PollUntilContextTimeout(context.Background(), 30*time.Second, 45*time.Minute, true, func(context.Context) (bool, error) {
		var err error
		pvm, err = ic.Get(insID)
		if err != nil {
			return false, err
		}
		if *pvm.Status == "ERROR" {
			if pvm.Fault != nil {
				return false, fmt.Errorf("failed to create the lpar: %s", pvm.Fault.Message)
			}
			return false, fmt.Errorf("failed to create the lpar")
		}
		// Check for `instanceReadyStatus` health status and also the final health status "OK"
		if *pvm.Status == "ACTIVE" && (pvm.Health.Status == "WARNING" || pvm.Health.Status == "OK") {
			return true, nil
		}

		return false, nil
	})

	if err != nil || pvm == nil {
		return nil, fmt.Errorf("failed to get target instance status: %v", err)
	}
	return pvm, err
}
func waitForInstanceSSH(publicIP string) error {
	err := wait.PollUntilContextTimeout(context.Background(), 30*time.Second, 30*time.Minute, true, func(context.Context) (bool, error) {
		var err error
		_, err = runRemoteCommand(publicIP, "hostname")
		if err != nil {
			return false, nil
		}
		return true, nil
	})

	if err != nil {
		return fmt.Errorf("failed to get ssh connection: %v", err)
	}

	return err
}

func (r *Remote) destroyPVSResources() {
	ic := instance.NewIBMPIInstanceClient(context.Background(), piSession, piID)
	kc := instance.NewIBMPIKeyClient(context.Background(), piSession, piID)

	err := kc.Delete(r.resName)
	if err != nil {
		klog.Warningf("failed to destroy pvs resources, might leave behind some stale instances: %v", err)
	}

	err = ic.Delete(r.resName)
	if err != nil {
		klog.Warningf("failed to destroy pvs resources, might leave behind some stale ssh keys: %v", err)
	}

	return
}
