package pfsenseapi

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"time"
)

var defaultTimeout = 5 * time.Second

type Client struct {
	client *http.Client
	Cfg    Config

	DHCP *DHCPService
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

func (c *Client) do(ctx context.Context, method, endpoint string, body []byte) (*http.Response, error) {
	baseURL := fmt.Sprintf("%s/%s", c.Cfg.Host, endpoint)
	req, err := http.NewRequestWithContext(ctx, method, baseURL, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	if c.Cfg.User != "" && c.Cfg.Password != "" {
		req.SetBasicAuth(c.Cfg.User, c.Cfg.Password)
	}

	if c.Cfg.ApiClientID != "" && c.Cfg.ApiClientToken != "" {
		req.Header.Add("Authorization", fmt.Sprintf("%s %s", c.Cfg.ApiClientID, c.Cfg.ApiClientToken))
	}

	return c.client.Do(req)
}

func (c *Client) get(ctx context.Context, endpoint string) ([]byte, error) {
	res, err := c.do(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = res.Body.Close()
	}()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *Client) post(ctx context.Context, endpoint string, body []byte) ([]byte, error) {
	res, err := c.do(ctx, http.MethodPost, endpoint, body)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = res.Body.Close()
	}()

	respbody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return respbody, nil
}

type apiResponse struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Return  int    `json:"return"`
	Message string `json:"message"`
}
