package pfsenseapi

import (
	"context"
	"encoding/json"
)

const (
	leasesEndpoint = "api/v1/services/dhcpd/lease"
)

type DHCPService service

func (s DHCPService) Leases(ctx context.Context) ([]*DHCPLease, error) {
	response, err := s.client.get(ctx, leasesEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(DhcpLeaseResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

type DhcpLeaseResponse struct {
	apiResponse
	Data []*DHCPLease `json:"data"`
}

type DHCPLease struct {
	IP                  *string
	Type                *string
	Mac                 *string
	If                  *string
	Starts              *string
	End                 *string
	Hostname            *string
	Descr               *string
	Online              *bool
	StaticmapArrayIndex *int
	State               *string
}
