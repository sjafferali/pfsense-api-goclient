package pfsenseapi

import (
	"context"
	"encoding/json"
)

const (
	leasesEndpoint        = "api/v1/services/dhcpd/lease"
	staticMappingEndpoint = "api/v1/services/dhcpd/static_mapping"
)

type DHCPService service

type dhcpLeaseResponse struct {
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

type dhcpStaticMappingResponse struct {
	apiResponse
	Data []*DHCPStaticMapping `json:"data"`
}

type DHCPStaticMapping struct {
	ID                     int    `json:"id"`
	Mac                    string `json:"mac"`
	Cid                    string `json:"cid"`
	IPaddr                 string `json:"ipaddr"`
	Hostname               string `json:"hostname"`
	Descr                  string `json:"descr"`
	Filename               string `json:"filename"`
	Rootpath               string `json:"rootpath"`
	DefaultLeaseTime       string `json:"defaultleasetime"`
	MaxLeaseTime           string `json:"maxleasetime"`
	Gateway                string `json:"gateway"`
	Domain                 string `json:"domain"`
	DomainSearchList       string `json:"domainsearchlist"`
	DDNSDomain             string `json:"ddnsdomain"`
	DDNSdomainprimary      string `json:"ddnsdomainprimary"`
	DDNSDomainSecondary    string `json:"ddnsdomainsecondary"`
	DDNSDomainkeyName      string `json:"ddnsdomainkeyname"`
	DDNSDomainkeyAlgorithm string `json:"ddnsdomainkeyalgorithm"`
	DDNSDomainkey          string `json:"ddnsdomainkey"`
	TFTP                   string `json:"tftp"`
	LDAP                   string `json:"ldap"`
	NextServer             string `json:"nextserver"`
	Filename32             string `json:"filename32"`
	Filename64             string `json:"filename64"`
	Filename32Arm          string `json:"filename32arm"`
	Filename64Arm          string `json:"filename64arm"`
	NumberOptions          string `json:"numberoptions"`
}

type DHCPStaticMappingRequest struct {
	ArpTableStaticEntry bool     `json:"arp_table_static_entry"`
	Cid                 string   `json:"cid"`
	Descr               string   `json:"descr"`
	DNSServer           []string `json:"dnsserver"`
	Domain              string   `json:"domain"`
	DomainSearchList    []string `json:"domainsearchlist"`
	Gateway             string   `json:"gateway"`
	Hostname            string   `json:"hostname"`
	Id                  int      `json:"id"`
	Interface           string   `json:"interface"`
	Ipaddr              string   `json:"ipaddr"`
	Mac                 string   `json:"mac"`
}

func (s DHCPService) ListLeases(ctx context.Context) ([]*DHCPLease, error) {
	response, err := s.client.get(ctx, leasesEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(dhcpLeaseResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (s DHCPService) ListStaticMappings(ctx context.Context, netInterface string) ([]*DHCPStaticMapping, error) {
	queryMap := map[string]string{
		"interface": netInterface,
	}
	response, err := s.client.get(ctx, staticMappingEndpoint, queryMap)
	if err != nil {
		return nil, err
	}

	resp := new(dhcpStaticMappingResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (s DHCPService) CreateStaticMapping(ctx context.Context, newStaticMapping DHCPStaticMappingRequest) error {
	jsonData, err := json.Marshal(newStaticMapping)
	if err != nil {
		return err
	}
	_, err = s.client.post(ctx, staticMappingEndpoint, nil, jsonData)
	if err != nil {
		return err
	}
	return nil
}
