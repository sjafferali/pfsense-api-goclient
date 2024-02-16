package pfsenseapi

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/markphelps/optional"
	"github.com/stretchr/testify/require"
)

func setupTestServer(t *testing.T, data string) *httptest.Server {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, err := io.WriteString(w, data)
		require.NoError(t, err)
	}

	return httptest.NewServer(http.HandlerFunc(handler))
}

func TestInterfaceService_ListInterfacesReturnsExpectedCount(t *testing.T) {
	data := mustReadFileString(t, "testdata/multipleinterface.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	response, err := newClient.Interface.ListInterfaces(context.Background())
	require.NoError(t, err)
	require.Len(t, response, 2)
}

func TestInterfaceService_ListInterfaceBridgesReturnsExpectedCount(t *testing.T) {
	data := mustReadFileString(t, "testdata/multipleinterfacebridge.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	response, err := newClient.Interface.ListInterfaceBridges(context.Background())
	require.NoError(t, err)
	require.Len(t, response, 2)
}

func TestInterfaceService_ListInterfaceGroupsReturnsExpectedCount(t *testing.T) {
	data := mustReadFileString(t, "testdata/multipleinterfacegroup.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	response, err := newClient.Interface.ListInterfaceGroups(context.Background())
	require.NoError(t, err)
	require.Len(t, response, 2)
}

func TestInterfaceService_ListVLANsReturnsExpectedCount(t *testing.T) {
	data := mustReadFileString(t, "testdata/multiplevlan.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	response, err := newClient.Interface.ListVLANs(context.Background())
	require.NoError(t, err)
	require.Len(t, response, 2)
}

func TestInterfaceService_GetInterfaceReturnsExpectedResponse(t *testing.T) {
	data := mustReadFileString(t, "testdata/singleinterface.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	response, err := newClient.Interface.GetInterface(context.Background(), "test_interface")
	require.NoError(t, err)
	require.NotNil(t, response)
}

func TestInterfaceService_GetVLANReturnsExpectedResponse(t *testing.T) {
	data := mustReadFileString(t, "testdata/singlevlan.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	response, err := newClient.Interface.GetVLAN(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, response)
}

func TestInterfaceService_GetInterfaceGroupReturnsExpectedResponse(t *testing.T) {
	data := mustReadFileString(t, "testdata/singleinterfacegroup.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	response, err := newClient.Interface.GetInterfaceGroup(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, response)
}

func TestInterfaceService_GetInterfaceBridgeReturnsExpectedResponse(t *testing.T) {
	data := mustReadFileString(t, "testdata/singleinterfacebridge.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	response, err := newClient.Interface.GetInterfaceBridge(context.Background(), "test_bridge")
	require.NoError(t, err)
	require.NotNil(t, response)
}

func TestInterfaceService_CreateInterfaceReturnsExpectedResponse(t *testing.T) {
	data := mustReadFileString(t, "testdata/singleinterface.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	enable := optional.NewBool(true)
	newInterface := InterfaceRequest{
		If:     "em3",
		Enable: &enable,
		Descr:  "Test Interface 3",
		Typev4: "staticv4",
		Ipaddr: "192.168.1.3",
		Subnet: 24,
	}
	response, err := newClient.Interface.CreateInterface(context.Background(), newInterface)
	require.NoError(t, err)
	require.NotNil(t, response)
}

func TestInterfaceService_CreateVLANReturnsExpectedResponse(t *testing.T) {
	data := mustReadFileString(t, "testdata/singlevlan.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	desc := optional.NewString("Test VLAN 3")
	newVLAN := VLANRequest{
		If:    "em3",
		Tag:   300,
		Descr: &desc,
	}
	response, err := newClient.Interface.CreateVLAN(context.Background(), newVLAN)
	require.NoError(t, err)
	require.NotNil(t, response)
}

func TestInterfaceService_CreateInterfaceGroupReturnsExpectedResponse(t *testing.T) {
	data := mustReadFileString(t, "testdata/singleinterfacegroup.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	newGroup := InterfaceGroupRequest{
		Ifname:  "group3",
		Members: []string{"em3", "em4"},
		Descr:   "Test Group 3",
	}
	response, err := newClient.Interface.CreateInterfaceGroup(context.Background(), newGroup)
	require.NoError(t, err)
	require.NotNil(t, response)
}

func TestInterfaceService_CreateInterfaceBridgeReturnsExpectedResponse(t *testing.T) {
	data := mustReadFileString(t, "testdata/singleinterfacebridge.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	newBridge := InterfaceBridgeRequest{
		Members:  []string{"em3", "em4"},
		Descr:    "Test Bridge 3",
		Bridgeif: "bridge2",
	}
	response, err := newClient.Interface.CreateInterfaceBridge(context.Background(), newBridge)
	require.NoError(t, err)
	require.NotNil(t, response)
}

func TestInterfaceService_DeleteInterfaceReturnsNoError(t *testing.T) {
	server := setupTestServer(t, "")
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	err := newClient.Interface.DeleteInterface(context.Background(), "test_interface")
	require.NoError(t, err)
}

func TestInterfaceService_DeleteVLANReturnsNoError(t *testing.T) {
	server := setupTestServer(t, "")
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	err := newClient.Interface.DeleteVLAN(context.Background(), 1)
	require.NoError(t, err)
}

func TestInterfaceService_DeleteInterfaceGroupReturnsNoError(t *testing.T) {
	server := setupTestServer(t, "")
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	err := newClient.Interface.DeleteInterfaceGroup(context.Background(), 1)
	require.NoError(t, err)
}

func TestInterfaceService_DeleteInterfaceBridgeReturnsNoError(t *testing.T) {
	server := setupTestServer(t, "")
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	err := newClient.Interface.DeleteInterfaceBridge(context.Background(), "test_bridge")
	require.NoError(t, err)
}

func TestInterfaceService_UpdateInterfaceReturnsExpectedResponse(t *testing.T) {
	data := mustReadFileString(t, "testdata/singleinterface.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	enable := optional.NewBool(true)
	updatedInterface := InterfaceRequest{
		If:     "em3",
		Enable: &enable,
		Descr:  "Updated Test Interface 3",
		Typev4: "staticv4",
		Ipaddr: "192.168.1.3",
		Subnet: 24,
	}
	response, err := newClient.Interface.UpdateInterface(context.Background(), "test_interface", updatedInterface)
	require.NoError(t, err)
	require.NotNil(t, response)
}

func TestInterfaceService_UpdateVLANReturnsExpectedResponse(t *testing.T) {
	data := mustReadFileString(t, "testdata/singlevlan.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	desc := optional.NewString("Updated Test VLAN 3")
	updatedVLAN := VLANRequest{
		If:    "em3",
		Tag:   300,
		Descr: &desc,
	}
	response, err := newClient.Interface.UpdateVLAN(context.Background(), 1, updatedVLAN)
	require.NoError(t, err)
	require.NotNil(t, response)
}

func TestInterfaceService_UpdateInterfaceGroupReturnsExpectedResponse(t *testing.T) {
	data := mustReadFileString(t, "testdata/singleinterfacegroup.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	updatedGroup := InterfaceGroupRequest{
		Ifname:  "group3",
		Members: []string{"em3", "em4", "em5"},
		Descr:   "Updated Test Group 3",
	}
	response, err := newClient.Interface.UpdateInterfaceGroup(context.Background(), 1, updatedGroup)
	require.NoError(t, err)
	require.NotNil(t, response)
}

func TestInterfaceService_UpdateInterfaceBridgeReturnsExpectedResponse(t *testing.T) {
	data := mustReadFileString(t, "testdata/singleinterfacebridge.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	updatedBridge := InterfaceBridgeRequest{
		Members:  []string{"em3", "em4", "em5"},
		Descr:    "Updated Test Bridge 3",
		Bridgeif: "bridge2",
	}
	response, err := newClient.Interface.UpdateInterfaceBridge(context.Background(), "test_bridge", updatedBridge)
	require.NoError(t, err)
	require.NotNil(t, response)
}

func TestInterfaceService_ApplyReturnsNoError(t *testing.T) {
	server := setupTestServer(t, "")
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	err := newClient.Interface.Apply(context.Background())
	require.NoError(t, err)
}
