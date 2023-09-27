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
