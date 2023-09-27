package pfsenseapi

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDHCPService_ListLeases(t *testing.T) {
	data := mustReadFileString(t, "testdata/listleases.json")

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, data)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	response, err := newClient.DHCP.ListLeases(context.Background())
	require.NoError(t, err)
	require.Len(t, response, 1)
}

func TestDHCPService_ListStaticMappings(t *testing.T) {
	data := mustReadFileString(t, "testdata/liststaticmappings.json")

	testInterface := "IOT"
	handler := func(w http.ResponseWriter, r *http.Request) {
		query, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = fmt.Fprintf(w, "invalid request")
			return
		}
		interfaceValue := query.Get("interface")
		if interfaceValue != testInterface {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = fmt.Fprintf(w, "invalid request")
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = io.WriteString(w, data)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	response, err := newClient.DHCP.ListStaticMappings(context.Background(), testInterface)
	require.NoError(t, err)
	require.Len(t, response, 1)
}
