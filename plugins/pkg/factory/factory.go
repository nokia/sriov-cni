package factory

import (
	"fmt"

	"github.com/nokia/sriov-cni/plugins/pkg/providers"
	"github.com/nokia/sriov-cni/plugins/pkg/types"
	"github.com/nokia/sriov-cni/plugins/pkg/utils"
)

const (
	//IntelProviderID Intel vendor ID
	IntelProviderID = "0x8086"
	//MellanoxProviderID Mellanox vendor ID
	MellanoxProviderID = "0x15b3"
)

// GetProviderConfig get Config for specific NIC
func GetProviderConfig(deviceID string) (types.VlanTrunkProviderConfig, error) {
	vendor, err := utils.GetVendorID(deviceID)
	if err != nil {
		return nil, fmt.Errorf("GetVendorID Error: %q", err)
	}

	switch vendor {
	case IntelProviderID:
		return providers.NewIntelTrunkProviderConfig(), nil
	case MellanoxProviderID:
		return providers.NewMellanoxTrunkProviderConfig(), nil
	default:
		return nil, fmt.Errorf("Not supported vendor: %q", vendor)
	}

}
