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

const (
	dhcpLeasesTestResponse = `{"status":"ok","code":200,"return":0,"message":"Success",
"data":[{"ip":"192.168.60.7","type":"static","mac":"b4:5f:56:22:d4:33","if":"opt10",
"starts":"","ends":"","hostname":"host1","descr":"host1",
"online":true,"staticmap_array_index":1,"state":"static"}]}`
	dhcpStaticMappingsTestResponse = `{"status":"ok","code":200,"return":0,"message":"Success",
"data":[{"id":0,"mac":"b4:5f:56:22:d4:33","cid":"","ipaddr":"192.168.1.2","hostname":"host1"
,"descr":"host1","filename":"","rootpath":"","defaultleasetime":"",
"maxleasetime":"","gateway":"","domain":"","domainsearchlist":"","ddnsdomain":"",
"ddnsdomainprimary":"","ddnsdomainsecondary":"","ddnsdomainkeyname":"","ddnsdomainkeyalgorithm":
"hmac-md5","ddnsdomainkey":"","tftp":"","ldap":"","nextserver":"","filename32":"","filename64":""
,"filename32arm":"","filename64arm":"","numberoptions":""}]}`
)

func TestDHCPService_ListLeases(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, dhcpLeasesTestResponse)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	response, err := newClient.DHCP.ListLeases(context.Background())
	require.NoError(t, err)
	require.Len(t, response, 1)
}

func TestDHCPService_ListStaticMappings(t *testing.T) {
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
		_, _ = io.WriteString(w, dhcpStaticMappingsTestResponse)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	response, err := newClient.DHCP.ListStaticMappings(context.Background(), testInterface)
	require.NoError(t, err)
	require.Len(t, response, 1)
}
