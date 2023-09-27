package pfsenseapi

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInterfaceService_ListInterfaceGroups(t *testing.T) {
	data := mustReadFileString(t, "testdata/listinterfacegroups.json")

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, data)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	response, err := newClient.Interface.ListInterfaceGroups(context.Background())
	require.NoError(t, err)
	require.Len(t, response, 2)
}

func TestInterfaceService_ListInterfaces(t *testing.T) {
	data := mustReadFileString(t, "testdata/listinterfaces.json")

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, data)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	response, err := newClient.Interface.ListInterfaces(context.Background())
	require.NoError(t, err)
	require.Len(t, response, 2)
}

func TestInterfaceService_ListVLANs(t *testing.T) {
	data := mustReadFileString(t, "testdata/listvlans.json")

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, data)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	response, err := newClient.Interface.ListVLANs(context.Background())
	require.NoError(t, err)
	require.Len(t, response, 2)
}
