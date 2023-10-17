package pfsenseapi

import (
	"context"
	"encoding/json"
)

const (
	userEndpoint        = "api/v1/user"
	groupEndpoint       = "api/v1/user/group"
	groupMemberEndpoint = "api/v1/user/group/member"
	privilegeEndpoint   = "api/v1/user/privilege"
)

// UserService provides User API methods
type UserService service

// User represents a single user.
type User struct {
	Scope            string   `json:"scope"`
	BcryptHash       string   `json:"bcrypt-hash"`
	Descr            string   `json:"descr"`
	Name             string   `json:"name"`
	Expires          string   `json:"expires"`
	Dashboardcolumns string   `json:"dashboardcolumns"`
	AuthorizedKeys   string   `json:"authorizedkeys"`
	Ipsecpsk         string   `json:"ipsecpsk"`
	Webguicss        string   `json:"webguicss"`
	Cert             []string `json:"cert"`
	Uid              string   `json:"uid"`
	GroupName        string   `json:"groupname"`
	Priv             []string `json:"priv"`
}

type userResponse struct {
	apiResponse
	Data []*User `json:"data"`
}

// ListUsers returns a list of the users.
func (s UserService) ListUsers(ctx context.Context) ([]*User, error) {
	response, err := s.client.get(ctx, userEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(userResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// DeleteUser deletes a user by username.
func (s UserService) DeleteUser(ctx context.Context, username string) error {
	_, err := s.client.delete(
		ctx,
		userEndpoint,
		map[string]string{
			"username": username,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

type UserRequest struct {
	AuthorizedKeys string   `json:"authorizedkeys"`
	Cert           []string `json:"cert"`
	Descr          string   `json:"descr"`
	Disabled       bool     `json:"disabled"`
	Expires        string   `json:"expires"`
	Ipsecpsk       string   `json:"ipsecpsk"`
	Password       string   `json:"password"`
	Priv           []string `json:"priv"`
	Username       string   `json:"username"`
}

type createUserResponse struct {
	apiResponse
	Data *User `json:"data"`
}

// CreateUser creates a new User.
func (s UserService) CreateUser(
	ctx context.Context,
	newUser UserRequest,
) (*User, error) {
	jsonData, err := json.Marshal(newUser)
	if err != nil {
		return nil, err
	}

	response, err := s.client.post(ctx, userEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(createUserResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// UpdateUser updates a user.
func (s UserService) UpdateUser(
	ctx context.Context,
	userData UserRequest,
) (*User, error) {
	jsonData, err := json.Marshal(userData)
	if err != nil {
		return nil, err
	}

	response, err := s.client.put(ctx, userEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(createUserResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// Group represents a single user group.
type Group struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Scope       string   `json:"scope"`
	Gid         int      `json:"gid"`
	Member      []int    `json:"member"`
	Priv        []string `json:"priv"`
}

type groupResponse struct {
	apiResponse
	Data []*Group `json:"data"`
}

// ListGroups returns a list of the groups.
func (s UserService) ListGroups(ctx context.Context) ([]*Group, error) {
	response, err := s.client.get(ctx, groupEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(groupResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// DeleteGroup deletes a group by group name.
func (s UserService) DeleteGroup(ctx context.Context, groupname string) error {
	_, err := s.client.delete(
		ctx,
		groupEndpoint,
		map[string]string{
			"id": groupname,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

type GroupRequest struct {
	Name        string   `json:"name"`
	Scope       string   `json:"scope"`
	Description string   `json:"description"`
	Member      []int    `json:"member"`
	Priv        []string `json:"priv"`
}

type createGroupResponse struct {
	apiResponse
	Data *Group `json:"data"`
}

// CreateGroup creates a new Group.
func (s UserService) CreateGroup(
	ctx context.Context,
	newGroup GroupRequest,
) (*Group, error) {
	jsonData, err := json.Marshal(newGroup)
	if err != nil {
		return nil, err
	}

	response, err := s.client.post(ctx, groupEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(createGroupResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

type groupRequestUpdate struct {
	GroupRequest
	Id string `json:"id"`
}

// UpdateGroup modifies an existing group.
func (s UserService) UpdateGroup(
	ctx context.Context,
	groupToUpdate string,
	newGroupData GroupRequest,
) (*Group, error) {
	requestData := groupRequestUpdate{
		GroupRequest: newGroupData,
		Id:           groupToUpdate,
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}
	response, err := s.client.put(ctx, groupEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(createGroupResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

type userGroupsRequest struct {
	Group    []string `json:"group"`
	Username string   `json:"username"`
}

// AddUserToGroups adds a user to existing groups.
func (s UserService) AddUserToGroups(ctx context.Context, username string, groups []string) error {
	requestBody := userGroupsRequest{
		Group:    groups,
		Username: username,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}

	if _, err := s.client.post(ctx, groupMemberEndpoint, nil, jsonData); err != nil {
		return err
	}
	return nil
}

// RemoveUserFromGroup removes a user from a group.
func (s UserService) RemoveUserFromGroup(ctx context.Context, username, groupname string) error {
	_, err := s.client.delete(
		ctx,
		groupMemberEndpoint,
		map[string]string{
			"username": username,
			"group":    groupname,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

// RemovePrivilegeFromUser removes a privilege from a user.
func (s UserService) RemovePrivilegeFromUser(ctx context.Context, username, privname string) error {
	_, err := s.client.delete(
		ctx,
		privilegeEndpoint,
		map[string]string{
			"username": username,
			"priv":     privname,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

type userPrivRequest struct {
	Priv     []string `json:"priv"`
	Username string   `json:"username"`
}

// AddPrivilegesToUser adds existing privileges to a user.
func (s UserService) AddPrivilegesToUser(ctx context.Context, username string, privileges []string) error {
	requestBody := userPrivRequest{
		Priv:     privileges,
		Username: username,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}

	if _, err := s.client.post(ctx, privilegeEndpoint, nil, jsonData); err != nil {
		return err
	}
	return nil
}
