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
	data := makeResultList(t, mustReadFileString(t, "testdata/listleases.json"))

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(data.popStatus())
		_, _ = io.WriteString(w, data.popResult())
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	response, err := newClient.DHCP.ListLeases(context.Background())
	require.NoError(t, err)
	require.Len(t, response, 1)

	response, err = newClient.DHCP.ListLeases(context.Background())
	require.Error(t, err)
	require.Nil(t, response)

	response, err = newClient.DHCP.ListLeases(context.Background())
	require.Error(t, err)
	require.Nil(t, response)
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
	require.Len(t, response, 4)
}

func TestDHCPService_DeleteStaticMappings(t *testing.T) {
	listResponse := mustReadFileString(t, "testdata/liststaticmappings.json")
	deleteResponse := mustReadFileString(t, "testdata/deletestaticmapping.json")

	testInterface := "IOT"
	mappingId := "3"
	mappingMac := "00:1d:93:aa:4c"

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

		if r.Method == http.MethodDelete {
			id := query.Get("id")

			if id != mappingId {
				w.WriteHeader(http.StatusBadRequest)
				_, _ = fmt.Fprintf(w, "invalid request")
			} else {
				w.WriteHeader(http.StatusOK)
				_, _ = io.WriteString(w, deleteResponse)
			}
		} else {
			w.WriteHeader(http.StatusOK)
			_, _ = io.WriteString(w, listResponse)
		}
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	err := newClient.DHCP.DeleteStaticMapping(context.Background(), testInterface, mappingMac)
	require.NoError(t, err)
}

func TestDHCPService_UpdateDHCPConfiguration(t *testing.T) {
	data := makeResultList(t, mustReadFileString(t, "testdata/dhcpconfiguration.json"))

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(data.popStatus())
		_, _ = io.WriteString(w, data.popResult())
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	response, err := newClient.DHCP.UpdateServerConfiguration(context.Background(), DHCPServerConfigurationRequest{})
	require.NotNil(t, response)
	require.NoError(t, err)

	response, err = newClient.DHCP.UpdateServerConfiguration(context.Background(), DHCPServerConfigurationRequest{})
	require.Nil(t, response)
	require.Error(t, err)

	response, err = newClient.DHCP.UpdateServerConfiguration(context.Background(), DHCPServerConfigurationRequest{})
	require.Nil(t, response)
	require.Error(t, err)
}
