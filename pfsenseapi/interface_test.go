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

func setupTestServer(t *testing.T, response string) *httptest.Server {
	data := makeResultList(t, response)

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(data.popStatus())
		_, err := io.WriteString(w, data.popResult())
		require.NoError(t, err)
	}

	return httptest.NewServer(http.HandlerFunc(handler))
}

func TestInterfaceService_ListInterfaces(t *testing.T) {
	data := mustReadFileString(t, "testdata/multipleinterface.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	response, err := newClient.Interface.ListInterfaces(context.Background())
	require.NoError(t, err)
	require.Len(t, response, 2)

	response, err = newClient.Interface.ListInterfaces(context.Background())
	require.Error(t, err)
	require.Nil(t, response)

	response, err = newClient.Interface.ListInterfaces(context.Background())
	require.Error(t, err)
	require.Nil(t, response)
}

func TestInterfaceService_ListInterfaceBridges(t *testing.T) {
	data := mustReadFileString(t, "testdata/multipleinterfacebridge.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	response, err := newClient.Interface.ListInterfaceBridges(context.Background())
	require.NoError(t, err)
	require.Len(t, response, 2)

	response, err = newClient.Interface.ListInterfaceBridges(context.Background())
	require.Error(t, err)
	require.Nil(t, response)

	response, err = newClient.Interface.ListInterfaceBridges(context.Background())
	require.Error(t, err)
	require.Nil(t, response)
}

func TestInterfaceService_ListInterfaceGroups(t *testing.T) {
	data := mustReadFileString(t, "testdata/multipleinterfacegroup.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	response, err := newClient.Interface.ListInterfaceGroups(context.Background())
	require.NoError(t, err)
	require.Len(t, response, 2)

	response, err = newClient.Interface.ListInterfaceGroups(context.Background())
	require.Error(t, err)
	require.Nil(t, response)

	response, err = newClient.Interface.ListInterfaceGroups(context.Background())
	require.Error(t, err)
	require.Nil(t, response)
}

func TestInterfaceService_ListVLANs(t *testing.T) {
	data := mustReadFileString(t, "testdata/multiplevlan.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	response, err := newClient.Interface.ListVLANs(context.Background())
	require.NoError(t, err)
	require.Len(t, response, 2)

	response, err = newClient.Interface.ListVLANs(context.Background())
	require.Error(t, err)
	require.Nil(t, response)

	response, err = newClient.Interface.ListVLANs(context.Background())
	require.Error(t, err)
	require.Nil(t, response)
}

func TestInterfaceService_GetInterface(t *testing.T) {
	data := mustReadFileString(t, "testdata/singleinterface.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	response, err := newClient.Interface.GetInterface(context.Background(), "test_interface")
	require.NoError(t, err)
	require.NotNil(t, response)

	response, err = newClient.Interface.GetInterface(context.Background(), "test_interface")
	require.Error(t, err)
	require.Nil(t, response)

	response, err = newClient.Interface.GetInterface(context.Background(), "test_interface")
	require.Error(t, err)
	require.Nil(t, response)
}

func TestInterfaceService_GetVLAN(t *testing.T) {
	data := mustReadFileString(t, "testdata/singlevlan.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	response, err := newClient.Interface.GetVLAN(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, response)

	response, err = newClient.Interface.GetVLAN(context.Background(), 1)
	require.Error(t, err)
	require.Nil(t, response)

	response, err = newClient.Interface.GetVLAN(context.Background(), 1)
	require.Error(t, err)
	require.Nil(t, response)
}

func TestInterfaceService_GetInterfaceGroup(t *testing.T) {
	data := mustReadFileString(t, "testdata/singleinterfacegroup.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	response, err := newClient.Interface.GetInterfaceGroup(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, response)

	response, err = newClient.Interface.GetInterfaceGroup(context.Background(), 1)
	require.Error(t, err)
	require.Nil(t, response)

	response, err = newClient.Interface.GetInterfaceGroup(context.Background(), 1)
	require.Error(t, err)
	require.Nil(t, response)
}

func TestInterfaceService_GetInterfaceBridge(t *testing.T) {
	data := mustReadFileString(t, "testdata/singleinterfacebridge.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	response, err := newClient.Interface.GetInterfaceBridge(context.Background(), "test_bridge")
	require.NoError(t, err)
	require.NotNil(t, response)

	response, err = newClient.Interface.GetInterfaceBridge(context.Background(), "test_bridge")
	require.Error(t, err)
	require.Nil(t, response)

	response, err = newClient.Interface.GetInterfaceBridge(context.Background(), "test_bridge")
	require.Error(t, err)
	require.Nil(t, response)
}

func TestInterfaceService_CreateInterface(t *testing.T) {
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

	response, err = newClient.Interface.CreateInterface(context.Background(), newInterface)
	require.Error(t, err)
	require.Nil(t, response)

	response, err = newClient.Interface.CreateInterface(context.Background(), newInterface)
	require.Error(t, err)
	require.Nil(t, response)
}

func TestInterfaceService_CreateVLAN(t *testing.T) {
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

	response, err = newClient.Interface.CreateVLAN(context.Background(), newVLAN)
	require.Error(t, err)
	require.Nil(t, response)

	response, err = newClient.Interface.CreateVLAN(context.Background(), newVLAN)
	require.Error(t, err)
	require.Nil(t, response)
}

func TestInterfaceService_CreateInterfaceGroup(t *testing.T) {
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

	response, err = newClient.Interface.CreateInterfaceGroup(context.Background(), newGroup)
	require.Error(t, err)
	require.Nil(t, response)

	response, err = newClient.Interface.CreateInterfaceGroup(context.Background(), newGroup)
	require.Error(t, err)
	require.Nil(t, response)
}

func TestInterfaceService_CreateInterfaceBridge(t *testing.T) {
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

	response, err = newClient.Interface.CreateInterfaceBridge(context.Background(), newBridge)
	require.Error(t, err)
	require.Nil(t, response)

	response, err = newClient.Interface.CreateInterfaceBridge(context.Background(), newBridge)
	require.Error(t, err)
	require.Nil(t, response)
}

func TestInterfaceService_DeleteInterface(t *testing.T) {
	server := setupTestServer(t, "{}")
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	err := newClient.Interface.DeleteInterface(context.Background(), "test_interface")
	require.NoError(t, err)

	err = newClient.Interface.DeleteInterface(context.Background(), "test_interface")
	require.Error(t, err)

	err = newClient.Interface.DeleteInterface(context.Background(), "test_interface")
	require.Error(t, err)
}

func TestInterfaceService_DeleteVLAN(t *testing.T) {
	server := setupTestServer(t, "{}")
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	err := newClient.Interface.DeleteVLAN(context.Background(), 1)
	require.NoError(t, err)

	err = newClient.Interface.DeleteVLAN(context.Background(), 1)
	require.Error(t, err)

	err = newClient.Interface.DeleteVLAN(context.Background(), 1)
	require.Error(t, err)
}

func TestInterfaceService_DeleteInterfaceGroup(t *testing.T) {
	server := setupTestServer(t, "{}")
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	err := newClient.Interface.DeleteInterfaceGroup(context.Background(), 1)
	require.NoError(t, err)

	err = newClient.Interface.DeleteInterfaceGroup(context.Background(), 1)
	require.Error(t, err)

	err = newClient.Interface.DeleteInterfaceGroup(context.Background(), 1)
	require.Error(t, err)
}

func TestInterfaceService_DeleteInterfaceBridge(t *testing.T) {
	server := setupTestServer(t, "{}")
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	err := newClient.Interface.DeleteInterfaceBridge(context.Background(), "test_bridge")
	require.NoError(t, err)

	err = newClient.Interface.DeleteInterfaceBridge(context.Background(), "test_bridge")
	require.Error(t, err)

	err = newClient.Interface.DeleteInterfaceBridge(context.Background(), "test_bridge")
	require.Error(t, err)
}

func TestInterfaceService_UpdateInterface(t *testing.T) {
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

	response, err = newClient.Interface.UpdateInterface(context.Background(), "test_interface", updatedInterface)
	require.Error(t, err)
	require.Nil(t, response)

	response, err = newClient.Interface.UpdateInterface(context.Background(), "test_interface", updatedInterface)
	require.Error(t, err)
	require.Nil(t, response)
}

func TestInterfaceService_UpdateVLAN(t *testing.T) {
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

	response, err = newClient.Interface.UpdateVLAN(context.Background(), 1, updatedVLAN)
	require.Error(t, err)
	require.Nil(t, response)

	response, err = newClient.Interface.UpdateVLAN(context.Background(), 1, updatedVLAN)
	require.Error(t, err)
	require.Nil(t, response)
}

func TestInterfaceService_UpdateInterfaceGroup(t *testing.T) {
	data := mustReadFileString(t, "testdata/singleinterfacegroup.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	updatedGroup := InterfaceGroupRequest{
		Ifname:  "group3",
		Members: []string{"em3", "em4"},
		Descr:   "Updated Test Group 3",
	}
	response, err := newClient.Interface.UpdateInterfaceGroup(context.Background(), 1, updatedGroup)
	require.NoError(t, err)
	require.NotNil(t, response)

	response, err = newClient.Interface.UpdateInterfaceGroup(context.Background(), 1, updatedGroup)
	require.Error(t, err)
	require.Nil(t, response)

	response, err = newClient.Interface.UpdateInterfaceGroup(context.Background(), 1, updatedGroup)
	require.Error(t, err)
	require.Nil(t, response)
}

func TestInterfaceService_UpdateInterfaceBridge(t *testing.T) {
	data := mustReadFileString(t, "testdata/singleinterfacebridge.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	updatedBridge := InterfaceBridgeRequest{
		Members:  []string{"em3", "em4"},
		Descr:    "Updated Test Bridge 3",
		Bridgeif: "bridge2",
	}
	response, err := newClient.Interface.UpdateInterfaceBridge(context.Background(), "test_bridge", updatedBridge)
	require.NoError(t, err)
	require.NotNil(t, response)

	response, err = newClient.Interface.UpdateInterfaceBridge(context.Background(), "test_bridge", updatedBridge)
	require.Error(t, err)
	require.Nil(t, response)

	response, err = newClient.Interface.UpdateInterfaceBridge(context.Background(), "test_bridge", updatedBridge)
	require.Error(t, err)
	require.Nil(t, response)
}

func TestInterfaceService_Apply(t *testing.T) {
	server := setupTestServer(t, "{}")
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	err := newClient.Interface.Apply(context.Background())
	require.NoError(t, err)

	err = newClient.Interface.Apply(context.Background())
	require.Error(t, err)

	err = newClient.Interface.Apply(context.Background())
	require.Error(t, err)
}

func TestInterfaceService_PutInterfaceGroups(t *testing.T) {
	data := mustReadFileString(t, "testdata/multipleinterfacegroup.json")

	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	newGroups := []*InterfaceGroupRequest{
		{
			Ifname:  "group1",
			Members: []string{"em1", "em2"},
			Descr:   "Test Group 1",
		},
		{
			Ifname:  "group2",
			Members: []string{"em3", "em4"},
			Descr:   "Test Group 2",
		},
	}
	response, err := newClient.Interface.PutInterfaceGroups(context.Background(), newGroups)
	require.NoError(t, err)
	require.Len(t, response, 2)

	response, err = newClient.Interface.PutInterfaceGroups(context.Background(), newGroups)
	require.Error(t, err)
	require.Nil(t, response)

	response, err = newClient.Interface.PutInterfaceGroups(context.Background(), newGroups)
	require.Error(t, err)
	require.Nil(t, response)
}
