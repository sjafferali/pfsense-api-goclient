package pfsenseapi

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStatusService_ListInterfaceStatus(t *testing.T) {
	data := mustReadFileString(t, "testdata/listinterfacestatus.json")

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, data)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	response, err := newClient.Status.ListInterfaceStatus(context.Background())
	require.NoError(t, err)
	require.Len(t, response, 2)
}

func TestStatusService_ListGatewayStatus(t *testing.T) {
	data := mustReadFileString(t, "testdata/listgatewaystatus.json")

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, data)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	response, err := newClient.Status.ListGatewayStatus(context.Background())
	require.NoError(t, err)
	require.Len(t, response, 2)
}

func TestStatusService_SystemLog(t *testing.T) {
	data := mustReadFileString(t, "testdata/systemlog.json")

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, data)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	response, err := newClient.Status.SystemLog(context.Background())
	require.NoError(t, err)
	require.Len(t, response, 2)
}

func TestStatusService_DHCPLog(t *testing.T) {
	data := mustReadFileString(t, "testdata/dhcplog.json")

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, data)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	response, err := newClient.Status.DHCPLog(context.Background())
	require.NoError(t, err)
	require.Len(t, response, 3)
}

func TestStatusService_FirewallLog(t *testing.T) {
	data := mustReadFileString(t, "testdata/firewalllog.json")

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, data)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	response, err := newClient.Status.FirewallLog(context.Background())
	require.NoError(t, err)
	require.Len(t, response, 2)
}
