package pfsenseapi

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUserService_ListGroups(t *testing.T) {
	data := mustReadFileString(t, "testdata/listgroups.json")

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, data)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	response, err := newClient.User.ListGroups(context.Background())
	require.NoError(t, err)
	require.Len(t, response, 2)
}

func TestUserService_ListUsers(t *testing.T) {
	data := mustReadFileString(t, "testdata/listusers.json")

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, data)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	response, err := newClient.User.ListUsers(context.Background())
	require.NoError(t, err)
	require.Len(t, response, 2)
}

func TestUserService_CreateUser(t *testing.T) {
	data := mustReadFileString(t, "testdata/createuser.json")

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, data)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	response, err := newClient.User.CreateUser(context.Background(), UserRequest{})
	require.NotNil(t, response)
	require.NoError(t, err)
}

func TestUserService_UpdateUser(t *testing.T) {
	data := mustReadFileString(t, "testdata/createuser.json")

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, data)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	response, err := newClient.User.UpdateUser(context.Background(), UserRequest{})
	require.NotNil(t, response)
	require.NoError(t, err)
}

func TestUserService_DeleteUser(t *testing.T) {
	data := mustReadFileString(t, "testdata/createuser.json")

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, data)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	err := newClient.User.DeleteUser(context.Background(), "testing123")
	require.NoError(t, err)
}

func TestUserService_CreateGroup(t *testing.T) {
	data := mustReadFileString(t, "testdata/creategroup.json")

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, data)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	response, err := newClient.User.CreateGroup(context.Background(), GroupRequest{})
	require.NotNil(t, response)
	require.NoError(t, err)
}

func TestUserService_UpdateGroup(t *testing.T) {
	data := mustReadFileString(t, "testdata/creategroup.json")

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, data)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	response, err := newClient.User.UpdateGroup(context.Background(), "admin", GroupRequest{})
	require.NotNil(t, response)
	require.NoError(t, err)
}

func TestUserService_DeleteGroup(t *testing.T) {
	data := mustReadFileString(t, "testdata/creategroup.json")

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, data)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	err := newClient.User.DeleteGroup(context.Background(), "admin")
	require.NoError(t, err)
}

func TestUserService_RemoveUserFromGroup(t *testing.T) {
	data := mustReadFileString(t, "testdata/addusertogroups.json")

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, data)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	err := newClient.User.RemoveUserFromGroup(context.Background(), "admin", "admins")
	require.NoError(t, err)
}

func TestUserService_AddUserToGroups(t *testing.T) {
	data := mustReadFileString(t, "testdata/addusertogroups.json")

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, data)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	err := newClient.User.AddUserToGroups(context.Background(), "admin", []string{"admins"})
	require.NoError(t, err)
}

func TestUserService_RemovePrivilegeFromUser(t *testing.T) {
	data := mustReadFileString(t, "testdata/addprivilegestouser.json")

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, data)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	err := newClient.User.RemovePrivilegeFromUser(context.Background(), "admin", "system-xmlrpc-ha-sync")
	require.NoError(t, err)
}

func TestUserService_AddPrivilegesToUser(t *testing.T) {
	data := mustReadFileString(t, "testdata/addprivilegestouser.json")

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, data)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	err := newClient.User.AddPrivilegesToUser(context.Background(), "admin", []string{"system-xmlrpc-ha-sync"})
	require.NoError(t, err)
}
