package pfsenseapi

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/exp/slices"
	"io"
	"net/http"
	"time"
)

var (
	defaultTimeout = 5 * time.Second

	// noAuthEndpoints is a list of endpoints that require no authentication
	noAuthEndpoints = []string{
		apiErrorEndpoint,
	}

	// localAuthEndpoints is a list of endpoints that always require local
	// authentication. This overrides the default behavior of authenticating with
	// whatever client the Client is constructed with.
	localAuthEndpoints = []string{
		tokenEndpoint,
	}
)

// Client provides client Methods
type Client struct {
	client *http.Client
	Cfg    Config

	System    *SystemService
	Token     *TokenService
	DHCP      *DHCPService
	Status    *StatusService
	Interface *InterfaceService
	Routing   *RoutingService
	Firewall  *FirewallService
	User      *UserService
}

// Config provides configuration for the client. These values are only read in
// when NewClient is called.
type Config struct {
	Host string

	LocalAuthEnabled bool
	User             string
	Password         string

	JWTAuthEnabled bool
	JWTToken       string

	TokenAuthEnabled bool
	ApiClientID      string
	ApiClientToken   string

	SkipTLS bool
	Timeout time.Duration
}

// authEnabled returns true if any authentication mechanism is enabled, or false
// if this is a NoAuth client.
func (c Config) authEnabled() bool {
	if !c.LocalAuthEnabled && !c.TokenAuthEnabled && !c.JWTAuthEnabled {
		return false
	}
	return true
}

// NewClient constructs a new Client
func NewClient(config Config) *Client {
	httpclient := &http.Client{
		Timeout: config.Timeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: config.SkipTLS},
		},
	}

	newClient := &Client{
		Cfg:    config,
		client: httpclient,
	}
	newClient.System = &SystemService{client: newClient}
	newClient.Token = &TokenService{client: newClient}
	newClient.DHCP = &DHCPService{client: newClient}
	newClient.Status = &StatusService{client: newClient}
	newClient.Interface = &InterfaceService{client: newClient}
	newClient.Routing = &RoutingService{client: newClient}
	newClient.Firewall = &FirewallService{client: newClient}
	newClient.User = &UserService{client: newClient}
	return newClient
}

// NewClientWithNoAuth constructs a new Client using defaults for everything
// except the host
func NewClientWithNoAuth(host string) *Client {
	config := Config{
		Host:    host,
		SkipTLS: true,
		Timeout: defaultTimeout,
	}

	return NewClient(config)
}

// NewClientWithLocalAuth constructs a new Client using Local username/password
// authentication
func NewClientWithLocalAuth(host, user, password string) *Client {
	config := Config{
		Host:             host,
		User:             user,
		Password:         password,
		SkipTLS:          true,
		Timeout:          defaultTimeout,
		LocalAuthEnabled: true,
	}

	return NewClient(config)
}

// NewClientWithJWTAuth constructs a new Client using JWT token authentication.
// The username and password provided here will be used to generate JWT tokens
// for authentication.
func NewClientWithJWTAuth(host, user, password string) *Client {
	config := Config{
		Host:           host,
		User:           user,
		JWTAuthEnabled: true,
		Password:       password,
		SkipTLS:        true,
		Timeout:        defaultTimeout,
	}

	return NewClient(config)
}

// NewClientWithTokenAuth constructs a new Client using token authentication
func NewClientWithTokenAuth(host, apiClientID, apiClientToken string) *Client {
	config := Config{
		Host:             host,
		ApiClientID:      apiClientID,
		ApiClientToken:   apiClientToken,
		SkipTLS:          true,
		Timeout:          defaultTimeout,
		TokenAuthEnabled: true,
	}
	return NewClient(config)
}

type service struct {
	client *Client
}

