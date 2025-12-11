package pfsenseapi

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/markphelps/optional"
)

const (
	firewallRuleEndpoint    = "api/v2/firewall/rule"
	firewallRulesEndpoint   = "api/v2/firewall/rules"
	firewallAliasEndpoint   = "api/v2/firewall/alias"
	firewallAliasesEndpoint = "api/v2/firewall/aliases"
	firewallApplyEndpoint   = "api/v2/firewall/apply"
)

// FirewallService provides Firewall API methods
type FirewallService service

// FirewallRule represents a firewall rule
type FirewallRule struct {
	FirewallRuleRequest
	Id int `json:"id"`
}

// FirewallRuleRequest represents a request to create or update a firewall rule
type FirewallRuleRequest struct {
	Type            *optional.String `json:"type,omitempty"`
	Interface       []string         `json:"interface,omitempty"`
	IPProtocol      *optional.String `json:"ipprotocol,omitempty"`
	Protocol        *optional.String `json:"protocol,omitempty"`
	Source          *optional.String `json:"source,omitempty"`
	SourcePort      *optional.String `json:"source_port,omitempty"`
	Destination     *optional.String `json:"destination,omitempty"`
	DestinationPort *optional.String `json:"destination_port,omitempty"`
	Gateway         *optional.String `json:"gateway,omitempty"`
	Sched           *optional.String `json:"sched,omitempty"`
	Descr           *optional.String `json:"descr,omitempty"`
	Log             *optional.Bool   `json:"log,omitempty"`
	Disabled        *optional.Bool   `json:"disabled,omitempty"`
	Direction       *optional.String `json:"direction,omitempty"`
	StateType       *optional.String `json:"statetype,omitempty"`
	Quick           *optional.Bool   `json:"quick,omitempty"`
	Floating        *optional.Bool   `json:"floating,omitempty"`
	DefaultQueue    *optional.String `json:"defaultqueue,omitempty"`
	AckQueue        *optional.String `json:"ackqueue,omitempty"`
	DnPipe          *optional.String `json:"dnpipe,omitempty"`
	PdnPipe         *optional.String `json:"pdnpipe,omitempty"`
	TcpFlagsAny     *optional.Bool   `json:"tcp_flags_any,omitempty"`
	TcpFlagsOutOf   []string         `json:"tcp_flags_out_of,omitempty"`
	TcpFlagsSet     []string         `json:"tcp_flags_set,omitempty"`
	Top             *optional.Bool   `json:"top,omitempty"`
}

type firewallRuleListResponse struct {
	apiResponse
	Data []*FirewallRule `json:"data"`
}

type firewallRuleGetResponse struct {
	apiResponse
	Data *FirewallRule `json:"data"`
}

// ListFirewallRules returns a list of firewall rules
func (s *FirewallService) ListFirewallRules(ctx context.Context) ([]*FirewallRule, error) {
	response, err := s.client.get(ctx, firewallRulesEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(firewallRuleListResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// GetFirewallRule returns a firewall rule by id
func (s *FirewallService) GetFirewallRule(ctx context.Context, id int) (*FirewallRule, error) {
	response, err := s.client.get(
		ctx,
		firewallRuleEndpoint,
		map[string]string{
			"id": strconv.Itoa(id),
		},
	)
	if err != nil {
		return nil, err
	}

	resp := new(firewallRuleGetResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// CreateFirewallRule creates a new firewall rule
func (s *FirewallService) CreateFirewallRule(ctx context.Context, newRule FirewallRuleRequest) (*FirewallRule, error) {
	jsonData, err := json.Marshal(newRule)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	response, err := s.client.post(ctx, firewallRuleEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(firewallRuleGetResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// UpdateFirewallRule updates a firewall rule
func (s *FirewallService) UpdateFirewallRule(ctx context.Context, id int, updatedRule FirewallRuleRequest) (*FirewallRule, error) {
	requestData := FirewallRule{
		FirewallRuleRequest: updatedRule,
		Id:                  id,
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	response, err := s.client.patch(ctx, firewallRuleEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(firewallRuleGetResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// DeleteFirewallRule deletes a firewall rule
func (s *FirewallService) DeleteFirewallRule(ctx context.Context, id int) (*FirewallRule, error) {
	response, err := s.client.delete(
		ctx,
		firewallRuleEndpoint,
		map[string]string{
			"id": strconv.Itoa(id),
		},
	)
	if err != nil {
		return nil, err
	}

	resp := new(firewallRuleGetResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// FirewallAlias represents a firewall alias
type FirewallAlias struct {
	FirewallAliasRequest
	Id int `json:"id"`
}

// FirewallAliasRequest represents a request to create or update a firewall alias
type FirewallAliasRequest struct {
	Name    string           `json:"name"`
	Type    string           `json:"type"`
	Descr   *optional.String `json:"descr,omitempty"`
	Address []string         `json:"address"`
	Detail  []string         `json:"detail"`
}

type firewallAliasListResponse struct {
	apiResponse
	Data []*FirewallAlias `json:"data"`
}

type firewallAliasGetResponse struct {
	apiResponse
	Data *FirewallAlias `json:"data"`
}

// ListFirewallAliases returns a list of firewall aliases
func (s *FirewallService) ListFirewallAliases(ctx context.Context) ([]*FirewallAlias, error) {
	response, err := s.client.get(ctx, firewallAliasesEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(firewallAliasListResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// GetFirewallAlias returns a firewall alias by id
func (s *FirewallService) GetFirewallAlias(ctx context.Context, id int) (*FirewallAlias, error) {
	response, err := s.client.get(
		ctx,
		firewallAliasEndpoint,
		map[string]string{
			"id": strconv.Itoa(id),
		},
	)
	if err != nil {
		return nil, err
	}

	resp := new(firewallAliasGetResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// CreateFirewallAlias creates a new firewall alias
func (s *FirewallService) CreateFirewallAlias(ctx context.Context, newAlias FirewallAliasRequest) (*FirewallAlias, error) {
	jsonData, err := json.Marshal(newAlias)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	response, err := s.client.post(ctx, firewallAliasEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(firewallAliasGetResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// UpdateFirewallAlias updates a firewall alias
func (s *FirewallService) UpdateFirewallAlias(ctx context.Context, id int, updatedAlias FirewallAliasRequest) (*FirewallAlias, error) {
	requestData := FirewallAlias{
		FirewallAliasRequest: updatedAlias,
		Id:                   id,
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	response, err := s.client.patch(ctx, firewallAliasEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(firewallAliasGetResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// DeleteFirewallAlias deletes a firewall alias
func (s *FirewallService) DeleteFirewallAlias(ctx context.Context, id int) (*FirewallAlias, error) {
	response, err := s.client.delete(
		ctx,
		firewallAliasEndpoint,
		map[string]string{
			"id": strconv.Itoa(id),
		},
	)
	if err != nil {
		return nil, err
	}

	resp := new(firewallAliasGetResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// ApplyFirewall applies pending firewall changes
func (s *FirewallService) ApplyFirewall(ctx context.Context) error {
	response, err := s.client.post(ctx, firewallApplyEndpoint, nil, nil)
	if err != nil {
		return err
	}

	resp := new(apiResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return fmt.Errorf("error unmarshalling response: %w", err)
	}

	return nil
}
