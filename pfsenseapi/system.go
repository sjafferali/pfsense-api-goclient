package pfsenseapi

import (
	"context"
	"encoding/json"
	"strconv"
)

const (
	apiEndpoint               = "api/v1/system/api"
	apiVersionEndpoint        = "api/v1/system/api/version"
	apiErrorEndpoint          = "api/v1/system/api/error"
	arpEndpoint               = "api/v1/system/arp"
	caCertificatesEndpoint    = "api/v1/system/ca"
	certificateEndpoint       = "api/v1/system/certificate"
	dnsConfigurationEndpoint  = "api/v1/system/dns"
	dnsServerEndpoint         = "api/v1/system/dns/server"
	haltEndpoint              = "api/v1/system/halt"
	hostnameEndpoint          = "api/v1/system/hostname"
	rebootEndpoint            = "api/v1/system/reboot"
	emailNotificationEndpoint = "api/v1/system/notifications/email"
	packageEndpoint           = "api/v1/system/package"
	tunableEndpoint           = "api/v1/system/tunable"
	versionEndpoint           = "api/v1/system/version"
	versionUpgradeEndpoint    = "api/v1/system/version/upgrade"
)

// SystemService provides System API methods
type SystemService service

// APIConfiguration represents the API configuration
type APIConfiguration struct {
	Enable            string `json:"enable"`
	Persist           string `json:"persist"`
	AllowedInterfaces string `json:"allowed_interfaces"`
	AuthMode          string `json:"authmode"`
	ContentType       string `json:"content_type"`
	JwtExp            string `json:"jwt_exp"`
	Keyhash           string `json:"keyhash"`
	Keybytes          string `json:"keybytes"`
	Keys              string `json:"keys"`
	AccessList        string `json:"access_list"`
}

type apiConfigurationResponse struct {
	apiResponse
	Data *APIConfiguration `json:"data"`
}

