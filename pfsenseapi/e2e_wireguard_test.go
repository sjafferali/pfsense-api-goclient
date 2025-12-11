// +build e2e

package pfsenseapi

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"testing"

	"github.com/markphelps/optional"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/curve25519"
)

// generateWireGuardKeyPair generates a WireGuard private/public key pair
func generateWireGuardKeyPair() (privateKey, publicKey string, err error) {
	var private [32]byte
	if _, err := rand.Read(private[:]); err != nil {
		return "", "", err
	}

	// Clamp the private key
	private[0] &= 248
	private[31] &= 127
	private[31] |= 64

	var public [32]byte
	curve25519.ScalarBaseMult(&public, &private)

	privateKey = base64.StdEncoding.EncodeToString(private[:])
	publicKey = base64.StdEncoding.EncodeToString(public[:])

	return privateKey, publicKey, nil
}

// TestE2E_WireGuardTunnel_CRUD tests the full CRUD lifecycle of a WireGuard tunnel
func TestE2E_WireGuardTunnel_CRUD(t *testing.T) {
	config, err := LoadE2EConfig()
	require.NoError(t, err, "Failed to load e2e config")

	client := NewClientWithLocalAuth(config.URL, config.Username, config.Password)
	ctx := context.Background()

	// Generate keys for the tunnel
	privateKey, publicKey, err := generateWireGuardKeyPair()
	require.NoError(t, err, "Failed to generate key pair")

	// Create a test WireGuard tunnel
	t.Log("Creating WireGuard tunnel...")
	enabled := optional.NewBool(false) // Create disabled to avoid interfering with actual VPN
	listenPort := optional.NewString("51820")
	mtu := optional.NewInt(1420)
	descr := optional.NewString("E2E Test Tunnel - Safe to Delete")
	pubKey := optional.NewString(publicKey)

	newTunnel := WireGuardTunnelRequest{
		Enabled:    &enabled,
		Name:       "tun_e2e_test",
		ListenPort: &listenPort,
		PrivateKey: privateKey,
		PublicKey:  &pubKey,
		MTU:        &mtu,
		Descr:      &descr,
		Address:    []string{"10.99.0.1/24"},
	}

	createdTunnel, err := client.WireGuard.CreateWireGuardTunnel(ctx, newTunnel)
	require.NoError(t, err, "Failed to create WireGuard tunnel")
	require.NotNil(t, createdTunnel)
	require.Equal(t, "tun_e2e_test", createdTunnel.Name)
	t.Logf("Created tunnel: %s", createdTunnel.Name)

	// Ensure cleanup
	defer func() {
		t.Log("Cleaning up: deleting WireGuard tunnel...")
		_, err := client.WireGuard.DeleteWireGuardTunnel(ctx, createdTunnel.Id)
		if err != nil {
			t.Logf("Warning: Failed to delete tunnel ID %d (%s): %v", createdTunnel.Id, createdTunnel.Name, err)
		} else {
			t.Logf("Successfully deleted tunnel ID %d (%s)", createdTunnel.Id, createdTunnel.Name)
		}
	}()

	// Read the tunnel back
	t.Log("Reading WireGuard tunnel...")
	readTunnel, err := client.WireGuard.GetWireGuardTunnel(ctx, createdTunnel.Id)
	require.NoError(t, err, "Failed to get WireGuard tunnel")
	require.NotNil(t, readTunnel)
	require.Equal(t, createdTunnel.Id, readTunnel.Id)
	require.Equal(t, createdTunnel.Name, readTunnel.Name)

	descrValue, err := readTunnel.Descr.Get()
	require.NoError(t, err)
	require.Equal(t, "E2E Test Tunnel - Safe to Delete", descrValue)

	// Update the tunnel
	t.Log("Updating WireGuard tunnel...")
	updatedDescr := optional.NewString("E2E Test Tunnel - Updated - Safe to Delete")
	updatedListenPort := optional.NewString("51821")

	updateTunnel := WireGuardTunnelRequest{
		Enabled:    &enabled,
		Name:       "tun_e2e_test",
		ListenPort: &updatedListenPort,
		PrivateKey: privateKey,
		PublicKey:  &pubKey,
		MTU:        &mtu,
		Descr:      &updatedDescr,
		Address:    []string{"10.99.0.1/24"},
	}

	updatedTunnel, err := client.WireGuard.UpdateWireGuardTunnel(ctx, createdTunnel.Id, updateTunnel)
	require.NoError(t, err, "Failed to update WireGuard tunnel")
	require.NotNil(t, updatedTunnel)
	require.Equal(t, createdTunnel.Id, updatedTunnel.Id)

	updatedDescrValue, err := updatedTunnel.Descr.Get()
	require.NoError(t, err)
	require.Equal(t, "E2E Test Tunnel - Updated - Safe to Delete", updatedDescrValue)

	updatedPortValue, err := updatedTunnel.ListenPort.Get()
	require.NoError(t, err)
	require.Equal(t, "51821", updatedPortValue)

	t.Log("WireGuard tunnel CRUD test completed successfully")
}

