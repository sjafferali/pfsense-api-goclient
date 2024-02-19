package pfsenseapi

import (
	"context"
	"testing"

	"github.com/markphelps/optional"
	"github.com/stretchr/testify/require"
)

func TestUserService_ListUsers(t *testing.T) {
	data := mustReadFileString(t, "testdata/multipleuser.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	users, err := newClient.User.ListUsers(context.Background())
	require.NoError(t, err)
	require.Len(t, users, 2)

	users, err = newClient.User.ListUsers(context.Background())
	require.Error(t, err)
	require.Nil(t, users)

	users, err = newClient.User.ListUsers(context.Background())
	require.Error(t, err)
	require.Nil(t, users)
}

func TestUserService_GetUser(t *testing.T) {
	data := mustReadFileString(t, "testdata/singleuser.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	user, err := newClient.User.GetUser(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, user)

	user, err = newClient.User.GetUser(context.Background(), 1)
	require.Error(t, err)
	require.Nil(t, user)

	user, err = newClient.User.GetUser(context.Background(), 1)
	require.Error(t, err)
	require.Nil(t, user)
}

func TestUserService_ListUserGroups(t *testing.T) {
	data := mustReadFileString(t, "testdata/multipleusergroup.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	userGroups, err := newClient.User.ListUserGroups(context.Background())
	require.NoError(t, err)
	require.Len(t, userGroups, 2)

	userGroups, err = newClient.User.ListUserGroups(context.Background())
	require.Error(t, err)
	require.Nil(t, userGroups)

	userGroups, err = newClient.User.ListUserGroups(context.Background())
	require.Error(t, err)
	require.Nil(t, userGroups)
}

func TestUserService_GetUserGroup(t *testing.T) {
	data := mustReadFileString(t, "testdata/singleusergroup.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	userGroup, err := newClient.User.GetUserGroup(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, userGroup)

	userGroup, err = newClient.User.GetUserGroup(context.Background(), 1)
	require.Error(t, err)
	require.Nil(t, userGroup)

	userGroup, err = newClient.User.GetUserGroup(context.Background(), 1)
	require.Error(t, err)
	require.Nil(t, userGroup)
}

func TestUserService_CreateUser(t *testing.T) {
	data := mustReadFileString(t, "testdata/singleuser.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	newUser := UserRequest{
		Name:           "newuser",
		Password:       "newpassword",
		Scope:          "user",
		Priv:           []string{"priv1", "priv2"},
		Disabled:       false,
		Descr:          "New User",
		Expires:        optional.NewString(""),
		Cert:           []string{"cert1", "cert2"},
		AuthorizedKeys: optional.NewString("key1"),
		IPSecPSK:       optional.NewString("psk1"),
	}
	user, err := newClient.User.CreateUser(context.Background(), newUser)
	require.NoError(t, err)
	require.NotNil(t, user)

	user, err = newClient.User.CreateUser(context.Background(), newUser)
	require.Error(t, err)
	require.Nil(t, user)

	user, err = newClient.User.CreateUser(context.Background(), newUser)
	require.Error(t, err)
	require.Nil(t, user)
}

func TestUserService_UpdateUser(t *testing.T) {
	data := mustReadFileString(t, "testdata/singleuser.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	updatedUser := UserRequest{
		Name:           "updateduser",
		Password:       "updatedpassword",
		Scope:          "user",
		Priv:           []string{"priv1", "priv2"},
		Disabled:       false,
		Descr:          "New User",
		Expires:        optional.NewString(""),
		Cert:           []string{"cert1", "cert2"},
		AuthorizedKeys: optional.NewString("key1"),
		IPSecPSK:       optional.NewString("psk1"),
	}
	user, err := newClient.User.UpdateUser(context.Background(), 1, updatedUser)
	require.NoError(t, err)
	require.NotNil(t, user)

	user, err = newClient.User.UpdateUser(context.Background(), 1, updatedUser)
	require.Error(t, err)
	require.Nil(t, user)

	user, err = newClient.User.UpdateUser(context.Background(), 1, updatedUser)
	require.Error(t, err)
	require.Nil(t, user)
}

func TestUserService_DeleteUser(t *testing.T) {
	data := mustReadFileString(t, "testdata/singleuser.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	user, err := newClient.User.DeleteUser(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, user)

	user, err = newClient.User.DeleteUser(context.Background(), 1)
	require.Error(t, err)
	require.Nil(t, user)

	user, err = newClient.User.DeleteUser(context.Background(), 1)
	require.Error(t, err)
	require.Nil(t, user)
}

func TestUserService_CreateUserGroup(t *testing.T) {
	data := mustReadFileString(t, "testdata/singleusergroup.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	newUserGroup := UserGroupRequest{
		Name:        "newgroup",
		Scope:       "group",
		Description: "New Group",
		Member:      []string{"user1", "user2"},
		Priv:        []string{"priv1", "priv2"},
	}
	userGroup, err := newClient.User.CreateUserGroup(context.Background(), newUserGroup)
	require.NoError(t, err)
	require.NotNil(t, userGroup)

	userGroup, err = newClient.User.CreateUserGroup(context.Background(), newUserGroup)
	require.Error(t, err)
	require.Nil(t, userGroup)

	userGroup, err = newClient.User.CreateUserGroup(context.Background(), newUserGroup)
	require.Error(t, err)
	require.Nil(t, userGroup)
}

func TestUserService_UpdateUserGroup(t *testing.T) {
	data := mustReadFileString(t, "testdata/singleusergroup.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	updatedUserGroup := UserGroupRequest{
		Name:        "updatedgroup",
		Scope:       "group",
		Description: "Updated Group",
		Member:      []string{"user1", "user2"},
		Priv:        []string{"priv1", "priv2"},
	}
	userGroup, err := newClient.User.UpdateUserGroup(context.Background(), 1, updatedUserGroup)
	require.NoError(t, err)
	require.NotNil(t, userGroup)

	userGroup, err = newClient.User.UpdateUserGroup(context.Background(), 1, updatedUserGroup)
	require.Error(t, err)
	require.Nil(t, userGroup)

	userGroup, err = newClient.User.UpdateUserGroup(context.Background(), 1, updatedUserGroup)
	require.Error(t, err)
	require.Nil(t, userGroup)
}

func TestUserService_DeleteUserGroup(t *testing.T) {
	data := mustReadFileString(t, "testdata/singleusergroup.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	userGroup, err := newClient.User.DeleteUserGroup(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, userGroup)

	userGroup, err = newClient.User.DeleteUserGroup(context.Background(), 1)
	require.Error(t, err)
	require.Nil(t, userGroup)

	userGroup, err = newClient.User.DeleteUserGroup(context.Background(), 1)
	require.Error(t, err)
	require.Nil(t, userGroup)
}

func TestUserService_PutUserGroups(t *testing.T) {
	data := mustReadFileString(t, "testdata/multipleusergroup.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	newUserGroups := []*UserGroupRequest{
		{
			Name:        "newgroup1",
			Scope:       "group",
			Description: "New Group 1",
			Member:      []string{"user1", "user2"},
			Priv:        []string{"priv1", "priv2"},
		},
		{
			Name:        "newgroup2",
			Scope:       "group",
			Description: "New Group 2",
			Member:      []string{"user3", "user4"},
			Priv:        []string{"priv3", "priv4"},
		},
	}
	userGroups, err := newClient.User.PutUserGroups(context.Background(), newUserGroups)
	require.NoError(t, err)
	require.Len(t, userGroups, 2)

	userGroups, err = newClient.User.PutUserGroups(context.Background(), newUserGroups)
	require.Error(t, err)
	require.Nil(t, userGroups)

	userGroups, err = newClient.User.PutUserGroups(context.Background(), newUserGroups)
	require.Error(t, err)
	require.Nil(t, userGroups)
}