// GetAPIConfiguration returns the API configuration
func (s SystemService) GetAPIConfiguration(ctx context.Context) (*APIConfiguration, error) {
	response, err := s.client.get(ctx, apiEndpoint, nil)
	if err != nil {
		return nil, err
	}
	resp := new(apiConfigurationResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// APIConfigurationRequest is the request used to update the API configuration
type APIConfigurationRequest struct {
	AccessList            []string            `json:"access_list"`
	AllowOptions          bool                `json:"allow_options"`
	AuthMode              string              `json:"authmode"`
	AllowedInterfaces     []string            `json:"allowed_interfaces"`
	CustomHeaders         []map[string]string `json:"custom_headers"`
	Enable                bool                `json:"enable"`
	EnableLoginProtection bool                `json:"enable_login_protection"`
	LogSuccessfulAuth     bool                `json:"log_successful_auth"`
	Hasync                bool                `json:"hasync"`
	HasyncHosts           []string            `json:"hasync_hosts"`
	HasyncPassword        string              `json:"hasync_password"`
	HasyncUsername        string              `json:"hasync_username"`
	JwtExp                int                 `json:"jwt_exp"`
	Keybytes              int                 `json:"keybytes"`
	Keyhash               string              `json:"keyhash"`
	Persist               bool                `json:"persist"`
	Readonly              bool                `json:"readonly"`
}

// UpdateAPIConfiguration updates the API configuration
func (s SystemService) UpdateAPIConfiguration(ctx context.Context, apiConfiguration APIConfigurationRequest) error {
	jsonData, err := json.Marshal(apiConfiguration)
	if err != nil {
		return err
	}
	_, err = s.client.put(ctx, apiEndpoint, nil, jsonData)
	if err != nil {
		return err
	}
	return nil
}

// APIVersion represents the API Versions.
type APIVersion struct {
	CurrentVersion  string `json:"current_version"`
	LatestVersion   string `json:"latest_version"`
	UpdateAvailable bool   `json:"update_available"`
}

type apiVersionResponse struct {
	apiResponse
	Data *APIVersion `json:"data"`
}

// GetAPIVersion returns the API versions
func (s SystemService) GetAPIVersion(ctx context.Context) (*APIVersion, error) {
	response, err := s.client.get(ctx, apiVersionEndpoint, nil)
	if err != nil {
		return nil, err
	}
	resp := new(apiVersionResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// ErrorDefinition represents a single error definition.
type ErrorDefinition struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Return  int    `json:"return"`
	Message string `json:"message"`
}

type errorDefinitionsResponse struct {
	apiResponse
	Data map[string]*ErrorDefinition `json:"data"`
}

// GetErrorDefinitions returns a map with the error code being the key and value
// being the error definition.
func (s SystemService) GetErrorDefinitions(ctx context.Context) (map[string]*ErrorDefinition, error) {
	response, err := s.client.get(ctx, apiErrorEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(errorDefinitionsResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

// ArpEntry represents a single arp entry in the arp table.
type ArpEntry struct {
	Ip        string `json:"ip"`
	Mac       string `json:"mac"`
	Interface string `json:"interface"`
	Status    string `json:"status"`
	Linktype  string `json:"linktype"`
}

type arpEntriesResponse struct {
	apiResponse
	Data []*ArpEntry `json:"data"`
}

// ListArpTable returns all the arp entries in the arp table.
func (s SystemService) ListArpTable(ctx context.Context) ([]*ArpEntry, error) {
	response, err := s.client.get(ctx, arpEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(arpEntriesResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

// DeleteArpEntry deletes a arp entry for an address from the arp table
func (s SystemService) DeleteArpEntry(ctx context.Context, address string) error {
	_, err := s.client.delete(
		ctx,
		arpEndpoint,
		map[string]string{
			"ip": address,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

// CACertificate represents a single CACertificate.
type CACertificate struct {
	Refid        string `json:"refid"`
	Descr        string `json:"descr"`
	Trust        string `json:"trust"`
	Randomserial string `json:"randomserial"`
	Crt          string `json:"crt"`
	Prv          string `json:"prv"`
	Serial       string `json:"serial"`
}

type caCertificatesResponse struct {
	apiResponse
	Data struct {
		CA []*CACertificate `json:"ca"`
	} `json:"data"`
}

// ListCACertificates returns all the CA certificates installed on the system.
func (s SystemService) ListCACertificates(ctx context.Context) ([]*CACertificate, error) {
	response, err := s.client.get(ctx, caCertificatesEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(caCertificatesResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}

	return resp.Data.CA, nil
}

// DeleteCACertificate deletes a CA certificate installed on the system.
func (s SystemService) DeleteCACertificate(ctx context.Context, refid string) error {
	_, err := s.client.delete(
		ctx,
		caCertificatesEndpoint,
		map[string]string{
			"refid": refid,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

// CACertificateRequest represents a single CA Certificate. This type is used for
// creating a new CA certificate.
type CACertificateRequest struct {
	Caref                string `json:"caref"`
	Crt                  string `json:"crt"`
	Descr                string `json:"descr"`
	DigestAlg            string `json:"digest_alg"`
	DnCity               string `json:"dn_city"`
	DnCommonname         string `json:"dn_commonname"`
	DnCountry            string `json:"dn_country"`
	DnOrganization       string `json:"dn_organization"`
	DnOrganizationalunit string `json:"dn_organizationalunit"`
	DnState              string `json:"dn_state"`
	Ecname               string `json:"ecname"`
	Keylen               int    `json:"keylen"`
	Keytype              string `json:"keytype"`
	Lifetime             int    `json:"lifetime"`
	Method               string `json:"method"`
	Prv                  string `json:"prv"`
	RandomSerial         bool   `json:"randomserial"`
	Serial               int    `json:"serial"`
	Trust                bool   `json:"trust"`
}

type createCACertificateResponse struct {
	apiResponse
	Data *CACertificate `json:"data"`
}

// CreateCACertificate generate or import new CA certificate.
func (s SystemService) CreateCACertificate(
	ctx context.Context,
	newCACertificate CACertificateRequest,
) (*CACertificate, error) {
	jsonData, err := json.Marshal(newCACertificate)
	if err != nil {
		return nil, err
	}
	response, err := s.client.post(ctx, caCertificatesEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(createCACertificateResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// Certificate represents a single installed SSL/TLS certificate.
type Certificate struct {
	Refid string `json:"refid"`
	Descr string `json:"descr"`
	Prv   string `json:"prv"`
	Crt   string `json:"crt"`
	Caref string `json:"caref"`
}

type certificatesResponse struct {
	apiResponse
	Data struct {
		Cert []*Certificate `json:"cert"`
	} `json:"data"`
}

// ListCertificates returns all the SSL/TLS certificates installed on the system.
func (s SystemService) ListCertificates(ctx context.Context) ([]*Certificate, error) {
	response, err := s.client.get(ctx, certificateEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(certificatesResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}

	return resp.Data.Cert, nil
}

// DeleteCertificate deletes a SSL/TLS certificate installed on the system.
func (s SystemService) DeleteCertificate(ctx context.Context, refid string) error {
	_, err := s.client.delete(
		ctx,
		certificateEndpoint,
		map[string]string{
			"refid": refid,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

// CertificateCreateRequest represents a single Certificate. This type is used to
// create a new certificate.
type CertificateCreateRequest struct {
	Active   bool `json:"active"`
	Altnames []struct {
		DNS   string `json:"dns,omitempty"`
		IP    string `json:"ip,omitempty"`
		URI   string `json:"uri,omitempty"`
		Email string `json:"email,omitempty"`
	} `json:"altnames"`
	Caref                string `json:"caref"`
	Crt                  string `json:"crt"`
	Descr                string `json:"descr"`
	DigestAlg            string `json:"digest_alg"`
	DnCity               string `json:"dn_city"`
	DnCommonname         string `json:"dn_commonname"`
	DnCountry            string `json:"dn_country"`
	DnOrganization       string `json:"dn_organization"`
	DnOrganizationalunit string `json:"dn_organizationalunit"`
	DnState              string `json:"dn_state"`
	Ecname               string `json:"ecname"`
	Keylen               int    `json:"keylen"`
	Keytype              string `json:"keytype"`
	Lifetime             int    `json:"lifetime"`
	Method               string `json:"method"`
	Prv                  string `json:"prv"`
	Type                 string `json:"type"`
}

type createCertificateResponse struct {
	apiResponse
	Data *Certificate `json:"data"`
}

// CreateCertificate generate or import new certificate.
func (s SystemService) CreateCertificate(
	ctx context.Context,
	newCertificate CertificateCreateRequest,
) (*Certificate, error) {
	jsonData, err := json.Marshal(newCertificate)
	if err != nil {
		return nil, err
	}
	response, err := s.client.post(ctx, certificateEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(createCertificateResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// CertificateUpdateRequest is used to update a certificate.
type CertificateUpdateRequest struct {
	Descr  string `json:"descr"`
	Prv    string `json:"prv"`
	Crt    string `json:"crt"`
	Active bool   `json:"active"`
}

type certificateUpdateRequest struct {
	CertificateUpdateRequest
	Refid string `json:"refid"`
}

// UpdateCertificate modifies an existing certificate
func (s SystemService) UpdateCertificate(
	ctx context.Context,
	refIDToUpdate string,
	newCertificateData CertificateUpdateRequest,
) (*Certificate, error) {
	requestData := certificateUpdateRequest{
		CertificateUpdateRequest: newCertificateData,
		Refid:                    refIDToUpdate,
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}
	response, err := s.client.put(ctx, certificateEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(createCertificateResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// DNSConfiguration represents the system DNS configuration.
type DNSConfiguration struct {
	Dnsserver        []string `json:"dnsserver"`
	Dnsallowoverride bool     `json:"dnsallowoverride"`
	Dnslocalhost     bool     `json:"dnslocalhost"`
}

type dnsConfigurationResponse struct {
	apiResponse
	Data *DNSConfiguration `json:"data"`
}

// GetDNSConfiguration returns the system DNS configuration.
func (s SystemService) GetDNSConfiguration(ctx context.Context) (*DNSConfiguration, error) {
	response, err := s.client.get(ctx, dnsConfigurationEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(dnsConfigurationResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

// UpdateDNSConfiguration updates the DNS configuration
func (s SystemService) UpdateDNSConfiguration(ctx context.Context, dnsConfiguration DNSConfiguration) error {
	jsonData, err := json.Marshal(dnsConfiguration)
	if err != nil {
		return err
	}
	_, err = s.client.put(ctx, dnsConfigurationEndpoint, nil, jsonData)
	if err != nil {
		return err
	}
	return nil
}

// DeleteDNSServer deletes a system DNS server.
func (s SystemService) DeleteDNSServer(ctx context.Context, dnsserver string) error {
	_, err := s.client.delete(
		ctx,
		dnsServerEndpoint,
		map[string]string{
			"dnsserver": dnsserver,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

type dnsServersRequest struct {
	Dnsserver []string `json:"dnsserver"`
}

// AddDNSServers adds new DNS servers to the system DNS configuration.
func (s SystemService) AddDNSServers(ctx context.Context, newDNSServers []string) error {
	requestData := dnsServersRequest{Dnsserver: newDNSServers}
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return err
	}
	_, err = s.client.post(ctx, dnsServerEndpoint, nil, jsonData)
	if err != nil {
		return err
	}
	return nil
}

// Halt shuts down the pfsense system.
func (s SystemService) Halt(ctx context.Context) error {
	_, err := s.client.post(ctx, haltEndpoint, nil, nil)
	if err != nil {
		return err
	}
	return nil
}

type SystemHostname struct {
	Hostname string `json:"hostname"`
	Domain   string `json:"domain"`
}

type hostnameResponse struct {
	apiResponse
	Data *SystemHostname `json:"data"`
}

// GetHostname returns the system hostname configuration.
func (s SystemService) GetHostname(ctx context.Context) (*SystemHostname, error) {
	response, err := s.client.get(ctx, hostnameEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(hostnameResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

// UpdateHostname updates the system hostname.
func (s SystemService) UpdateHostname(ctx context.Context, newHostname SystemHostname) error {
	jsonData, err := json.Marshal(newHostname)
	if err != nil {
		return err
	}

	if _, err = s.client.put(ctx, hostnameEndpoint, nil, jsonData); err != nil {
		return err
	}

	return nil
}

// Reboot initiates a system reboot.
func (s SystemService) Reboot(ctx context.Context) error {
	_, err := s.client.post(ctx, rebootEndpoint, nil, nil)
	if err != nil {
		return err
	}
	return nil
}

// EmailNotification represents the email notification configuration.
type EmailNotification struct {
	Ipaddress               string `json:"ipaddress"`
	Port                    string `json:"port"`
	Sslvalidate             string `json:"sslvalidate"`
	Timeout                 string `json:"timeout"`
	Notifyemailaddress      string `json:"notifyemailaddress"`
	Username                string `json:"username"`
	Password                string `json:"password"`
	AuthenticationMechanism string `json:"authentication_mechanism"`
	Fromaddress             string `json:"fromaddress"`
	Disable                 string `json:"disable"`
}

type emailNotificationResponse struct {
	apiResponse
	Data *EmailNotification `json:"data"`
}

// GetEmailNotification returns the system email notification configuration.
func (s SystemService) GetEmailNotification(ctx context.Context) (*EmailNotification, error) {
	response, err := s.client.get(ctx, emailNotificationEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(emailNotificationResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

type EmailNotificationRequest struct {
	AuthenticationMechanism string `json:"authentication_mechanism"`
	Disabled                bool   `json:"disabled"`
	FromAddress             string `json:"fromaddress"`
	Ipaddress               string `json:"ipaddress"`
	Notifyemailaddress      string `json:"notifyemailaddress"`
	Password                string `json:"password"`
	Port                    int    `json:"port"`
	Ssl                     bool   `json:"ssl"`
	SslValidate             bool   `json:"sslvalidate"`
	Timeout                 int    `json:"timeout"`
	Username                string `json:"username"`
}

// UpdateEmailNotification updates the system email notification configuration.
func (s SystemService) UpdateEmailNotification(ctx context.Context, newConfig EmailNotificationRequest) error {
	jsonData, err := json.Marshal(newConfig)
	if err != nil {
		return err
	}

	if _, err = s.client.put(ctx, emailNotificationEndpoint, nil, jsonData); err != nil {
		return err
	}

	return nil
}

// Package represents a single package.
type Package struct {
	Name             string `json:"name"`
	Version          string `json:"version"`
	InstalledVersion string `json:"installed_version"`
	Descr            string `json:"descr"`
	Installed        bool   `json:"installed"`
	UpdateAvailable  bool   `json:"update_available"`
}

type packageResponse struct {
	apiResponse
	Data []*Package `json:"data"`
}

// ListPackages returns a list of the packages. Passing the true to the all
// argument here includes all pfSense packages available, even packages that are
// not installed.
func (s SystemService) ListPackages(ctx context.Context, all bool) ([]*Package, error) {
	response, err := s.client.get(
		ctx,
		packageEndpoint,
		map[string]string{
			"all": strconv.FormatBool(all),
		},
	)
	if err != nil {
		return nil, err
	}

	resp := new(packageResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// UninstallPackage uninstalls a package.
func (s SystemService) UninstallPackage(ctx context.Context, name string) error {
	_, err := s.client.delete(
		ctx,
		packageEndpoint,
		map[string]string{
			"name": name,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

type newPackage struct {
	Name string `json:"name"`
}

// InstallPackage installs a new package on the system.
func (s SystemService) InstallPackage(ctx context.Context, name string) error {
	requestData := newPackage{Name: name}
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return err
	}
	_, err = s.client.post(ctx, packageEndpoint, nil, jsonData)
	if err != nil {
		return err
	}
	return nil
}

// Tunable represents a single system tunable.
type Tunable struct {
	Tunable  string `json:"tunable"`
	Value    string `json:"value"`
	Descr    string `json:"descr"`
	Modified bool   `json:"modified"`
}

type tunableResponse struct {
	apiResponse
	Data []*Tunable `json:"data"`
}

// ListTunables returns a list of the system tunables.
func (s SystemService) ListTunables(ctx context.Context) ([]*Tunable, error) {
	response, err := s.client.get(ctx, tunableEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(tunableResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// DeleteTunable deletes a system tunable.
func (s SystemService) DeleteTunable(ctx context.Context, tunableID int) error {
	_, err := s.client.delete(ctx, tunableEndpoint, map[string]string{"id": strconv.Itoa(tunableID)})
	if err != nil {
		return err
	}
	return nil
}

type TunableRequest struct {
	Descr   string `json:"descr"`
	Tunable string `json:"tunable"`
	Value   string `json:"value"`
}

type createTunableResponse struct {
	apiResponse
	Data *Tunable `json:"data"`
}

// CreateTunable creates a new system tunable.
func (s SystemService) CreateTunable(
	ctx context.Context,
	newTunable TunableRequest,
) (*Tunable, error) {
	jsonData, err := json.Marshal(newTunable)
	if err != nil {
		return nil, err
	}
	response, err := s.client.post(ctx, tunableEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(createTunableResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

type tunableRequestUpdate struct {
	TunableRequest
	Id string `json:"id"`
}

// UpdateTunable modifies an existing tunable.
func (s SystemService) UpdateTunable(
	ctx context.Context,
	tunableToUpdate string,
	newTunable TunableRequest,
) (*Tunable, error) {
	requestData := tunableRequestUpdate{
		TunableRequest: newTunable,
		Id:             tunableToUpdate,
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}

	response, err := s.client.put(ctx, tunableEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(createTunableResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// Version represents the system version.
type Version struct {
	Version    string `json:"version"`
	Base       string `json:"base"`
	Patch      string `json:"patch"`
	Buildtime  string `json:"buildtime"`
	Lastcommit string `json:"lastcommit"`
	Program    int    `json:"program"`
}

type versionResponse struct {
	apiResponse
	Data *Version `json:"data"`
}

// GetVersion returns the system version.
func (s SystemService) GetVersion(ctx context.Context) (*Version, error) {
	response, err := s.client.get(ctx, versionEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(versionResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

// VersionUpgradeStatus represents the system version upgrade status.
type VersionUpgradeStatus struct {
	Version           string `json:"version"`
	InstalledVersion  string `json:"installed_version"`
	PkgVersionCompare string `json:"pkg_version_compare"`
}

type versionUpgradeResponse struct {
	apiResponse
	Data *VersionUpgradeStatus `json:"data"`
}

// GetVersionUpgradeStatus checks if there is an version upgrade available, but
// does not perform the upgrade.
func (s SystemService) GetVersionUpgradeStatus(ctx context.Context, useCache bool) (*VersionUpgradeStatus, error) {
	response, err := s.client.get(
		ctx,
		versionUpgradeEndpoint,
		map[string]string{
			"use_cache": strconv.FormatBool(useCache),
		},
	)
	if err != nil {
		return nil, err
	}

	resp := new(versionUpgradeResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
