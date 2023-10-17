package pfsenseapi

import (
	"context"
	"encoding/json"
	"strconv"

	"golang.org/x/exp/maps"
)

const (
	aliasEndpoint         = "api/v1/firewall/alias"
	aliasEntryEndpoint    = "api/v1/firewall/alias"
	ruleEndpoint          = "api/v1/firewall/rule"
	firewallApplyEndpoint = "api/v1/firewall/apply"
)

// FirewallService provides firewall API methods
type FirewallService service

// FirewallAlias represents a single firewall alias
type FirewallAlias struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Address string `json:"address"`
	Descr   string `json:"descr"`
	Detail  string `json:"detail"`
}

type firewallAliasListResponse struct {
	apiResponse
	Data []*FirewallAlias `json:"data"`
}

// ListAliases returns the aliases
func (s FirewallService) ListAliases(ctx context.Context) ([]*FirewallAlias, error) {
	response, err := s.client.get(ctx, aliasEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(firewallAliasListResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

type FirewallAliasRequest struct {
	Address []string `json:"address"`
	Descr   string   `json:"descr"`
	Detail  []string `json:"detail"`
	Name    string   `json:"name"`
	Type    string   `json:"type"`
}

type firewallAliasRequestCreate struct {
	FirewallAliasRequest
	Apply bool `json:"apply"`
}

type createAliasResponse struct {
	apiResponse
	Data *FirewallAlias `json:"data"`
}

// CreateAlias creates a new Alias.
func (s FirewallService) CreateAlias(
	ctx context.Context,
	newAlias FirewallAliasRequest,
	apply bool,
) (*FirewallAlias, error) {
	requestData := firewallAliasRequestCreate{
		FirewallAliasRequest: newAlias,
		Apply:                apply,
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}
	response, err := s.client.post(ctx, aliasEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(createAliasResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// DeleteAlias deletes a firewall Alias
func (s FirewallService) DeleteAlias(ctx context.Context, aliasToDelete string, apply bool) error {
	_, err := s.client.delete(
		ctx,
		aliasEndpoint,
		map[string]string{
			"id":    aliasToDelete,
			"apply": strconv.FormatBool(apply),
		},
	)
	if err != nil {
		return err
	}
	return nil
}

type firewallAliasRequestUpdate struct {
	FirewallAliasRequest
	Apply bool   `json:"apply"`
	Id    string `json:"id"`
}

// UpdateAlias modifies an existing alias
func (s FirewallService) UpdateAlias(
	ctx context.Context,
	aliasToUpdate string,
	newAliasData FirewallAliasRequest,
	apply bool,
) (*FirewallAlias, error) {
	requestData := firewallAliasRequestUpdate{
		FirewallAliasRequest: newAliasData,
		Apply:                apply,
		Id:                   aliasToUpdate,
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}

	response, err := s.client.put(ctx, aliasEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(createAliasResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// DeleteAliasEntry deletes a address from a firewall alias
func (s FirewallService) DeleteAliasEntry(ctx context.Context, aliasName string, address string, apply bool) error {
	_, err := s.client.delete(
		ctx,
		aliasEntryEndpoint,
		map[string]string{
			"name":    aliasName,
			"address": address,
			"apply":   strconv.FormatBool(apply),
		},
	)
	if err != nil {
		return err
	}
	return nil
}

type addAliasEntryRequest struct {
	Address []string `json:"address"`
	Apply   bool     `json:"apply"`
	Detail  []string `json:"detail"`
	Name    string   `json:"name"`
}

// AddAliasEntry adds an address to an existing Alias. The addresses to add is
// represented by a map with the address to add being the key, and the
// description being the value.
func (s FirewallService) AddAliasEntry(ctx context.Context, aliasName string, toAdd map[string]string, apply bool) error {
	newRequest := addAliasEntryRequest{
		Address: maps.Keys(toAdd),
		Apply:   apply,
		Detail:  maps.Values(toAdd),
		Name:    aliasName,
	}
	jsonData, err := json.Marshal(newRequest)
	if err != nil {
		return err
	}
	_, err = s.client.post(ctx, aliasEntryEndpoint, nil, jsonData)
	if err != nil {
		return err
	}
	return nil
}

// Apply applies pending firewall changes
func (s FirewallService) Apply(ctx context.Context) error {
	_, err := s.client.post(ctx, firewallApplyEndpoint, nil, nil)
	if err != nil {
		return err
	}
	return nil
}

type FirewallRule struct {
	Id           string            `json:"id"`
	Tracker      string            `json:"tracker"`
	Type         string            `json:"type"`
	Interface    string            `json:"interface"`
	Ipprotocol   string            `json:"ipprotocol"`
	Tag          string            `json:"tag"`
	Tagged       string            `json:"tagged"`
	Max          string            `json:"max"`
	MaxSrcNodes  string            `json:"max-src-nodes"`
	MaxSrcConn   string            `json:"max-src-conn"`
	MaxSrcStates string            `json:"max-src-states"`
	Statetimeout string            `json:"statetimeout"`
	Statetype    string            `json:"statetype"`
	Os           string            `json:"os"`
	Source       map[string]string `json:"source"`
	Destination  map[string]string `json:"destination"`
	Descr        string            `json:"descr"`
	Updated      struct {
		Time     string `json:"time"`
		Username string `json:"username"`
	} `json:"updated"`
	Created struct {
		Time     string `json:"time"`
		Username string `json:"username"`
	} `json:"created"`
}

type firewallRuleListResponse struct {
	apiResponse
	Data []*FirewallRule `json:"data"`
}

// ListRules returns the rules
func (s FirewallService) ListRules(ctx context.Context) ([]*FirewallRule, error) {
	response, err := s.client.get(ctx, ruleEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(firewallRuleListResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

// DeleteRule deletes a firewall Rule
func (s FirewallService) DeleteRule(ctx context.Context, tracker int, apply bool) error {
	_, err := s.client.delete(
		ctx,
		ruleEndpoint,
		map[string]string{
			"tracker": strconv.Itoa(tracker),
			"apply":   strconv.FormatBool(apply),
		},
	)
	if err != nil {
		return err
	}
	return nil
}

type FirewallRuleRequest struct {
	AckQueue     string   `json:"ackqueue"`
	DefaultQueue string   `json:"defaultqueue"`
	Descr        string   `json:"descr"`
	Direction    string   `json:"direction"`
	Disabled     bool     `json:"disabled"`
	Dnpipe       string   `json:"dnpipe"`
	Dst          string   `json:"dst"`
	DstPort      string   `json:"dstport"`
	Floating     bool     `json:"floating"`
	Gateway      string   `json:"gateway"`
	IcmpType     []string `json:"icmptype"`
	Interface    []string `json:"interface"`
	IpProtocol   string   `json:"ipprotocol"`
	Log          bool     `json:"log"`
	Pdnpipe      string   `json:"pdnpipe"`
	Protocol     string   `json:"protocol"`
	Quick        bool     `json:"quick"`
	Sched        string   `json:"sched"`
	Src          string   `json:"src"`
	SrcPort      string   `json:"srcport"`
	StateType    string   `json:"statetype"`
	TcpFlagsAny  bool     `json:"tcpflags_any"`
	TcpFlags1    []string `json:"tcpflags1"`
	TcpFlags2    []string `json:"tcpflags2"`
	Top          bool     `json:"top"`
	Type         string   `json:"type"`
}

type firewallRuleCreateRequest struct {
	FirewallRuleRequest
	Apply bool `json:"apply"`
}

type createRuleResponse struct {
	apiResponse
	Data *FirewallRule `json:"data"`
}

// CreateRule creates a new Rule
func (s FirewallService) CreateRule(
	ctx context.Context,
	newRule FirewallRuleRequest,
	apply bool,
) (*FirewallRule, error) {
	requestData := firewallRuleCreateRequest{
		FirewallRuleRequest: newRule,
		Apply:               apply,
	}
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}
	response, err := s.client.post(ctx, ruleEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(createRuleResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

type firewallRuleUpdateRequest struct {
	FirewallRuleRequest
	Apply   bool `json:"apply"`
	Tracker int  `json:"tracker"`
}

// UpdateRule modifies an existing rule
func (s FirewallService) UpdateRule(
	ctx context.Context,
	ruleToUpdate int,
	newRuleData FirewallRuleRequest,
	apply bool,
) (*FirewallRule, error) {
	requestData := firewallRuleUpdateRequest{
		FirewallRuleRequest: newRuleData,
		Apply:               apply,
		Tracker:             ruleToUpdate,
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}

	response, err := s.client.put(ctx, ruleEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(createRuleResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