func (c *Client) do(ctx context.Context, method, endpoint string, queryMap map[string]string, body []byte) (*http.Response, error) {
	res, err := c.doRequest(ctx, method, endpoint, queryMap, body)
	if err != nil {
		return nil, err
	}

	// refresh token and try again if expired
	if c.Cfg.JWTAuthEnabled && res.StatusCode == 401 {
		if _, err = c.generateToken(ctx); err != nil {
			return nil, err
		}

		res, err = c.doRequest(ctx, method, endpoint, queryMap, body)
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}

func (c *Client) doRequest(ctx context.Context, method, endpoint string, queryMap map[string]string, body []byte) (*http.Response, error) {
	baseURL := fmt.Sprintf("%s/%s", c.Cfg.Host, endpoint)
	req, err := http.NewRequestWithContext(ctx, method, baseURL, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	for key, value := range queryMap {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	req.Header.Add("Accept", "application/json")

	req, err = configureAuthForRequest(ctx, req, c, endpoint)
	if err != nil {
		return nil, err
	}

	return c.client.Do(req)
}

func configureAuthForRequest(
	ctx context.Context,
	req *http.Request,
	c *Client,
	endpoint string,
) (*http.Request, error) {
	if !c.Cfg.authEnabled() {
		return req, nil
	}

	if slices.Contains(noAuthEndpoints, endpoint) {
		return req, nil
	}

	if slices.Contains(localAuthEndpoints, endpoint) {
		if c.Cfg.User == "" || c.Cfg.Password == "" {
			return nil, errors.New("endpoint requires local authentication, but no user/pass available in client")
		}

		req.SetBasicAuth(c.Cfg.User, c.Cfg.Password)
		return req, nil
	}

	switch {
	case c.Cfg.JWTAuthEnabled:
		token, err := c.getToken(ctx)
		if err != nil {
			return nil, err
		}
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	case c.Cfg.LocalAuthEnabled:
		req.SetBasicAuth(c.Cfg.User, c.Cfg.Password)
	case c.Cfg.TokenAuthEnabled:
		req.Header.Add("Authorization", fmt.Sprintf("%s %s", c.Cfg.ApiClientID, c.Cfg.ApiClientToken))
	}
	return req, nil
}

// getToken returns the token if already set, otherwise generates a new token
// prior to returning
func (c *Client) getToken(ctx context.Context) (string, error) {
	if c.Cfg.JWTToken != "" {
		return c.Cfg.JWTToken, nil
	}

	return c.generateToken(ctx)
}

// generateToken creates a new token and updates client
func (c *Client) generateToken(ctx context.Context) (string, error) {
	token, err := c.Token.CreateAccessToken(ctx)
	if err != nil {
		return "", err
	}
	c.Cfg.JWTToken = token
	return token, nil
}

func (c *Client) get(ctx context.Context, endpoint string, queryMap map[string]string) ([]byte, error) {
	res, err := c.do(ctx, http.MethodGet, endpoint, queryMap, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		_, _ = io.Copy(io.Discard, res.Body)
		_ = res.Body.Close()
	}()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		resp := new(apiResponse)
		if err = json.Unmarshal(body, resp); err != nil {
			return nil, fmt.Errorf("non 2xx response code received: %d", res.StatusCode)
		}
		return nil, fmt.Errorf("%s, response code %d", resp.Message, res.StatusCode)
	}

	return body, nil
}

func (c *Client) post(ctx context.Context, endpoint string, queryMap map[string]string, body []byte) ([]byte, error) {
	res, err := c.do(ctx, http.MethodPost, endpoint, queryMap, body)
	if err != nil {
		return nil, err
	}
	defer func() {
		_, _ = io.Copy(io.Discard, res.Body)
		_ = res.Body.Close()
	}()

	respbody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		resp := new(apiResponse)
		if err = json.Unmarshal(respbody, resp); err != nil {
			return nil, fmt.Errorf("non 2xx response code received: %d", res.StatusCode)
		}
		return nil, fmt.Errorf("%s, response code %d", resp.Message, res.StatusCode)
	}

	return respbody, nil
}

func (c *Client) put(ctx context.Context, endpoint string, queryMap map[string]string, body []byte) ([]byte, error) {
	res, err := c.do(ctx, http.MethodPut, endpoint, queryMap, body)
	if err != nil {
		return nil, err
	}
	defer func() {
		_, _ = io.Copy(io.Discard, res.Body)
		_ = res.Body.Close()
	}()

	respbody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		resp := new(apiResponse)
		if err = json.Unmarshal(respbody, resp); err != nil {
			return nil, fmt.Errorf("non 2xx response code received: %d", res.StatusCode)
		}
		return nil, fmt.Errorf("%s, response code %d", resp.Message, res.StatusCode)
	}

	return respbody, nil
}

func (c *Client) delete(ctx context.Context, endpoint string, queryMap map[string]string) ([]byte, error) {
	res, err := c.do(ctx, http.MethodDelete, endpoint, queryMap, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		_, _ = io.Copy(io.Discard, res.Body)
		_ = res.Body.Close()
	}()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		resp := new(apiResponse)
		if err = json.Unmarshal(body, resp); err != nil {
			return nil, fmt.Errorf("non 2xx response code received: %d", res.StatusCode)
		}
		return nil, fmt.Errorf("%s, response code %d", resp.Message, res.StatusCode)
	}

	return body, nil
}

type apiResponse struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Return  int    `json:"return"`
	Message string `json:"message"`
}
