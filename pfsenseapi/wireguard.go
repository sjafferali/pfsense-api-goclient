package pfsenseapi

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/markphelps/optional"
)

const (
	wireguardTunnelEndpoint  = "api/v2/vpn/wireguard/tunnel"
	wireguardTunnelsEndpoint = "api/v2/vpn/wireguard/tunnels"
	wireguardPeerEndpoint    = "api/v2/vpn/wireguard/peer"
	wireguardPeersEndpoint   = "api/v2/vpn/wireguard/peers"
	wireguardApplyEndpoint   = "api/v2/vpn/wireguard/apply"
)

// WireGuardService provides WireGuard API methods
type WireGuardService service

// WireGuardTunnel represents a WireGuard tunnel
type WireGuardTunnel struct {
	WireGuardTunnelRequest
	Id int `json:"id"`
}

// WireGuardTunnelRequest represents a request to create or update a WireGuard tunnel
type WireGuardTunnelRequest struct {
	Enabled    *optional.Bool   `json:"enabled,omitempty"`
	Name       string           `json:"name"`
	ListenPort *optional.String `json:"listenport,omitempty"`
	PrivateKey string           `json:"privatekey"`
	PublicKey  *optional.String `json:"publickey,omitempty"`
	MTU        *optional.Int    `json:"mtu,omitempty"`
	Descr      *optional.String `json:"descr,omitempty"`
	Address    []string         `json:"address"`
}

type wireguardTunnelListResponse struct {
	apiResponse
	Data []*WireGuardTunnel `json:"data"`
}

type wireguardTunnelGetResponse struct {
	apiResponse
	Data *WireGuardTunnel `json:"data"`
}

// ListWireGuardTunnels returns a list of WireGuard tunnels
func (s *WireGuardService) ListWireGuardTunnels(ctx context.Context) ([]*WireGuardTunnel, error) {
	response, err := s.client.get(ctx, wireguardTunnelsEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(wireguardTunnelListResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// GetWireGuardTunnel returns a WireGuard tunnel by id
func (s *WireGuardService) GetWireGuardTunnel(ctx context.Context, id int) (*WireGuardTunnel, error) {
	response, err := s.client.get(
		ctx,
		wireguardTunnelEndpoint,
		map[string]string{
			"id": strconv.Itoa(id),
		},
	)
	if err != nil {
		return nil, err
	}

	resp := new(wireguardTunnelGetResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// CreateWireGuardTunnel creates a new WireGuard tunnel
func (s *WireGuardService) CreateWireGuardTunnel(ctx context.Context, newTunnel WireGuardTunnelRequest) (*WireGuardTunnel, error) {
	jsonData, err := json.Marshal(newTunnel)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	response, err := s.client.post(ctx, wireguardTunnelEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(wireguardTunnelGetResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// UpdateWireGuardTunnel updates a WireGuard tunnel by id
func (s *WireGuardService) UpdateWireGuardTunnel(ctx context.Context, id int, updatedTunnel WireGuardTunnelRequest) (*WireGuardTunnel, error) {
	requestData := WireGuardTunnel{
		WireGuardTunnelRequest: updatedTunnel,
		Id:                     id,
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	response, err := s.client.patch(ctx, wireguardTunnelEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(wireguardTunnelGetResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// DeleteWireGuardTunnel deletes a WireGuard tunnel by id
func (s *WireGuardService) DeleteWireGuardTunnel(ctx context.Context, id int) (*WireGuardTunnel, error) {
	response, err := s.client.delete(
		ctx,
		wireguardTunnelEndpoint,
		map[string]string{
			"id": strconv.Itoa(id),
		},
	)
	if err != nil {
		return nil, err
	}

	resp := new(wireguardTunnelGetResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// WireGuardPeer represents a WireGuard peer
type WireGuardPeer struct {
	WireGuardPeerRequest
	Id int `json:"id"`
}

// WireGuardPeerRequest represents a request to create or update a WireGuard peer
type WireGuardPeerRequest struct {
	Enabled             *optional.Bool       `json:"enabled,omitempty"`
	Tun                 string               `json:"tun"`
	Endpoint            *optional.String     `json:"endpoint,omitempty"`
	Port                *optional.String     `json:"port,omitempty"`
	Descr               *optional.String     `json:"descr,omitempty"`
	PersistentKeepalive *optional.Int        `json:"persistentkeepalive,omitempty"`
	PublicKey           string               `json:"publickey"`
	PresharedKey        *optional.String     `json:"presharedkey,omitempty"`
	AllowedIPs          []WireGuardAllowedIP `json:"allowedips"`
}

// WireGuardAllowedIP represents an allowed IP for a WireGuard peer
type WireGuardAllowedIP struct {
	ParentId *optional.Int    `json:"parent_id,omitempty"`
	Id       *optional.Int    `json:"id,omitempty"`
	Address  string           `json:"address"`
	Mask     int              `json:"mask"`
	Descr    *optional.String `json:"descr,omitempty"`
}

type wireguardPeerListResponse struct {
	apiResponse
	Data []*WireGuardPeer `json:"data"`
}

type wireguardPeerGetResponse struct {
	apiResponse
	Data *WireGuardPeer `json:"data"`
}

// ListWireGuardPeers returns a list of WireGuard peers
func (s *WireGuardService) ListWireGuardPeers(ctx context.Context) ([]*WireGuardPeer, error) {
	response, err := s.client.get(ctx, wireguardPeersEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(wireguardPeerListResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// GetWireGuardPeer returns a WireGuard peer by id
func (s *WireGuardService) GetWireGuardPeer(ctx context.Context, id int) (*WireGuardPeer, error) {
	response, err := s.client.get(
		ctx,
		wireguardPeerEndpoint,
		map[string]string{
			"id": strconv.Itoa(id),
		},
	)
	if err != nil {
		return nil, err
	}

	resp := new(wireguardPeerGetResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// CreateWireGuardPeer creates a new WireGuard peer
func (s *WireGuardService) CreateWireGuardPeer(ctx context.Context, newPeer WireGuardPeerRequest) (*WireGuardPeer, error) {
	jsonData, err := json.Marshal(newPeer)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	response, err := s.client.post(ctx, wireguardPeerEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(wireguardPeerGetResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// UpdateWireGuardPeer updates a WireGuard peer
func (s *WireGuardService) UpdateWireGuardPeer(ctx context.Context, id int, updatedPeer WireGuardPeerRequest) (*WireGuardPeer, error) {
	requestData := WireGuardPeer{
		WireGuardPeerRequest: updatedPeer,
		Id:                   id,
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	response, err := s.client.patch(ctx, wireguardPeerEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(wireguardPeerGetResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// DeleteWireGuardPeer deletes a WireGuard peer
func (s *WireGuardService) DeleteWireGuardPeer(ctx context.Context, id int) (*WireGuardPeer, error) {
	response, err := s.client.delete(
		ctx,
		wireguardPeerEndpoint,
		map[string]string{
			"id": strconv.Itoa(id),
		},
	)
	if err != nil {
		return nil, err
	}

	resp := new(wireguardPeerGetResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// ApplyWireGuard applies pending WireGuard changes
func (s *WireGuardService) ApplyWireGuard(ctx context.Context) error {
	response, err := s.client.post(ctx, wireguardApplyEndpoint, nil, nil)
	if err != nil {
		return err
	}

	resp := new(apiResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return fmt.Errorf("error unmarshalling response: %w", err)
	}

	return nil
}
