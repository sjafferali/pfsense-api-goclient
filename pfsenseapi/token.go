package pfsenseapi

import (
	"context"
	"encoding/json"
)

const (
	tokenEndpoint = "api/v1/access_token"
)

// TokenService provides Token API methods
type TokenService service

// accessToken represents a single API access token
type accessToken struct {
	Token string `json:"token"`
}

type accessTokenResponse struct {
	apiResponse
	Data *accessToken `json:"data"`
}

// CreateAccessToken creates a new AccessToken
func (s TokenService) CreateAccessToken(ctx context.Context) (string, error) {
	response, err := s.client.post(ctx, tokenEndpoint, nil, nil)
	if err != nil {
		return "", err
	}
	resp := new(accessTokenResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return "", err
	}
	return resp.Data.Token, nil
}
