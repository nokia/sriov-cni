package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/nokia/sriov-cni/plugins/pkg/types"
)

var (
	// NetDirectory sysfs net directory
	NetDirectory = "/sys/class/net"
	// SysBusPci is sysfs pci device directory
	SysBusPci = "/sys/bus/pci/devices"
	//ExecCommand used for os.exec
	execCommand = exec.Command
	// TrunkFileDirectory trunk file directoy
	TrunkFileDirectory = "/sys/class/net/%s/device/sriov/%d/trunk"
)

// GetVendorID returns ID of installed vendor
func GetVendorID(deviceID string) (string, error) {
	path := filepath.Join(SysBusPci, deviceID, "vendor")

	readVendor, err := ioutil.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("Error reading vendor file %q, %q", path, err)
	}

	vendorCode := strings.Split(string(readVendor), "\n")[0]

	return vendorCode, nil
}

// GetVlanTrunkRange creates VlanTrunkRangeData from vlanTrunkString
func GetVlanTrunkRange(vlanTrunkString string) (types.VlanTrunkRangeData, error) {
	var vlanRange = []types.VlanTrunkRange{}
	trunkingRanges := strings.Split(vlanTrunkString, ",")

	for _, r := range trunkingRanges {
		values := strings.Split(r, "-")
		v1, errconv1 := strconv.Atoi(values[0])
		v2, errconv2 := strconv.Atoi(values[len(values)-1])

		if errconv1 != nil || errconv2 != nil {
			return types.VlanTrunkRangeData{}, fmt.Errorf("Trunk range error: invalid values")
		}

		v := types.VlanTrunkRange{
			Start: uint(v1),
			End:   uint(v2),
		}

		vlanRange = append(vlanRange, v)
	}
	if err := ValidateVlanTrunkRange(vlanRange); err != nil {
		return types.VlanTrunkRangeData{}, err
	}

	vlanRanges := types.VlanTrunkRangeData{
		VlanTrunkRanges: vlanRange,
	}
	return vlanRanges, nil

}

// ValidateVlanTrunkRange checks if given vlan trunking ranges are of correct form
func ValidateVlanTrunkRange(vlanRanges []types.VlanTrunkRange) error {

	for i, r1 := range vlanRanges {
		if r1.Start > r1.End {
			return fmt.Errorf("Invalid VlanTrunk range values")
		}

		if r1.Start < 1 || r1.End > 4094 {
			return fmt.Errorf("Invalid VlanTrunk range values")
		}

		for j, r2 := range vlanRanges {
			if r1.End > r2.Start && i < j {
				return fmt.Errorf("Invalid VlanTrunk range values")
			}
		}

	}
	return nil
}

// GetInfraVlanData returns vlan ranges used by cloud infra-structure
func GetInfraVlanData() ([]string, []int, error) {
	type Member struct {
		Type string `json:"type"`
		Name string `json:"name"`
	}
	type Network struct {
		Type    string   `json:"type"`
		Name    string   `json:"name,omitempty"`
		Members []Member `json:"members,omitempty"`
		VlanID  int      `json:"vlan_id,omitempty"`
		Device  string   `json:"device,omitempty"`
	}
	type NetworkConfig struct {
		Networks []Network `json:"network_config"`
	}

	path := "/etc/os-net-config/config.json"
	byteValue, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, nil, fmt.Errorf("Error reading os-net-config json file %q, %q", path, err)
	}

	var networkConfig NetworkConfig
	json.Unmarshal(byteValue, &networkConfig)

	var infraInterfaces []string
	var infraVlans []int
	for _, network := range networkConfig.Networks {
		if network.Name == "infra-bond" {
			for _, member := range network.Members {
				if member.Type == "interface" {
					infraInterfaces = append(infraInterfaces, member.Name)
				}
			}
		}
		if network.Type == "vlan" && network.Device == "infra-bond" {
			infraVlans = append(infraVlans, network.VlanID)
		}
	}

	return infraInterfaces, infraVlans, nil
}
