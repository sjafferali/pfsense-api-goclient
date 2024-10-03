package pfsenseapi

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
)

const (
	leasesEndpoint        = "api/v1/services/dhcpd/lease"
	staticMappingEndpoint = "api/v1/services/dhcpd/static_mapping"
	serverEndpoint        = "api/v1/services/dhcpd"
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
	ID                     int           `json:"id"`
	Mac                    string        `json:"mac"`
	Cid                    string        `json:"cid"`
	IPaddr                 string        `json:"ipaddr"`
	Hostname               string        `json:"hostname"`
	Descr                  string        `json:"descr"`
	Filename               string        `json:"filename"`
	Rootpath               string        `json:"rootpath"`
	DefaultLeaseTime       string        `json:"defaultleasetime"`
	MaxLeaseTime           string        `json:"maxleasetime"`
	Gateway                string        `json:"gateway"`
	Domain                 string        `json:"domain"`
	DomainSearchList       string        `json:"domainsearchlist"`
	DDNSDomain             string        `json:"ddnsdomain"`
	DDNSDomainPrimary      string        `json:"ddnsdomainprimary"`
	DDNSDomainSecondary    string        `json:"ddnsdomainsecondary"`
	DDNSDomainkeyName      string        `json:"ddnsdomainkeyname"`
	DDNSDomainkeyAlgorithm string        `json:"ddnsdomainkeyalgorithm"`
	DDNSDomainkey          string        `json:"ddnsdomainkey"`
	DNSServers             []string      `json:"dnsserver"`
	TFTP                   string        `json:"tftp"`
	LDAP                   string        `json:"ldap"`
	NextServer             string        `json:"nextserver"`
	Filename32             string        `json:"filename32"`
	Filename64             string        `json:"filename64"`
	Filename32Arm          string        `json:"filename32arm"`
	Filename64Arm          string        `json:"filename64arm"`
	NumberOptions          string        `json:"numberoptions"`
	ArpTableStaticEntry    TrueIfPresent `json:"arp_table_static_entry"`
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

func (s DHCPService) getStaticMappingObjectId(ctx context.Context, mappingInterface string, macAddress string) (int, error) {
	mappings, err := s.ListStaticMappings(ctx, mappingInterface)

	if err != nil {
		return 0, err
	}

	for i, mapping := range mappings {
		if mapping.Mac == macAddress {
			return i, nil
		}
	}

	return 0, fmt.Errorf("Unable to find static mapping on interface %s with mac %s", mappingInterface, macAddress)
}

// UpdateStaticMapping modifies a DHCP static mapping.
func (s DHCPService) UpdateStaticMapping(
	ctx context.Context,
	macAddress string,
	mappingData DHCPStaticMappingRequest,
) (*DHCPStaticMapping, error) {
	id, err := s.getStaticMappingObjectId(ctx, mappingData.Interface, macAddress)

	if err != nil {
		return nil, err
	}

	requestData := dhcpStaticMappingRequestUpdate{
		DHCPStaticMappingRequest: mappingData,
		Id:                       id,
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
func (s DHCPService) DeleteStaticMapping(ctx context.Context, mappingInterface string, macAddress string) error {
	id, err := s.getStaticMappingObjectId(ctx, mappingInterface, macAddress)

	if err != nil {
		return err
	}

	_, err = s.client.delete(
		ctx,
		staticMappingEndpoint,
		map[string]string{
			"interface": mappingInterface,
			"id":        strconv.Itoa(id),
		},
	)
	if err != nil {
		return err
	}
	return nil
}

// DHCPServerConfigurationRequest updates the current DHCP Server (dhcpd) configuration for a specified interface
type DHCPServerConfigurationRequest struct {
	DefaultLeaseTime *int          `json:"defaultleasetime"`
	DenyUnknown      bool          `json:"denyunknown"`
	DNSServer        []string      `json:"dnsserver,omitempty"`
	Domain           string        `json:"domain,omitempty"`
	DomainSearchList []string      `json:"domainsearchlist,omitempty"`
	Enable           bool          `json:"enable"`
	Gateway          string        `json:"gateway,omitempty"`
	IgnoreBootP      bool          `json:"ignorebootp,omitempty"`
	Interface        string        `json:"interface"`
	MacAllow         []string      `json:"mac_allow,omitempty"`
	MacDeny          []string      `json:"mac_deny,omitempty"`
	MaxLeaseTime     *int          `json:"maxleasetime,omitempty"`
	NumberOptions    []interface{} `json:"numberoptions,omitempty"`
	RangeFrom        string        `json:"range_from,omitempty"`
	RangeTo          string        `json:"range_to,omitempty"`
	StaticARP        bool          `json:"staticarp"`
}

type DHCPRange struct {
	From string `json:"from"`
	To   string `json:"to"`
}

// DHCPServerConfiguration describes the current DHCP Server (dhcpd) configuration for a specified interface
type DHCPServerConfiguration struct {
	DefaultLeaseTime OptionalJSONInt `json:"defaultleasetime"`
	DenyUnknown      TrueIfPresent   `json:"denyunknown"`
	DNSServer        []string        `json:"dnsserver"`
	Domain           string          `json:"domain"`
	DomainSearchList string          `json:"domainsearchlist"`
	Enable           TrueIfPresent   `json:"enable"`
	Gateway          string          `json:"gateway"`
	IgnoreBootP      bool            `json:"ignorebootp"`
	Interface        string          `json:"interface"`
	MacAllow         string          `json:"mac_allow"`
	MacDeny          string          `json:"mac_deny"`
	MaxLeaseTime     OptionalJSONInt `json:"maxleasetime"`
	NumberOptions    string          `json:"numberoptions"`
	Range            *DHCPRange      `json:"range"`
	StaticARP        TrueIfPresent   `json:"staticarp"`
}

type dhcpServerResponse struct {
	apiResponse
	Data []*DHCPServerConfiguration `json:"data"`
}

// ListServerConfigurations lists all DHCP server configurations
func (s DHCPService) ListServerConfigurations(ctx context.Context) ([]*DHCPServerConfiguration, error) {
	response, err := s.client.get(ctx, serverEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(dhcpServerResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

type dhcpServerUpdateResponse struct {
	apiResponse
	Data *DHCPServerConfiguration `json:"data"`
}

// UpdateServerConfiguration modifies a DHCP server configuration.
func (s DHCPService) UpdateServerConfiguration(
	ctx context.Context,
	dhcpConfigData DHCPServerConfigurationRequest,
) (*DHCPServerConfiguration, error) {
	jsonData, err := json.Marshal(dhcpConfigData)
	if err != nil {
		return nil, err
	}
	response, err := s.client.put(ctx, serverEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(dhcpServerUpdateResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}

	resp.Data.Interface = dhcpConfigData.Interface
	return resp.Data, nil
}
