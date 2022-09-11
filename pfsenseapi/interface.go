package pfsenseapi

import (
	"context"
	"encoding/json"
)

const (
	interfaceEndpoint     = "api/v1/interface"
	interfaceVLANEndpoint = "api/v1/interface/vlan"
)

// InterfaceService provides interface API methods
type InterfaceService service

type Interface struct {
	Enable                          string `json:"enable"`
	If                              string `json:"if"`
	Descr                           string `json:"descr"`
	AliasAddress                    string `json:"alias-address"`
	AliasSubnet                     string `json:"alias-subnet"`
	Ipaddr                          string `json:"ipaddr"`
	Dhcprejectfrom                  string `json:"dhcprejectfrom"`
	AdvDhcpPtTimeout                string `json:"adv_dhcp_pt_timeout"`
	AdvDhcpPtRetry                  string `json:"adv_dhcp_pt_retry"`
	AdvDhcpPtSelectTimeout          string `json:"adv_dhcp_pt_select_timeout"`
	AdvDhcpPtReboot                 string `json:"adv_dhcp_pt_reboot"`
	AdvDhcpPtBackoffCutoff          string `json:"adv_dhcp_pt_backoff_cutoff"`
	AdvDhcpPtInitialInterval        string `json:"adv_dhcp_pt_initial_interval"`
	AdvDhcpPtValues                 string `json:"adv_dhcp_pt_values"`
	AdvDhcpSendOptions              string `json:"adv_dhcp_send_options"`
	AdvDhcpRequestOptions           string `json:"adv_dhcp_request_options"`
	AdvDhcpRequiredOptions          string `json:"adv_dhcp_required_options"`
	AdvDhcpOptionModifiers          string `json:"adv_dhcp_option_modifiers"`
	AdvDhcpConfigAdvanced           string `json:"adv_dhcp_config_advanced"`
	AdvDhcpConfigFileOverride       string `json:"adv_dhcp_config_file_override"`
	AdvDhcpConfigFileOverridePath   string `json:"adv_dhcp_config_file_override_path"`
	Ipaddrv6                        string `json:"ipaddrv6"`
	Dhcp6Duid                       string `json:"dhcp6-duid"`
	Dhcp6IaPdLen                    string `json:"dhcp6-ia-pd-len"`
	AdvDhcp6PrefixSelectedInterface string `json:"adv_dhcp6_prefix_selected_interface"`
	Blockpriv                       string `json:"blockpriv"`
	Blockbogons                     string `json:"blockbogons"`
	Subnet                          string `json:"subnet"`
	Spoofmac                        string `json:"spoofmac"`
	Name                            string `json:"name"`
}

type VLAN struct {
	If     string `json:"if"`
	Tag    string `json:"tag"`
	Pcp    string `json:"pcp"`
	Descr  string `json:"descr"`
	Vlanif string `json:"vlanif"`
}

type interfaceListResponse struct {
	apiResponse
	Data map[string]*Interface `json:"data"`
}

type vlanListResponse struct {
	apiResponse
	Data []*VLAN `json:"data"`
}

// ListInterfaces returns the interfaces
func (s InterfaceService) ListInterfaces(ctx context.Context) ([]*Interface, error) {
	response, err := s.client.get(ctx, interfaceEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(interfaceListResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}

	interfaces := make([]*Interface, 0, len(resp.Data))
	for interfaceName, interfaceDetails := range resp.Data {
		interfaceDetails.Name = interfaceName
		interfaces = append(interfaces, interfaceDetails)
	}
	return interfaces, nil
}

// ListVLANs returns the interfaces
func (s InterfaceService) ListVLANs(ctx context.Context) ([]*VLAN, error) {
	response, err := s.client.get(ctx, interfaceVLANEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(vlanListResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}