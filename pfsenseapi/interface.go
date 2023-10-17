package pfsenseapi

import (
	"context"
	"encoding/json"
	"strconv"
)

const (
	interfaceEndpoint      = "api/v1/interface"
	interfaceVLANEndpoint  = "api/v1/interface/vlan"
	interfaceGroupEndpoint = "api/v1/interface/group"
	interfaceApplyEndpoint = "api/v1/interface/apply"
)

// InterfaceService provides interface API methods
type InterfaceService service

// Interface represents a single interface.
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

type interfaceListResponse struct {
	apiResponse
	Data map[string]*Interface `json:"data"`
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

// DeleteInterface deletes the interface. The interfaceID can be specified in
// either the interface's descriptive name, the pfSense ID (wan, lan, optx), or
// the physical interface id (e.g. igb0).
func (s InterfaceService) DeleteInterface(ctx context.Context, interfaceID string) error {
	_, err := s.client.delete(
		ctx,
		interfaceEndpoint,
		map[string]string{
			"if": interfaceID,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

type InterfaceRequest struct {
	AdvDhcpConfigAdvanced         bool     `json:"adv_dhcp_config_advanced"`
	AdvDhcpConfigFileOverride     bool     `json:"adv_dhcp_config_file_override"`
	AdvDhcpConfigFileOverrideFile string   `json:"adv_dhcp_config_file_override_file"`
	AdvDhcpOptionModifiers        string   `json:"adv_dhcp_option_modifiers"`
	AdvDhcpPtBackoffCutoff        int      `json:"adv_dhcp_pt_backoff_cutoff"`
	AdvDhcpPtInitialInterval      int      `json:"adv_dhcp_pt_initial_interval"`
	AdvDhcpPtReboot               int      `json:"adv_dhcp_pt_reboot"`
	AdvDhcpPtRetry                int      `json:"adv_dhcp_pt_retry"`
	AdvDhcpPtSelectTimeout        int      `json:"adv_dhcp_pt_select_timeout"`
	AdvDhcpPtTimeout              int      `json:"adv_dhcp_pt_timeout"`
	AdvDhcpRequestOptions         string   `json:"adv_dhcp_request_options"`
	AdvDhcpRequiredOptions        string   `json:"adv_dhcp_required_options"`
	AdvDhcpSendOptions            string   `json:"adv_dhcp_send_options"`
	AliasAddress                  string   `json:"alias-address"`
	AliasSubnet                   int      `json:"alias-subnet"`
	Apply                         bool     `json:"apply"`
	Blockbogons                   bool     `json:"blockbogons"`
	Blockpriv                     bool     `json:"blockpriv"`
	Descr                         string   `json:"descr"`
	Dhcpcvpt                      int      `json:"dhcpcvpt"`
	Dhcphostname                  string   `json:"dhcphostname"`
	Dhcprejectfrom                []string `json:"dhcprejectfrom"`
	Dhcpvlanenable                bool     `json:"dhcpvlanenable"`
	Enable                        bool     `json:"enable"`
	Gateway                       string   `json:"gateway"`
	Gateway6Rd                    string   `json:"gateway-6rd"`
	Gatewayv6                     string   `json:"gatewayv6"`
	If                            string   `json:"if"`
	Ipaddr                        string   `json:"ipaddr"`
	Ipaddrv6                      string   `json:"ipaddrv6"`
	Ipv6Usev4Iface                bool     `json:"ipv6usev4iface"`
	Media                         string   `json:"media"`
	Mss                           string   `json:"mss"`
	Mtu                           int      `json:"mtu"`
	Prefix6Rd                     string   `json:"prefix-6rd"`
	Prefix6RdV4Plen               int      `json:"prefix-6rd-v4plen"`
	Spoofmac                      string   `json:"spoofmac"`
	Subnet                        int      `json:"subnet"`
	Subnetv6                      string   `json:"subnetv6"`
	Track6Interface               string   `json:"track6-interface"`
	Track6PrefixIdHex             int      `json:"track6-prefix-id-hex"`
	Type                          string   `json:"type"`
	Type6                         string   `json:"type6"`
}

type createInterfaceResponse struct {
	apiResponse
	Data *Interface `json:"data"`
}

// CreateInterface creates a new interface.
func (s InterfaceService) CreateInterface(
	ctx context.Context,
	newInterface InterfaceRequest,
) (*Interface, error) {
	jsonData, err := json.Marshal(newInterface)
	if err != nil {
		return nil, err
	}

	response, err := s.client.post(ctx, interfaceEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(createInterfaceResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

type interfaceRequestUpdate struct {
	InterfaceRequest
	Id string `json:"id"`
}

// UpdateInterface modifies an existing interface.
func (s InterfaceService) UpdateInterface(
	ctx context.Context,
	idToUpdate int,
	interfaceData InterfaceRequest,
) (*Interface, error) {
	requestData := interfaceRequestUpdate{
		InterfaceRequest: interfaceData,
		Id:               strconv.Itoa(idToUpdate),
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}

	response, err := s.client.put(ctx, interfaceEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(createInterfaceResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// VLAN represents a single VLAN.
type VLAN struct {
	If     string `json:"if"`
	Tag    string `json:"tag"`
	Pcp    string `json:"pcp"`
	Descr  string `json:"descr"`
	Vlanif string `json:"vlanif"`
}

type vlanListResponse struct {
	apiResponse
	Data []*VLAN `json:"data"`
}

// ListVLANs returns the VLANs
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

// DeleteVLAN deletes a VLAN.
func (s InterfaceService) DeleteVLAN(ctx context.Context, idToDelete int) error {
	_, err := s.client.delete(
		ctx,
		interfaceVLANEndpoint,
		map[string]string{
			"id": strconv.Itoa(idToDelete),
		},
	)
	if err != nil {
		return err
	}
	return nil
}

type VLANRequest struct {
	Descr string `json:"descr"`
	If    string `json:"if"`
	Pcp   int    `json:"pcp"`
	Tag   int    `json:"tag"`
}

type createVLANResponse struct {
	apiResponse
	Data *VLAN `json:"data"`
}

// CreateVLAN creates a new VLAN.
func (s InterfaceService) CreateVLAN(
	ctx context.Context,
	newVLAN VLANRequest,
) (*VLAN, error) {
	jsonData, err := json.Marshal(newVLAN)
	if err != nil {
		return nil, err
	}

	response, err := s.client.post(ctx, interfaceVLANEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(createVLANResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

type vlanRequestUpdate struct {
	VLANRequest
	Id int `json:"id"`
}

// UpdateVLAN modifies an existing VLAN.
func (s InterfaceService) UpdateVLAN(
	ctx context.Context,
	idToUpdate int,
	vlanData VLANRequest,
) (*VLAN, error) {
	requestData := vlanRequestUpdate{
		VLANRequest: vlanData,
		Id:          idToUpdate,
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}

	response, err := s.client.put(ctx, interfaceEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(createVLANResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

type InterfaceGroup struct {
	Members string `json:"members"`
	Descr   string `json:"descr"`
	Ifname  string `json:"ifname"`
}

type interfaceGroupListResponse struct {
	apiResponse
	Data []*InterfaceGroup `json:"data"`
}

// ListInterfaceGroups returns the interface groups.
func (s InterfaceService) ListInterfaceGroups(ctx context.Context) ([]*InterfaceGroup, error) {
	response, err := s.client.get(ctx, interfaceGroupEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(interfaceGroupListResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

// DeleteInterfaceGroup deletes an interface group.
func (s InterfaceService) DeleteInterfaceGroup(ctx context.Context, idToDelete int) error {
	_, err := s.client.delete(
		ctx,
		interfaceGroupEndpoint,
		map[string]string{
			"id": strconv.Itoa(idToDelete),
		},
	)
	if err != nil {
		return err
	}
	return nil
}

type InterfaceGroupRequestCreate struct {
	Descr   string   `json:"descr"`
	Members []string `json:"members"`
	Ifname  string   `json:"ifname"`
}

type createInterfaceGroupResponse struct {
	apiResponse
	Data *InterfaceGroup `json:"data"`
}

// CreateInterfaceGroup creates a new interface group.
func (s InterfaceService) CreateInterfaceGroup(
	ctx context.Context,
	newGroup InterfaceGroupRequestCreate,
) (*InterfaceGroup, error) {
	jsonData, err := json.Marshal(newGroup)
	if err != nil {
		return nil, err
	}
	response, err := s.client.post(ctx, interfaceGroupEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(createInterfaceGroupResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

type InterfaceGroupRequestUpdate struct {
	Descr   string   `json:"descr"`
	Id      string   `json:"id"`
	Members []string `json:"members"`
}

// UpdateInterfaceGroup updates an existing interface group.
func (s InterfaceService) UpdateInterfaceGroup(
	ctx context.Context,
	groupData InterfaceGroupRequestUpdate,
) (*InterfaceGroup, error) {
	jsonData, err := json.Marshal(groupData)
	if err != nil {
		return nil, err
	}

	response, err := s.client.put(ctx, interfaceGroupEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(createInterfaceGroupResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

type applyInterfaceRequest struct {
	Async bool `json:"async"`
}

// Apply applies pending interface changes
func (s InterfaceService) Apply(ctx context.Context, async bool) error {
	requestData := applyInterfaceRequest{
		Async: async,
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return err
	}

	_, err = s.client.post(ctx, interfaceApplyEndpoint, nil, jsonData)
	if err != nil {
		return err
	}
	return nil
}
