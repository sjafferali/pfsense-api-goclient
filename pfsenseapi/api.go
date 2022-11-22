package pfsenseapi

import (
	"context"
	"encoding/json"
)

const (
	apiEndpoint        = "api/v1/system/api"
	apiVersionEndpoint = "api/v1/system/api/version"
	apiErrorEndpoint   = "api/v1/system/api/error"
)

// APIService provides API API methods
type APIService service

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
func (s APIService) GetAPIConfiguration(ctx context.Context) (*APIConfiguration, error) {
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
func (s APIService) UpdateAPIConfiguration(ctx context.Context, apiConfiguration APIConfigurationRequest) error {
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
func (s APIService) GetAPIVersion(ctx context.Context) (*APIVersion, error) {
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
func (s APIService) GetErrorDefinitions(ctx context.Context) (map[string]*ErrorDefinition, error) {
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
