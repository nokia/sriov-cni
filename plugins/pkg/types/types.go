package types

import (
	sriovtypes "github.com/k8snetworkplumbingwg/sriov-cni/pkg/types"
)

// VlanTrunkProviderConfig provdes methods for provider configuration
type VlanTrunkProviderConfig interface {
	InitConfig(vlanRanges *VlanTrunkRangeData)
	ApplyConfig(conf *sriovtypes.NetConf) error
	RemoveConfig(conf *sriovtypes.NetConf) error
	GetVlanData(vlanRanges *VlanTrunkRangeData)
}

// VlanTrunkRange strores trunking range
type VlanTrunkRange struct {
	Start uint
	End   uint
}

// VlanTrunkRangeData stores an array of VlanTrunkRange
type VlanTrunkRangeData struct {
	VlanTrunkRanges []VlanTrunkRange
}
