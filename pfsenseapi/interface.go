package pfsenseapi

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/markphelps/optional"
)

const (
	interfaceEndpoint        = "api/v2/interface"
	interfacesEndpoint       = "api/v2/interfaces"
	interfaceVLANEndpoint    = "api/v2/interface/vlan"
	interfaceVLANsEndpoint   = "api/v2/interface/vlans"
	interfaceGroupEndpoint   = "api/v2/interface/group"
	interfaceGroupsEndpoint  = "api/v2/interface/groups"
	interfaceBridgeEndpoint  = "api/v2/interface/bridge"
	interfaceBridgesEndpoint = "api/v2/interface/bridges"
	interfaceApplyEndpoint   = "api/v2/interface/apply"
)

// InterfaceService provides interface API methods
type InterfaceService service

type Interface struct {
	InterfaceRequest
	Id string `json:"id"`
}

type interfaceListResponse struct {
	apiResponse
	Data []*Interface `json:"data"`
}

// GetInterface returns a single interface.
func (s InterfaceService) GetInterface(ctx context.Context, interfaceID string) (*Interface, error) {
	response, err := s.client.get(
		ctx,
		interfaceEndpoint,
		map[string]string{
			"if": interfaceID,
		},
	)
	if err != nil {
		return nil, err
	}

	resp := new(createInterfaceResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return resp.Data, nil
}

// ListInterfaces returns a list of the interfaces.
func (s InterfaceService) ListInterfaces(ctx context.Context) ([]*Interface, error) {
	response, err := s.client.get(ctx, interfacesEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(interfaceListResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// DeleteInterface deletes the interface. The interfaceID can be specified in
// either the interface's descriptive name, the pfSense ID (wan, lan, optx), or
// the physical interface id (e.g. igb0).
func (s InterfaceService) DeleteInterface(ctx context.Context, interfaceID string) error {
	response, err := s.client.delete(
		ctx,
		interfaceEndpoint,
		map[string]string{
			"if": interfaceID,
		},
	)
	if err != nil {
		return err
	}

	resp := new(apiResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return fmt.Errorf("error unmarshalling response: %w", err)
	}

	return nil
}

type InterfaceRequest struct {
	If                            string           `json:"if"`
	Enable                        *optional.Bool   `json:"enable,omitempty"`
	Descr                         string           `json:"descr"`
	Spoofmac                      *optional.String `json:"spoofmac,omitempty"`
	Mtu                           *optional.Int32  `json:"mtu,omitempty"`
	Mss                           *optional.Int32  `json:"mss,omitempty"`
	Media                         *optional.String `json:"media,omitempty"`
	Mediaopt                      *optional.String `json:"mediaopt,omitempty"`
	Blockpriv                     *optional.Bool   `json:"blockpriv,omitempty"`
	Blockbogons                   *optional.Bool   `json:"blockbogons,omitempty"`
	Typev4                        string           `json:"typev4"`
	Ipaddr                        string           `json:"ipaddr"`
	Subnet                        int32            `json:"subnet"`
	Gateway                       *optional.String `json:"gateway,omitempty"`
	AliasSubnet                   *optional.Int32  `json:"alias_subnet,omitempty"`
	AdvDhcpPtTimeout              *optional.Int32  `json:"adv_dhcp_pt_timeout,omitempty"`
	AdvDhcpPtRetry                *optional.Int32  `json:"adv_dhcp_pt_retry,omitempty"`
	AdvDhcpPtSelectTimeout        *optional.Int32  `json:"adv_dhcp_pt_select_timeout,omitempty"`
	AdvDhcpPtReboot               *optional.Int32  `json:"adv_dhcp_pt_reboot,omitempty"`
	AdvDhcpPtBackoffCutoff        *optional.Int32  `json:"adv_dhcp_pt_backoff_cutoff,omitempty"`
	AdvDhcpPtInitialInterval      *optional.Int32  `json:"adv_dhcp_pt_initial_interval,omitempty"`
	AdvDhcpSendOptions            *optional.String `json:"adv_dhcp_send_options,omitempty"`
	AdvDhcpRequestOptions         *optional.String `json:"adv_dhcp_request_options,omitempty"`
	AdvDhcpRequiredOptions        *optional.String `json:"adv_dhcp_required_options,omitempty"`
	AdvDhcpOptionModifiers        *optional.String `json:"adv_dhcp_option_modifiers,omitempty"`
	AdvDhcpConfigFileOverridePath *optional.String `json:"adv_dhcp_config_file_override_path,omitempty"`
	Typev6                        *optional.String `json:"typev6,omitempty"`
	Ipaddrv6                      string           `json:"ipaddrv6"`
	Subnetv6                      int32            `json:"subnetv6"`
	Gatewayv6                     *optional.String `json:"gatewayv6,omitempty"`
	Prefix6Rd                     string           `json:"prefix_6rd"`
	Gateway6Rd                    string           `json:"gateway_6rd"`
	Prefix6RdV4Plen               int32            `json:"prefix_6rd_v4plen"`
	Track6Interface               string           `json:"track6_interface"`
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
		return nil, fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	response, err := s.client.post(ctx, interfaceEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(createInterfaceResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return resp.Data, nil
}

// UpdateInterface modifies an existing interface.
func (s InterfaceService) UpdateInterface(
	ctx context.Context,
	idToUpdate string,
	interfaceData InterfaceRequest,
) (*Interface, error) {
	requestData := Interface{
		InterfaceRequest: interfaceData,
		Id:               idToUpdate,
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	response, err := s.client.patch(ctx, interfaceEndpoint, nil, jsonData)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	resp := new(createInterfaceResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// VLAN represents a single VLAN.
type VLAN struct {
	VLANRequest
	Id int `json:"id"`
}

type vlanListResponse struct {
	apiResponse
	Data []*VLAN `json:"data"`
}

// ListVLANs returns the VLANs
func (s InterfaceService) ListVLANs(ctx context.Context) ([]*VLAN, error) {
	response, err := s.client.get(ctx, interfaceVLANsEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(vlanListResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// GetVLAN returns the VLAN with the given ID.
func (s InterfaceService) GetVLAN(ctx context.Context, id int) (*VLAN, error) {
	response, err := s.client.get(
		ctx,
		interfaceVLANEndpoint,
		map[string]string{
			"id": strconv.Itoa(id),
		},
	)
	if err != nil {
		return nil, err
	}

	resp := new(createVLANResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return resp.Data, nil
}

// DeleteVLAN deletes a VLAN.
func (s InterfaceService) DeleteVLAN(ctx context.Context, idToDelete int) error {
	response, err := s.client.delete(
		ctx,
		interfaceVLANEndpoint,
		map[string]string{
			"id": strconv.Itoa(idToDelete),
		},
	)
	if err != nil {
		return err
	}

	resp := new(apiResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return fmt.Errorf("error unmarshalling response: %w", err)
	}

	return nil
}

type VLANRequest struct {
	If     string           `json:"if"`
	Tag    int              `json:"tag"`
	Vlanif *optional.String `json:"vlanif,omitempty"`
	Pcp    *optional.Int    `json:"pcp,omitempty"`
	Descr  *optional.String `json:"descr,omitempty"`
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
		return nil, fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	response, err := s.client.post(ctx, interfaceVLANEndpoint, nil, jsonData)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	resp := new(createVLANResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return resp.Data, nil
}

// UpdateVLAN modifies an existing VLAN.
func (s InterfaceService) UpdateVLAN(
	ctx context.Context,
	idToUpdate int,
	vlanData VLANRequest,
) (*VLAN, error) {
	requestData := VLAN{
		VLANRequest: vlanData,
		Id:          idToUpdate,
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	response, err := s.client.patch(ctx, interfaceVLANEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(createVLANResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return resp.Data, nil
}

type InterfaceGroup struct {
	InterfaceGroupRequest
	Id int `json:"id"`
}

type interfaceGroupListResponse struct {
	apiResponse
	Data []*InterfaceGroup `json:"data"`
}

// ListInterfaceGroups returns the interface groups.
func (s InterfaceService) ListInterfaceGroups(ctx context.Context) ([]*InterfaceGroup, error) {
	response, err := s.client.get(ctx, interfaceGroupsEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(interfaceGroupListResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// PutInterfaceGroups replaces all interface groups with the given list.
func (s InterfaceService) PutInterfaceGroups(ctx context.Context, groups []*InterfaceGroupRequest) ([]*InterfaceGroup, error) {
	jsonData, err := json.Marshal(groups)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	response, err := s.client.put(ctx, interfaceGroupsEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(interfaceGroupListResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// GetInterfaceGroup returns the interface group with the given ID.
func (s InterfaceService) GetInterfaceGroup(ctx context.Context, id int) (*InterfaceGroup, error) {
	response, err := s.client.get(
		ctx,
		interfaceGroupEndpoint,
		map[string]string{
			"id": strconv.Itoa(id),
		},
	)
	if err != nil {
		return nil, err
	}

	resp := new(interfaceGroupResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return resp.Data, nil

}

// DeleteInterfaceGroup deletes an interface group.
func (s InterfaceService) DeleteInterfaceGroup(ctx context.Context, idToDelete int) error {
	response, err := s.client.delete(
		ctx,
		interfaceGroupEndpoint,
		map[string]string{
			"id": strconv.Itoa(idToDelete),
		},
	)
	if err != nil {
		return err
	}

	resp := new(apiResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return fmt.Errorf("error unmarshalling response: %w", err)
	}

	return nil
}

// InterfaceGroupRequest represents the request to create or update an interface group.
type InterfaceGroupRequest struct {
	Ifname  string   `json:"ifname"`
	Members []string `json:"members"`
	Descr   string   `json:"descr"`
}

type interfaceGroupResponse struct {
	apiResponse
	Data *InterfaceGroup `json:"data"`
}

// CreateInterfaceGroup creates a new interface group.
func (s InterfaceService) CreateInterfaceGroup(
	ctx context.Context,
	newGroup InterfaceGroupRequest,
) (*InterfaceGroup, error) {
	jsonData, err := json.Marshal(newGroup)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request payload into json: %w", err)
	}
	response, err := s.client.post(ctx, interfaceGroupEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(interfaceGroupResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return resp.Data, nil
}

// UpdateInterfaceGroup updates an existing interface group.
func (s InterfaceService) UpdateInterfaceGroup(
	ctx context.Context,
	idToUpdate int,
	groupData InterfaceGroupRequest,
) (*InterfaceGroup, error) {
	requestData := InterfaceGroup{
		InterfaceGroupRequest: groupData,
		Id:                    idToUpdate,
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	response, err := s.client.patch(ctx, interfaceGroupEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(interfaceGroupResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return resp.Data, nil
}

// Apply applies pending interface changes
func (s InterfaceService) Apply(ctx context.Context) error {
	response, err := s.client.post(ctx, interfaceApplyEndpoint, nil, nil)
	if err != nil {
		return err
	}

	resp := new(apiResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return fmt.Errorf("error unmarshalling response: %w", err)
	}

	return nil
}

// InterfaceBridge represents a single bridge.
type InterfaceBridge struct {
	InterfaceBridgeRequest
	Id string `json:"id"`
}

type interfaceBridgeListResponse struct {
	apiResponse
	Data []*InterfaceBridge `json:"data"`
}

// ListInterfaceBridges returns the bridges.
func (s InterfaceService) ListInterfaceBridges(ctx context.Context) ([]*InterfaceBridge, error) {
	response, err := s.client.get(ctx, interfaceBridgesEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(interfaceBridgeListResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

type interfaceBridgeResponse struct {
	apiResponse
	Data *InterfaceBridge `json:"data"`
}

// GetInterfaceBridge returns the bridge with the given ID.
func (s InterfaceService) GetInterfaceBridge(ctx context.Context, id string) (*InterfaceBridge, error) {
	response, err := s.client.get(
		ctx,
		interfaceBridgeEndpoint,
		map[string]string{
			"id": id,
		},
	)
	if err != nil {
		return nil, err
	}

	resp := new(interfaceBridgeResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return resp.Data, nil
}

// DeleteInterfaceBridge deletes a bridge.
func (s InterfaceService) DeleteInterfaceBridge(ctx context.Context, idToDelete string) error {
	response, err := s.client.delete(
		ctx,
		interfaceBridgeEndpoint,
		map[string]string{
			"id": idToDelete,
		},
	)
	if err != nil {
		return err
	}

	resp := new(apiResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return fmt.Errorf("error unmarshalling response: %w", err)
	}

	return nil
}

// CreateInterfaceBridge creates a new bridge.
func (s InterfaceService) CreateInterfaceBridge(
	ctx context.Context,
	newBridge InterfaceBridgeRequest,
) (*InterfaceBridge, error) {
	jsonData, err := json.Marshal(newBridge)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	response, err := s.client.post(ctx, interfaceBridgeEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(interfaceBridgeResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return resp.Data, nil
}

// UpdateInterfaceBridge updates an existing bridge.
func (s InterfaceService) UpdateInterfaceBridge(
	ctx context.Context,
	idToUpdate string,
	bridgeData InterfaceBridgeRequest,
) (*InterfaceBridge, error) {
	requestData := InterfaceBridge{
		InterfaceBridgeRequest: bridgeData,
		Id:                     idToUpdate,
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	response, err := s.client.patch(ctx, interfaceBridgeEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(interfaceBridgeResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return resp.Data, nil
}

// InterfaceBridgeRequest represents the request to create or update a bridge.
type InterfaceBridgeRequest struct {
	Members  []string `json:"members"`
	Descr    string   `json:"descr"`
	Bridgeif string   `json:"bridgeif"`
}
