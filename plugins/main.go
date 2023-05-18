package main

import (
	"fmt"

	sriovtypes "github.com/k8snetworkplumbingwg/sriov-cni/pkg/types"
	"github.com/nokia/sriov-cni/plugins/pkg/factory"
	"github.com/nokia/sriov-cni/plugins/pkg/utils"
)

func ApplyConfig(conf *sriovtypes.NetConf) error {
	vlanTrunkRange, err := utils.GetVlanTrunkRange(conf.VlanTrunk)
	if err != nil {
		return fmt.Errorf("GetVlanTrunkRange Error: %q", err)
	}

	vlanTrunkProviderConfig, err := factory.GetProviderConfig(conf.DeviceID)
	if err != nil {
		return fmt.Errorf("GetProviderConfig Error: %q", err)
	}

	vlanTrunkProviderConfig.InitConfig(&vlanTrunkRange)

	if err := vlanTrunkProviderConfig.ApplyConfig(conf); err != nil {
		return fmt.Errorf("ApplyConfig Error: %q", err)
	}
	return nil
}

func RemoveConfig(conf *sriovtypes.NetConf) error {
	vlanTrunkRange, err := utils.GetVlanTrunkRange(conf.VlanTrunk)
	if err != nil {
		return fmt.Errorf("GetVlanTrunkRange Error: %q", err)
	}

	vlanTrunkProviderConfig, err := factory.GetProviderConfig(conf.DeviceID)
	if err != nil {
		return fmt.Errorf("GetProviderConfig Error: %q", err)
	}

	vlanTrunkProviderConfig.InitConfig(&vlanTrunkRange)

	if err := vlanTrunkProviderConfig.RemoveConfig(conf); err != nil {
		return fmt.Errorf("RemoveConfig Error: %q", err)
	}
	return nil
}