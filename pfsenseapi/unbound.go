package pfsenseapi

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
)

const (
	hostOverrideEndpoint = "api/v1/services/unbound/host_override"
)

// Unbound provides Unbound API methods
type UnboundService service

// Gateway represents a single routing gateway
type UnboundHostOverride struct {
	Aliases     *UnboundAliasesList `json:"aliases,omitempty"`
	Description string              `json:"descr"`
	Domain      string              `json:"domain"`
	Host        string              `json:"host"`
	IP          StringArray         `json:"ip"`
}

type UnboundAliasesList struct {
	Items []*UnboundHostOverrideAlias `json:"item"`
}

func (ual *UnboundAliasesList) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || (len(data) == 2 && string(data) == "\"\"") {
		*ual = UnboundAliasesList{}
		return nil
	}

	type unboundAliasesList UnboundAliasesList

	aux := &struct {
		*unboundAliasesList
	}{
		unboundAliasesList: (*unboundAliasesList)(ual),
	}

	return json.Unmarshal(data, &aux)
}

type UnboundHostOverrideAlias struct {
	Host        string `json:"host"`
	Domain      string `json:"domain"`
	Description string `json:"description"`
}

type apiWriteResponse[ResponseType any] struct {
	apiResponse
	Data *ResponseType `json:"data"`
}

type apiListResponse[ResponseType any] struct {
	apiResponse
	Data []ResponseType `json:"data"`
}

type createHostOverride struct {
	UnboundHostOverride
	Apply bool `json:"apply"`
}

func (s UnboundService) CreateHostOverride(
	ctx context.Context,
	hostOverride *UnboundHostOverride,
	apply bool,
) (*UnboundHostOverride, error) {
	jsonData, err := json.Marshal(&createHostOverride{
		UnboundHostOverride: *hostOverride,
		Apply:               apply,
	})

	if err != nil {
		return nil, err
	}

	response, err := s.client.post(ctx, hostOverrideEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	return s.parseWriteResponse(response)
}

type updateHostOverride struct {
	UnboundHostOverride
	Apply bool   `json:"apply"`
	Id    string `json:"id"`
}

func (s UnboundService) parseWriteResponse(
	response []byte,
) (*UnboundHostOverride, error) {
	resp := new(apiWriteResponse[UnboundHostOverride])
	if err := json.Unmarshal(response, resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func (s UnboundService) UpdateHostOverride(
	ctx context.Context,
	hostOverride *UnboundHostOverride,
	apply bool,
) (*UnboundHostOverride, error) {
	id, err := s.getHostOverridesObjectId(ctx, hostOverride.Host, hostOverride.Domain)
	if err != nil {
		return nil, fmt.Errorf("error finding override: %v", err)
	}

	jsonData, err := json.Marshal(&updateHostOverride{
		UnboundHostOverride: *hostOverride,
		Apply:               apply,
		Id:                  fmt.Sprint(id),
	})

	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %v", err)
	}

	response, err := s.client.put(ctx, hostOverrideEndpoint, nil, jsonData)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}

	return s.parseWriteResponse(response)
}

// Gets the index of the Unbound Host Override, you can only have one host override with the same host, name and ip type
func (s UnboundService) getHostOverridesObjectId(ctx context.Context, host string, domain string) (int, error) {
	list, err := s.ListHostOverrides(ctx)

	if err != nil {
		return 0, err
	}

	for i, item := range list {
		if item.Host == host && item.Domain == domain {
			return i, nil
		}
	}

	return 0, fmt.Errorf("Unable to find host override with host %s, domain %s", host, domain)
}

func (s UnboundService) ListHostOverrides(ctx context.Context) ([]*UnboundHostOverride, error) {
	response, err := s.client.get(ctx, hostOverrideEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(apiListResponse[*UnboundHostOverride])
	if err = json.Unmarshal(response, &resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (s UnboundService) DeleteHostOverride(ctx context.Context, host string, domain string, apply bool) error {
	id, err := s.getHostOverridesObjectId(ctx, host, domain)

	if err != nil {
		return err
	}

	queryMap := map[string]string{
		"id":    fmt.Sprint(id),
		"apply": strconv.FormatBool(apply),
	}

	_, err = s.client.delete(ctx, hostOverrideEndpoint, queryMap)

	if err != nil {
		return err
	}

	return nil
}
