package pfsenseapi

import (
	"context"
	"encoding/json"
	"strconv"
)

const (
	leasesEndpoint        = "api/v1/services/dhcpd/lease"
	staticMappingEndpoint = "api/v1/services/dhcpd/static_mapping"
)

// DHCPService provides DHCP API methods
type DHCPService service

type dhcpLeaseResponse struct {
	apiResponse
	Data []*DHCPLease `json:"data"`
}

// DHCPLease represents a single DHCP lease
type DHCPLease struct {
	Ip                  string `json:"ip"`
	Type                string `json:"type"`
	Mac                 string `json:"mac"`
	If                  string `json:"if"`
	Starts              string `json:"starts"`
	Ends                string `json:"ends"`
	Hostname            string `json:"hostname"`
	Descr               string `json:"descr"`
	Online              string `json:"online"`
	StaticmapArrayIndex int    `json:"staticmap_array_index"`
	State               string `json:"state"`
}

type dhcpStaticMappingResponse struct {
	apiResponse
	Data []*DHCPStaticMapping `json:"data"`
}

// DHCPStaticMapping represents a single DHCP static reservation
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
	DDNSDomainPrimary      string `json:"ddnsdomainprimary"`
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

// DHCPStaticMappingRequest represents a single DHCP static reservation. This
// type is used for updating or creating a new static reservation.
type DHCPStaticMappingRequest struct {
	ArpTableStaticEntry bool     `json:"arp_table_static_entry"`
	Cid                 string   `json:"cid"`
	Descr               string   `json:"descr"`
	DNSServer           []string `json:"dnsserver"`
	Domain              string   `json:"domain"`
	DomainSearchList    []string `json:"domainsearchlist"`
	Gateway             string   `json:"gateway"`
	Hostname            string   `json:"hostname"`
	Interface           string   `json:"interface"`
	Ipaddr              string   `json:"ipaddr"`
	Mac                 string   `json:"mac"`
}

// ListLeases returns a list of the DHCP leases
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

// ListStaticMappings returns a list of the static reservations for the interface
// provided. The interface can be either the interface's
// descriptive name, the pfSense interface ID (e.g. wan, lan, optx), or the real
// interface ID (e.g. igb0).
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

type createStaticMappingResponse struct {
	apiResponse
	Data *DHCPStaticMapping `json:"data"`
}

// CreateStaticMapping creates a new DHCP static mapping.
func (s DHCPService) CreateStaticMapping(
	ctx context.Context,
	newStaticMapping DHCPStaticMappingRequest,
) (*DHCPStaticMapping, error) {
	jsonData, err := json.Marshal(newStaticMapping)
	if err != nil {
		return nil, err
	}
	response, err := s.client.post(ctx, staticMappingEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(createStaticMappingResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

type dhcpStaticMappingRequestUpdate struct {
	DHCPStaticMappingRequest
	Id int `json:"id"`
}

// UpdateStaticMapping modifies a DHCP static mapping.
func (s DHCPService) UpdateStaticMapping(
	ctx context.Context,
	idToUpdate int,
	mappingData DHCPStaticMappingRequest,
) (*DHCPStaticMapping, error) {
	requestData := dhcpStaticMappingRequestUpdate{
		DHCPStaticMappingRequest: mappingData,
		Id:                       idToUpdate,
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}
	response, err := s.client.put(ctx, staticMappingEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(createStaticMappingResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// DeleteStaticMapping deletes a DHCP static mapping.
func (s DHCPService) DeleteStaticMapping(ctx context.Context, mappingInterface string, idToDelete int) error {
	_, err := s.client.delete(
		ctx,
		staticMappingEndpoint,
		map[string]string{
			"interface": mappingInterface,
			"id":        strconv.Itoa(idToDelete),
		},
	)
	if err != nil {
		return err
	}
	return nil
}
