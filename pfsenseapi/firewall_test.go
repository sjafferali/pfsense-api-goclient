package pfsenseapi

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFirewall_ListRules(t *testing.T) {
	data := makeResultList(t, mustReadFileString(t, "testdata/listfirewallrules.json"))

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(data.popStatus())
		_, _ = io.WriteString(w, data.popResult())
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	response, err := newClient.Firewall.ListRules(context.Background())
	require.NoError(t, err)
	require.Len(t, response, 1)

	response, err = newClient.Firewall.ListRules(context.Background())
	require.Nil(t, response)
	require.Error(t, err)

	response, err = newClient.Firewall.ListRules(context.Background())
	require.Nil(t, response)
	require.Error(t, err)
}

func TestFirewall_CreateRule(t *testing.T) {
	data := mustReadFileString(t, "testdata/createfirewallrule.json")

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = io.WriteString(w, data)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)

	response, err := newClient.Firewall.CreateRule(context.Background(), FirewallRuleRequest{}, true)
	require.NoError(t, err)
	require.NotNil(t, response)
}
