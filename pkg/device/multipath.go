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
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"

	"k8s.io/klog/v2"
)

const (
	multipathd           = "multipathd"
	dmsetupcommand       = "dmsetup"
	majorMinorPattern    = "(.*)\\((?P<Major>\\d+),\\s+(?P<Minor>\\d+)\\)"
	orphanPathsPattern   = ".*\\s+(?P<host>\\d+):(?P<channel>\\d+):(?P<target>\\d+):(?P<lun>\\d+).*orphan"
	deviceDoesNotExist   = "No such device or address"
	scsiDeviceDeletePath = "/sys/class/scsi_device/%s:%s:%s:%s/device/delete"
)

var (
	showPathsFormat  = []string{"show", "paths", "raw", "format", "%w %d %t %i %o %T %z %s %m"}
	orphanPathRegexp = regexp.MustCompile(orphanPathsPattern)
)

// getPathsCount get number of slaves for a given device
func getPathsCount(mapper string) (count int, err error) {
	// TODO: This can be achieved reading the full line processing instead of piped command
	statusCmd := fmt.Sprintf("dmsetup status --target multipath %s | awk 'BEGIN{RS=\" \";active=0}/[0-9]+:[0-9]+/{dev=1}/A/{if (dev == 1) active++; dev=0} END{ print active }'", mapper)

	outBytes, err := exec.Command("bash", "-c", statusCmd).CombinedOutput()
	out := strings.TrimSuffix(string(outBytes), "\n")
	if err != nil || out == "" || strings.Contains(out, "Command failed") {
		return 0, fmt.Errorf("error while running dmsetup show command: %s : %v", out, err)
	}
	return strconv.Atoi(out)
}

// isMultipathTimeoutError check for timeout or similar error msg
func isMultipathTimeoutError(msg string) bool {
	return strings.Contains(msg, "timeout") || strings.Contains(msg, "receiving packet")
}

// retryCleanupDevice retry for maxtries for device Cleanup
func retryCleanupDevice(dev *Device) (err error) {
	maxTries := 10
	for try := 0; try < maxTries; try++ {
		err = cleanupDevice(dev)
		if err == nil {
			return nil
		}
		time.Sleep(5 * time.Second)
	}
	return nil
}

// cleanupDevice remove the multipath devices
func cleanupDevice(dev *Device) (err error) {
	return multipathRemoveDmDevice(dev.Mapper)
}

// multipathDisableQueuing disable queueing on the multipath device
func multipathDisableQueuing(mapper string) (err error) {
	args := []string{"message", mapper, "0", "fail_if_no_path"}
	outBytes, err := exec.Command(dmsetupcommand, args...).CombinedOutput()
	if err != nil {
		return err
	}
	out := string(outBytes)
	if out != "" && strings.Contains(out, deviceDoesNotExist) {
		return fmt.Errorf("failed to disable queuing for %s. Error: %s", mapper, out)
	}
	return nil
}

// multipathRemoveDmDevice remove multipath device via dmsetup
func multipathRemoveDmDevice(mapper string) (err error) {
	if strings.HasSuffix(mapper, "mpatha") {
		klog.Warning("skipping remove mpatha which is root")
		return
	}
	_ = multipathDisableQueuing(mapper)

	args := []string{"remove", "--force", mapper}
	outBytes, err := exec.Command(dmsetupcommand, args...).CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to remove multipath map for %s, error: %v", mapper, err)
	}
	out := string(outBytes)
	if out != "" && !strings.Contains(out, "ok") && !strings.Contains(out, deviceDoesNotExist) {
		return fmt.Errorf("failed to remove device map for %s, error: %s", mapper, out)
	}
	return nil
}

// cleanupOrphanPaths find orphan paths and remove them (best effort)
func cleanupOrphanPaths() {
	// run multipathd show paths and fetch orphan maps
	outBytes, err := exec.Command(multipathd, showPathsFormat...).CombinedOutput()
	if err != nil {
		klog.Warningf("failed to run multipathd %v, err: %s", showPathsFormat, err)
		return
	}
	out := string(outBytes)
	// rc can be 0 on the below error conditions as well
	if isMultipathTimeoutError(out) {
		klog.Warningf("failed to get multipathd %v, out %s", showPathsFormat, out)
		return
	}

	listOrphanPaths := orphanPathRegexp.FindAllString(out, -1)
	for _, orphanPath := range listOrphanPaths {
		result := findStringSubmatchMap(orphanPath, orphanPathRegexp)
		deletePath := fmt.Sprintf(scsiDeviceDeletePath, result["host"], result["channel"], result["target"], result["lun"])
		if err := deleteSdDevice(deletePath); err != nil {
			// ignore errors as its a best effort to cleanup all orphan maps
			klog.Warningf("error while deleting device: %v", err)
		}
	}
}
