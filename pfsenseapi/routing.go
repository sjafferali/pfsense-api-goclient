package pfsenseapi

import (
	"context"
	"encoding/json"
	"strconv"

	"golang.org/x/exp/maps"
)

const (
	gatewayEndpoint        = "api/v1/routing/gateway"
	defaultGatewayEndpoint = "api/v1/routing/gateway/default"
	routingApplyEndpoint   = "api/v1/routing/apply"
)

// RoutingService provides routing API methods
type RoutingService service

// Gateway represents a single routing gateway
type Gateway struct {
	Dynamic         bool   `json:"dynamic"`
	IpProtocol      string `json:"ipprotocol"`
	Gateway         string `json:"gateway"`
	Interface       string `json:"interface"`
	FriendlyIface   string `json:"friendlyiface"`
	FriendlyIfDescr string `json:"friendlyifdescr"`
	Name            string `json:"name"`
	Attribute       any    `json:"attribute"`
	IsDefaultGW     bool   `json:"isdefaultgw"`
	Monitor         string `json:"monitor"`
	Descr           string `json:"descr"`
	TierName        string `json:"tiername"`
}

type gatewayListResponse struct {
	apiResponse
	Data map[string]*Gateway `json:"data"`
}

// ListGateways returns the gateways
func (s RoutingService) ListGateways(ctx context.Context) ([]*Gateway, error) {
	response, err := s.client.get(ctx, gatewayEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(gatewayListResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}

	return maps.Values(resp.Data), nil
}

// GatewayRequest represents a single gateway to be created or modified. This
// type is use for creations and updates.
type GatewayRequest struct {
	ActionDisable  bool   `json:"action_disable"`
	AlertInterval  int    `json:"alert_interval"`
	Apply          bool   `json:"apply"`
	DataPayload    int    `json:"data_payload"`
	Descr          string `json:"descr"`
	Disabled       bool   `json:"disabled"`
	ForceDown      bool   `json:"force_down"`
	Gateway        string `json:"gateway"`
	Interface      string `json:"interface"`
	Interval       int    `json:"interval"`
	IpProtocol     string `json:"ipprotocol"`
	LatencyHigh    int    `json:"latencyhigh"`
	LatencyLow     int    `json:"latencylow"`
	LossInterval   int    `json:"loss_interval"`
	LossHigh       int    `json:"losshigh"`
	LossLow        int    `json:"losslow"`
	Monitor        string `json:"monitor"`
	MonitorDisable bool   `json:"monitor_disable"`
	Name           string `json:"name"`
	TimePeriod     int    `json:"time_period"`
	Weight         int    `json:"weight"`
}

type createGatewayResponse struct {
	apiResponse
	Data *Gateway `json:"data"`
}

// CreateGateway creates a new Gateway
func (s RoutingService) CreateGateway(
	ctx context.Context,
	newGateway GatewayRequest,
) (*Gateway, error) {
	jsonData, err := json.Marshal(newGateway)
	if err != nil {
		return nil, err
	}

	response, err := s.client.post(ctx, gatewayEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(createGatewayResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// DeleteGateway deletes a Gateway
func (s RoutingService) DeleteGateway(ctx context.Context, gatewayID int) error {
	_, err := s.client.delete(ctx, gatewayEndpoint, map[string]string{"id": strconv.Itoa(gatewayID)})
	if err != nil {
		return err
	}
	return nil
}

// UpdateGateway modifies a existing gateway
func (s RoutingService) UpdateGateway(
	ctx context.Context,
	gatewayToUpdate GatewayRequest,
) (*Gateway, error) {
	jsonData, err := json.Marshal(gatewayToUpdate)
	if err != nil {
		return nil, err
	}

	response, err := s.client.put(ctx, gatewayEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(createGatewayResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

type DefaultGatewayRequest struct {
	DefaultGW4 string `json:"defaultgw4"`
	DefaultGW6 string `json:"defaultgw6"`
	Apply      bool   `json:"apply"`
}

// SetDefaultGateway sets the default gateway
func (s RoutingService) SetDefaultGateway(ctx context.Context, newDefaultGateway DefaultGatewayRequest) error {
	jsonData, err := json.Marshal(newDefaultGateway)
	if err != nil {
		return err
	}
	_, err = s.client.put(ctx, defaultGatewayEndpoint, nil, jsonData)
	if err != nil {
		return err
	}
	return nil
}

// Apply applies pending routing changes
func (s RoutingService) Apply(ctx context.Context) error {
	_, err := s.client.post(ctx, routingApplyEndpoint, nil, nil)
	if err != nil {
		return err
	}
	return nil
}