// TestE2E_WireGuardTunnels_List tests listing WireGuard tunnels
func TestE2E_WireGuardTunnels_List(t *testing.T) {
	config, err := LoadE2EConfig()
	require.NoError(t, err, "Failed to load e2e config")

	client := NewClientWithLocalAuth(config.URL, config.Username, config.Password)
	ctx := context.Background()

	t.Log("Listing WireGuard tunnels...")
	tunnels, err := client.WireGuard.ListWireGuardTunnels(ctx)
	require.NoError(t, err, "Failed to list WireGuard tunnels")
	require.NotNil(t, tunnels)

	t.Logf("Found %d WireGuard tunnels", len(tunnels))

	// Verify structure of returned tunnels
	for i, tunnel := range tunnels {
		require.NotEmpty(t, tunnel.Name, "Tunnel %d should have a name", i)
		t.Logf("Tunnel %d: Name=%s", i, tunnel.Name)
	}
}

// TestE2E_WireGuardPeer_CRUD tests the full CRUD lifecycle of a WireGuard peer
// Note: This test requires a WireGuard tunnel to exist
func TestE2E_WireGuardPeer_CRUD(t *testing.T) {
	config, err := LoadE2EConfig()
	require.NoError(t, err, "Failed to load e2e config")

	client := NewClientWithLocalAuth(config.URL, config.Username, config.Password)
	ctx := context.Background()

	// First, create a tunnel for the peer
	privateKey, publicKey, err := generateWireGuardKeyPair()
	require.NoError(t, err, "Failed to generate tunnel key pair")

	t.Log("Creating WireGuard tunnel for peer test...")
	enabled := optional.NewBool(false)
	listenPort := optional.NewString("51822")
	mtu := optional.NewInt(1420)
	descr := optional.NewString("E2E Test Tunnel for Peer - Safe to Delete")
	pubKey := optional.NewString(publicKey)

	newTunnel := WireGuardTunnelRequest{
		Enabled:    &enabled,
		Name:       "tun_e2e_peer_test",
		ListenPort: &listenPort,
		PrivateKey: privateKey,
		PublicKey:  &pubKey,
		MTU:        &mtu,
		Descr:      &descr,
		Address:    []string{"10.98.0.1/24"},
	}

	createdTunnel, err := client.WireGuard.CreateWireGuardTunnel(ctx, newTunnel)
	require.NoError(t, err, "Failed to create WireGuard tunnel")
	require.NotNil(t, createdTunnel)
	t.Logf("Created tunnel: %s", createdTunnel.Name)

	// Ensure tunnel cleanup
	defer func() {
		t.Log("Cleaning up: deleting WireGuard tunnel...")
		_, err := client.WireGuard.DeleteWireGuardTunnel(ctx, createdTunnel.Id)
		if err != nil {
			t.Logf("Warning: Failed to delete tunnel ID %d (%s): %v", createdTunnel.Id, createdTunnel.Name, err)
		} else {
			t.Logf("Successfully deleted tunnel ID %d (%s)", createdTunnel.Id, createdTunnel.Name)
		}
	}()

	// Generate keys for the peer
	_, peerPublicKey, err := generateWireGuardKeyPair()
	require.NoError(t, err, "Failed to generate peer key pair")

	// Create a test WireGuard peer
	t.Log("Creating WireGuard peer...")
	peerEnabled := optional.NewBool(true)
	endpoint := optional.NewString("192.168.1.100")
	port := optional.NewString("51820")
	peerDescr := optional.NewString("E2E Test Peer - Safe to Delete")
	keepalive := optional.NewInt(25)

	newPeer := WireGuardPeerRequest{
		Enabled:             &peerEnabled,
		Tun:                 createdTunnel.Name,
		Endpoint:            &endpoint,
		Port:                &port,
		Descr:               &peerDescr,
		PersistentKeepalive: &keepalive,
		PublicKey:           peerPublicKey,
		AllowedIPs: []WireGuardAllowedIP{
			{
				Address: "10.98.0.2",
				Mask:    32,
				Descr:   &peerDescr,
			},
		},
	}

	createdPeer, err := client.WireGuard.CreateWireGuardPeer(ctx, newPeer)
	require.NoError(t, err, "Failed to create WireGuard peer")
	require.NotNil(t, createdPeer)
	require.Greater(t, createdPeer.Id, 0)
	t.Logf("Created peer with ID: %d", createdPeer.Id)

	// Ensure peer cleanup
	defer func() {
		t.Log("Cleaning up: deleting WireGuard peer...")
		_, err := client.WireGuard.DeleteWireGuardPeer(ctx, createdPeer.Id)
		if err != nil {
			t.Logf("Warning: Failed to delete peer %d: %v", createdPeer.Id, err)
		} else {
			t.Logf("Successfully deleted peer %d", createdPeer.Id)
		}
	}()

	// Read the peer back
	t.Log("Reading WireGuard peer...")
	readPeer, err := client.WireGuard.GetWireGuardPeer(ctx, createdPeer.Id)
	require.NoError(t, err, "Failed to get WireGuard peer")
	require.NotNil(t, readPeer)
	require.Equal(t, createdPeer.Id, readPeer.Id)

	readDescrValue, err := readPeer.Descr.Get()
	require.NoError(t, err)
	require.Equal(t, "E2E Test Peer - Safe to Delete", readDescrValue)

	// Update the peer
	t.Log("Updating WireGuard peer...")
	updatedPeerDescr := optional.NewString("E2E Test Peer - Updated - Safe to Delete")
	updatedEndpoint := optional.NewString("192.168.1.101")

	updatePeer := WireGuardPeerRequest{
		Enabled:             &peerEnabled,
		Tun:                 createdTunnel.Name,
		Endpoint:            &updatedEndpoint,
		Port:                &port,
		Descr:               &updatedPeerDescr,
		PersistentKeepalive: &keepalive,
		PublicKey:           peerPublicKey,
		AllowedIPs: []WireGuardAllowedIP{
			{
				Address: "10.98.0.2",
				Mask:    32,
				Descr:   &updatedPeerDescr,
			},
		},
	}

	updatedPeer, err := client.WireGuard.UpdateWireGuardPeer(ctx, createdPeer.Id, updatePeer)
	require.NoError(t, err, "Failed to update WireGuard peer")
	require.NotNil(t, updatedPeer)

	updatedPeerDescrValue, err := updatedPeer.Descr.Get()
	require.NoError(t, err)
	require.Equal(t, "E2E Test Peer - Updated - Safe to Delete", updatedPeerDescrValue)

	updatedEndpointValue, err := updatedPeer.Endpoint.Get()
	require.NoError(t, err)
	require.Equal(t, "192.168.1.101", updatedEndpointValue)

	t.Log("WireGuard peer CRUD test completed successfully")
}

// TestE2E_WireGuardPeers_List tests listing WireGuard peers
func TestE2E_WireGuardPeers_List(t *testing.T) {
	config, err := LoadE2EConfig()
	require.NoError(t, err, "Failed to load e2e config")

	client := NewClientWithLocalAuth(config.URL, config.Username, config.Password)
	ctx := context.Background()

	t.Log("Listing WireGuard peers...")
	peers, err := client.WireGuard.ListWireGuardPeers(ctx)
	require.NoError(t, err, "Failed to list WireGuard peers")
	require.NotNil(t, peers)

	t.Logf("Found %d WireGuard peers", len(peers))

	// Verify structure of returned peers
	for i, peer := range peers {
		require.GreaterOrEqual(t, peer.Id, 0, "Peer %d should have a valid ID", i)
		t.Logf("Peer %d: ID=%d", i, peer.Id)
	}
}
