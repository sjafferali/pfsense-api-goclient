package pfsenseapi

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/markphelps/optional"
)

const (
	userEndpoint   = "api/v2/user"
	usersEndpoint  = "api/v2/users"
	groupEndpoint  = "api/v2/user/group"
	groupsEndpoint = "api/v2/user/groups"
)

// UserService provides user API methods
type UserService service

type User struct {
	UserRequest
	Id  int `json:"id"`
	UID int `json:"uid"`
}

type UserRequest struct {
	Name           string          `json:"name"`
	Password       string          `json:"password"`
	Scope          string          `json:"scope"`
	Priv           []string        `json:"priv"`
	Disabled       bool            `json:"disabled"`
	Descr          string          `json:"descr"`
	Expires        optional.String `json:"expires"`
	Cert           []string        `json:"cert"`
	AuthorizedKeys optional.String `json:"authorizedkeys"`
	IPSecPSK       optional.String `json:"ipsecpsk"`
}

// userListResponse is the response that contains multiple users.
type userListResponse struct {
	apiResponse
	Data []*User `json:"data"`
}

// ListUsers returns a list of users.
func (s *UserService) ListUsers(ctx context.Context) ([]*User, error) {
	response, err := s.client.get(ctx, usersEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(userListResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// userGetResponse is the response that contains a single user.
type userGetResponse struct {
	apiResponse
	Data *User `json:"data"`
}

// GetUser returns a user by id.
func (s *UserService) GetUser(ctx context.Context, id int) (*User, error) {
	response, err := s.client.get(
		ctx,
		userEndpoint,
		map[string]string{
			"id": strconv.Itoa(id),
		},
	)
	if err != nil {
		return nil, err
	}

	resp := new(userGetResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// CreateUser creates a new user.
func (s *UserService) CreateUser(ctx context.Context, newUser UserRequest) (*User, error) {
	jsonData, err := json.Marshal(newUser)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	response, err := s.client.post(ctx, userEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(userGetResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// UpdateUser updates a user.
func (s *UserService) UpdateUser(ctx context.Context, id int, updatedUser UserRequest) (*User, error) {
	requestData := User{
		UserRequest: updatedUser,
		Id:          id,
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	response, err := s.client.patch(ctx, userEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(userGetResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// DeleteUser deletes a user.
func (s *UserService) DeleteUser(ctx context.Context, id int) (*User, error) {
	response, err := s.client.delete(
		ctx,
		userEndpoint,
		map[string]string{
			"id": strconv.Itoa(id),
		},
	)
	if err != nil {
		return nil, err
	}

	resp := new(userGetResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

type UserGroup struct {
	UserGroupRequest
	Id  int `json:"id"`
	GID int `json:"gid"`
}

// userGroupListResponse is the response that contains multiple user groups.
type userGroupListResponse struct {
	apiResponse
	Data []*UserGroup `json:"data"`
}

// userGroupGetResponse is the response that contains a single user group.
type userGroupGetResponse struct {
	apiResponse
	Data *UserGroup `json:"data"`
}

// ListUserGroups returns a list of user groups.
func (s *UserService) ListUserGroups(ctx context.Context) ([]*UserGroup, error) {
	response, err := s.client.get(ctx, groupsEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(userGroupListResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// GetUserGroup returns a user group by id.
func (s *UserService) GetUserGroup(ctx context.Context, id int) (*UserGroup, error) {
	response, err := s.client.get(
		ctx,
		groupEndpoint,
		map[string]string{
			"id": strconv.Itoa(id),
		},
	)
	if err != nil {
		return nil, err
	}

	resp := new(userGroupGetResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

type UserGroupRequest struct {
	Name        string   `json:"name"`
	Scope       string   `json:"scope"`
	Description string   `json:"description"`
	Member      []string `json:"member"`
	Priv        []string `json:"priv"`
}

// CreateUserGroup creates a new user group.
func (s *UserService) CreateUserGroup(ctx context.Context, newUserGroup UserGroupRequest) (*UserGroup, error) {
	jsonData, err := json.Marshal(newUserGroup)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	response, err := s.client.post(ctx, groupEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(userGroupGetResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// UpdateUserGroup updates a user group.
func (s *UserService) UpdateUserGroup(ctx context.Context, id int, updatedUserGroup UserGroupRequest) (*UserGroup, error) {
	requestData := UserGroup{
		UserGroupRequest: updatedUserGroup,
		Id:               id,
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	response, err := s.client.patch(ctx, groupEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(userGroupGetResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// DeleteUserGroup deletes a user group.
func (s *UserService) DeleteUserGroup(ctx context.Context, id int) (*UserGroup, error) {
	response, err := s.client.delete(
		ctx,
		groupEndpoint,
		map[string]string{
			"id": strconv.Itoa(id),
		},
	)
	if err != nil {
		return nil, err
	}

	resp := new(userGroupGetResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// PutUserGroups replaces all user groups with the provided list.
func (s *UserService) PutUserGroups(ctx context.Context, userGroups []*UserGroupRequest) ([]*UserGroup, error) {
	jsonData, err := json.Marshal(userGroups)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	response, err := s.client.put(ctx, groupsEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(userGroupListResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}
