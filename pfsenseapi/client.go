package pfsenseapi

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var defaultTimeout = 5 * time.Second

// Client provides client Methods
type Client struct {
	client *http.Client
	Cfg    Config

	DHCP      *DHCPService
	Status    *StatusService
	Interface *InterfaceService
}

// Config provides configuration for the client. These values are only read in
// when NewClient is called.
type Config struct {
	Host string

	User     string
	Password string

	ApiClientID    string
	ApiClientToken string

	SkipTLS bool
	Timeout time.Duration
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
	newClient.DHCP = &DHCPService{client: newClient}
	newClient.Status = &StatusService{client: newClient}
	newClient.Interface = &InterfaceService{client: newClient}
	return newClient
}

// NewClientWithNoAuth constructs a new Client using defaults for everything
// except the host
func NewClientWithNoAuth(host string) *Client {
	config := Config{
		Host:    host,
		SkipTLS: false,
		Timeout: defaultTimeout,
	}
	return NewClient(config)
}

// NewClientFromLocalAuth constructs a new Client using Local username/password
// authentication
func NewClientFromLocalAuth(host, user, password string) *Client {
	config := Config{
		Host:     host,
		User:     user,
		Password: password,
		SkipTLS:  false,
		Timeout:  defaultTimeout,
	}

	return NewClient(config)
}

// NewClientFromTokenAuth constructs a new Client using token authentication
func NewClientFromTokenAuth(host, apiClientID, apiClientToken string) *Client {
	config := Config{
		Host:           host,
		ApiClientID:    apiClientID,
		ApiClientToken: apiClientToken,
		SkipTLS:        false,
		Timeout:        defaultTimeout,
	}
	return NewClient(config)
}

type service struct {
	client *Client
}

func (c *Client) do(ctx context.Context, method, endpoint string, queryMap map[string]string, body []byte) (*http.Response, error) {
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
	if c.Cfg.User != "" && c.Cfg.Password != "" {
		req.SetBasicAuth(c.Cfg.User, c.Cfg.Password)
	}

	if c.Cfg.ApiClientID != "" && c.Cfg.ApiClientToken != "" {
		req.Header.Add("Authorization", fmt.Sprintf("%s %s", c.Cfg.ApiClientID, c.Cfg.ApiClientToken))
	}

	return c.client.Do(req)
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
