package pfsenseapi

import (
	"context"
	"testing"

	"github.com/markphelps/optional"
	"github.com/stretchr/testify/require"
)

func TestWireGuardService_ListWireGuardTunnels(t *testing.T) {
	data := mustReadFileString(t, "testdata/multiplewireguardtunnel.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	tunnels, err := newClient.WireGuard.ListWireGuardTunnels(context.Background())
	require.NoError(t, err)
	require.Len(t, tunnels, 2)

	tunnels, err = newClient.WireGuard.ListWireGuardTunnels(context.Background())
	require.Error(t, err)
	require.Nil(t, tunnels)

	tunnels, err = newClient.WireGuard.ListWireGuardTunnels(context.Background())
	require.Error(t, err)
	require.Nil(t, tunnels)
}

func TestWireGuardService_GetWireGuardTunnel(t *testing.T) {
	data := mustReadFileString(t, "testdata/singlewireguardtunnel.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	tunnel, err := newClient.WireGuard.GetWireGuardTunnel(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, tunnel)
	require.Equal(t, 1, tunnel.Id)
	require.Equal(t, "tun_wg0", tunnel.Name)

	tunnel, err = newClient.WireGuard.GetWireGuardTunnel(context.Background(), 1)
	require.Error(t, err)
	require.Nil(t, tunnel)

	tunnel, err = newClient.WireGuard.GetWireGuardTunnel(context.Background(), 1)
	require.Error(t, err)
	require.Nil(t, tunnel)
}

func TestWireGuardService_CreateWireGuardTunnel(t *testing.T) {
	data := mustReadFileString(t, "testdata/singlewireguardtunnel.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	enabled := optional.NewBool(true)
	listenPort := optional.NewString("51820")
	mtu := optional.NewInt(1420)
	descr := optional.NewString("Test WireGuard Tunnel")

	newTunnel := WireGuardTunnelRequest{
		Enabled:    &enabled,
		Name:       "tun_wg0",
		ListenPort: &listenPort,
		PrivateKey: "aBcDeFgHiJkLmNoPqRsTuVwXyZ0123456789ABCDEF=",
		MTU:        &mtu,
		Descr:      &descr,
		Address:    []string{"10.0.0.1/24"},
	}

	tunnel, err := newClient.WireGuard.CreateWireGuardTunnel(context.Background(), newTunnel)
	require.NoError(t, err)
	require.NotNil(t, tunnel)

	tunnel, err = newClient.WireGuard.CreateWireGuardTunnel(context.Background(), newTunnel)
	require.Error(t, err)
	require.Nil(t, tunnel)

	tunnel, err = newClient.WireGuard.CreateWireGuardTunnel(context.Background(), newTunnel)
	require.Error(t, err)
	require.Nil(t, tunnel)
}

func TestWireGuardService_UpdateWireGuardTunnel(t *testing.T) {
	data := mustReadFileString(t, "testdata/singlewireguardtunnel.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	enabled := optional.NewBool(true)
	listenPort := optional.NewString("51820")
	mtu := optional.NewInt(1420)
	descr := optional.NewString("Updated WireGuard Tunnel")

	updatedTunnel := WireGuardTunnelRequest{
		Enabled:    &enabled,
		Name:       "tun_wg0",
		ListenPort: &listenPort,
		PrivateKey: "aBcDeFgHiJkLmNoPqRsTuVwXyZ0123456789ABCDEF=",
		MTU:        &mtu,
		Descr:      &descr,
		Address:    []string{"10.0.0.1/24"},
	}

	tunnel, err := newClient.WireGuard.UpdateWireGuardTunnel(context.Background(), 1, updatedTunnel)
	require.NoError(t, err)
	require.NotNil(t, tunnel)
	require.Equal(t, 1, tunnel.Id)

	tunnel, err = newClient.WireGuard.UpdateWireGuardTunnel(context.Background(), 1, updatedTunnel)
	require.Error(t, err)
	require.Nil(t, tunnel)

	tunnel, err = newClient.WireGuard.UpdateWireGuardTunnel(context.Background(), 1, updatedTunnel)
	require.Error(t, err)
	require.Nil(t, tunnel)
}

func TestWireGuardService_DeleteWireGuardTunnel(t *testing.T) {
	data := mustReadFileString(t, "testdata/singlewireguardtunnel.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	tunnel, err := newClient.WireGuard.DeleteWireGuardTunnel(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, tunnel)
	require.Equal(t, 1, tunnel.Id)

	tunnel, err = newClient.WireGuard.DeleteWireGuardTunnel(context.Background(), 1)
	require.Error(t, err)
	require.Nil(t, tunnel)

	tunnel, err = newClient.WireGuard.DeleteWireGuardTunnel(context.Background(), 1)
	require.Error(t, err)
	require.Nil(t, tunnel)
}

func TestWireGuardService_ListWireGuardPeers(t *testing.T) {
	data := mustReadFileString(t, "testdata/multiplewireguardpeer.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	peers, err := newClient.WireGuard.ListWireGuardPeers(context.Background())
	require.NoError(t, err)
	require.Len(t, peers, 2)

	peers, err = newClient.WireGuard.ListWireGuardPeers(context.Background())
	require.Error(t, err)
	require.Nil(t, peers)

	peers, err = newClient.WireGuard.ListWireGuardPeers(context.Background())
	require.Error(t, err)
	require.Nil(t, peers)
}

func TestWireGuardService_GetWireGuardPeer(t *testing.T) {
	data := mustReadFileString(t, "testdata/singlewireguardpeer.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	peer, err := newClient.WireGuard.GetWireGuardPeer(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, peer)
	require.Equal(t, 1, peer.Id)

	peer, err = newClient.WireGuard.GetWireGuardPeer(context.Background(), 1)
	require.Error(t, err)
	require.Nil(t, peer)

	peer, err = newClient.WireGuard.GetWireGuardPeer(context.Background(), 1)
	require.Error(t, err)
	require.Nil(t, peer)
}

func TestWireGuardService_CreateWireGuardPeer(t *testing.T) {
	data := mustReadFileString(t, "testdata/singlewireguardpeer.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	enabled := optional.NewBool(true)
	endpoint := optional.NewString("192.168.1.100")
	port := optional.NewString("51820")
	descr := optional.NewString("Test Peer")
	keepalive := optional.NewInt(25)
	parentId := optional.NewInt(1)
	allowedId := optional.NewInt(0)
	allowedDescr := optional.NewString("Peer IP")

	newPeer := WireGuardPeerRequest{
		Enabled:             &enabled,
		Tun:                 "tun_wg0",
		Endpoint:            &endpoint,
		Port:                &port,
		Descr:               &descr,
		PersistentKeepalive: &keepalive,
		PublicKey:           "PeerPublicKey123456789ABCDEF=",
		AllowedIPs: []WireGuardAllowedIP{
			{
				ParentId: &parentId,
				Id:       &allowedId,
				Address:  "10.0.0.2",
				Mask:     32,
				Descr:    &allowedDescr,
			},
		},
	}

	peer, err := newClient.WireGuard.CreateWireGuardPeer(context.Background(), newPeer)
	require.NoError(t, err)
	require.NotNil(t, peer)

	peer, err = newClient.WireGuard.CreateWireGuardPeer(context.Background(), newPeer)
	require.Error(t, err)
	require.Nil(t, peer)

	peer, err = newClient.WireGuard.CreateWireGuardPeer(context.Background(), newPeer)
	require.Error(t, err)
	require.Nil(t, peer)
}

func TestWireGuardService_UpdateWireGuardPeer(t *testing.T) {
	data := mustReadFileString(t, "testdata/singlewireguardpeer.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	enabled := optional.NewBool(true)
	endpoint := optional.NewString("192.168.1.100")
	port := optional.NewString("51820")
	descr := optional.NewString("Updated Peer")
	keepalive := optional.NewInt(25)
	parentId := optional.NewInt(1)
	allowedId := optional.NewInt(0)
	allowedDescr := optional.NewString("Peer IP")

	updatedPeer := WireGuardPeerRequest{
		Enabled:             &enabled,
		Tun:                 "tun_wg0",
		Endpoint:            &endpoint,
		Port:                &port,
		Descr:               &descr,
		PersistentKeepalive: &keepalive,
		PublicKey:           "PeerPublicKey123456789ABCDEF=",
		AllowedIPs: []WireGuardAllowedIP{
			{
				ParentId: &parentId,
				Id:       &allowedId,
				Address:  "10.0.0.2",
				Mask:     32,
				Descr:    &allowedDescr,
			},
		},
	}

	peer, err := newClient.WireGuard.UpdateWireGuardPeer(context.Background(), 1, updatedPeer)
	require.NoError(t, err)
	require.NotNil(t, peer)

	peer, err = newClient.WireGuard.UpdateWireGuardPeer(context.Background(), 1, updatedPeer)
	require.Error(t, err)
	require.Nil(t, peer)

	peer, err = newClient.WireGuard.UpdateWireGuardPeer(context.Background(), 1, updatedPeer)
	require.Error(t, err)
	require.Nil(t, peer)
}

func TestWireGuardService_DeleteWireGuardPeer(t *testing.T) {
	data := mustReadFileString(t, "testdata/singlewireguardpeer.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	peer, err := newClient.WireGuard.DeleteWireGuardPeer(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, peer)

	peer, err = newClient.WireGuard.DeleteWireGuardPeer(context.Background(), 1)
	require.Error(t, err)
	require.Nil(t, peer)

	peer, err = newClient.WireGuard.DeleteWireGuardPeer(context.Background(), 1)
	require.Error(t, err)
	require.Nil(t, peer)
}
