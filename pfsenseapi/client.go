package pfsenseapi

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	host     string
	client   *http.Client
	user     string
	password string

	DHCP *DHCPService
}

type service struct {
	client *Client
}

func NewClient(host, user, password string, timeout time.Duration) *Client {
	client := &http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	newClient := &Client{
		host:     host,
		client:   client,
		user:     user,
		password: password,
	}
	newClient.DHCP = &DHCPService{client: newClient}
	return newClient
}

func (c *Client) do(ctx context.Context, method, endpoint string) (*http.Response, error) {
	baseURL := fmt.Sprintf("%s/%s", c.host, endpoint)
	req, err := http.NewRequestWithContext(ctx, method, baseURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	if c.user != "" && c.password != "" {
		req.SetBasicAuth(c.user, c.password)
	}
	return c.client.Do(req)
}

func (c *Client) get(ctx context.Context, endpoint string) ([]byte, error) {
	res, err := c.do(ctx, http.MethodGet, endpoint)
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

type apiResponse struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Return  int    `json:"return"`
	Message string `json:"message"`
}
