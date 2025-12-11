package pfsenseapi

import (
	"context"
	"testing"

	"github.com/markphelps/optional"
	"github.com/stretchr/testify/require"
)

func TestFirewallService_ListFirewallRules(t *testing.T) {
	data := mustReadFileString(t, "testdata/multiplefirewallrule.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	rules, err := newClient.Firewall.ListFirewallRules(context.Background())
	require.NoError(t, err)
	require.Len(t, rules, 2)

	rules, err = newClient.Firewall.ListFirewallRules(context.Background())
	require.Error(t, err)
	require.Nil(t, rules)

	rules, err = newClient.Firewall.ListFirewallRules(context.Background())
	require.Error(t, err)
	require.Nil(t, rules)
}

func TestFirewallService_GetFirewallRule(t *testing.T) {
	data := mustReadFileString(t, "testdata/singlefirewallrule.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	rule, err := newClient.Firewall.GetFirewallRule(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, rule)
	require.Equal(t, 1, rule.Id)

	rule, err = newClient.Firewall.GetFirewallRule(context.Background(), 1)
	require.Error(t, err)
	require.Nil(t, rule)

	rule, err = newClient.Firewall.GetFirewallRule(context.Background(), 1)
	require.Error(t, err)
	require.Nil(t, rule)
}

func TestFirewallService_CreateFirewallRule(t *testing.T) {
	data := mustReadFileString(t, "testdata/singlefirewallrule.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	ruleType := optional.NewString("pass")
	ipProto := optional.NewString("inet")
	proto := optional.NewString("tcp")
	src := optional.NewString("any")
	dst := optional.NewString("any")
	dstPort := optional.NewString("443")
	descr := optional.NewString("Allow HTTPS")
	log := optional.NewBool(true)
	disabled := optional.NewBool(false)

	newRule := FirewallRuleRequest{
		Type:            &ruleType,
		Interface:       []string{"wan"},
		IPProtocol:      &ipProto,
		Protocol:        &proto,
		Source:          &src,
		Destination:     &dst,
		DestinationPort: &dstPort,
		Descr:           &descr,
		Log:             &log,
		Disabled:        &disabled,
	}

	rule, err := newClient.Firewall.CreateFirewallRule(context.Background(), newRule)
	require.NoError(t, err)
	require.NotNil(t, rule)

	rule, err = newClient.Firewall.CreateFirewallRule(context.Background(), newRule)
	require.Error(t, err)
	require.Nil(t, rule)

	rule, err = newClient.Firewall.CreateFirewallRule(context.Background(), newRule)
	require.Error(t, err)
	require.Nil(t, rule)
}

func TestFirewallService_UpdateFirewallRule(t *testing.T) {
	data := mustReadFileString(t, "testdata/singlefirewallrule.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	ruleType := optional.NewString("pass")
	ipProto := optional.NewString("inet")
	proto := optional.NewString("tcp")
	src := optional.NewString("any")
	dst := optional.NewString("any")
	dstPort := optional.NewString("443")
	descr := optional.NewString("Updated HTTPS Rule")
	log := optional.NewBool(true)
	disabled := optional.NewBool(false)

	updatedRule := FirewallRuleRequest{
		Type:            &ruleType,
		Interface:       []string{"wan"},
		IPProtocol:      &ipProto,
		Protocol:        &proto,
		Source:          &src,
		Destination:     &dst,
		DestinationPort: &dstPort,
		Descr:           &descr,
		Log:             &log,
		Disabled:        &disabled,
	}

	rule, err := newClient.Firewall.UpdateFirewallRule(context.Background(), 1, updatedRule)
	require.NoError(t, err)
	require.NotNil(t, rule)

	rule, err = newClient.Firewall.UpdateFirewallRule(context.Background(), 1, updatedRule)
	require.Error(t, err)
	require.Nil(t, rule)

	rule, err = newClient.Firewall.UpdateFirewallRule(context.Background(), 1, updatedRule)
	require.Error(t, err)
	require.Nil(t, rule)
}

func TestFirewallService_DeleteFirewallRule(t *testing.T) {
	data := mustReadFileString(t, "testdata/singlefirewallrule.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	rule, err := newClient.Firewall.DeleteFirewallRule(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, rule)

	rule, err = newClient.Firewall.DeleteFirewallRule(context.Background(), 1)
	require.Error(t, err)
	require.Nil(t, rule)

	rule, err = newClient.Firewall.DeleteFirewallRule(context.Background(), 1)
	require.Error(t, err)
	require.Nil(t, rule)
}

func TestFirewallService_ListFirewallAliases(t *testing.T) {
	data := mustReadFileString(t, "testdata/multiplefirewallalias.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	aliases, err := newClient.Firewall.ListFirewallAliases(context.Background())
	require.NoError(t, err)
	require.Len(t, aliases, 2)

	aliases, err = newClient.Firewall.ListFirewallAliases(context.Background())
	require.Error(t, err)
	require.Nil(t, aliases)

	aliases, err = newClient.Firewall.ListFirewallAliases(context.Background())
	require.Error(t, err)
	require.Nil(t, aliases)
}

func TestFirewallService_GetFirewallAlias(t *testing.T) {
	data := mustReadFileString(t, "testdata/singlefirewallalias.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	alias, err := newClient.Firewall.GetFirewallAlias(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, alias)
	require.Equal(t, 1, alias.Id)

	alias, err = newClient.Firewall.GetFirewallAlias(context.Background(), 1)
	require.Error(t, err)
	require.Nil(t, alias)

	alias, err = newClient.Firewall.GetFirewallAlias(context.Background(), 1)
	require.Error(t, err)
	require.Nil(t, alias)
}

func TestFirewallService_CreateFirewallAlias(t *testing.T) {
	data := mustReadFileString(t, "testdata/singlefirewallalias.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	descr := optional.NewString("Test Alias")

	newAlias := FirewallAliasRequest{
		Name:    "test_alias",
		Type:    "host",
		Descr:   &descr,
		Address: []string{"192.168.1.1", "192.168.1.2"},
		Detail:  []string{"Host 1", "Host 2"},
	}

	alias, err := newClient.Firewall.CreateFirewallAlias(context.Background(), newAlias)
	require.NoError(t, err)
	require.NotNil(t, alias)

	alias, err = newClient.Firewall.CreateFirewallAlias(context.Background(), newAlias)
	require.Error(t, err)
	require.Nil(t, alias)

	alias, err = newClient.Firewall.CreateFirewallAlias(context.Background(), newAlias)
	require.Error(t, err)
	require.Nil(t, alias)
}

func TestFirewallService_UpdateFirewallAlias(t *testing.T) {
	data := mustReadFileString(t, "testdata/singlefirewallalias.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	descr := optional.NewString("Updated Alias")

	updatedAlias := FirewallAliasRequest{
		Name:    "test_alias",
		Type:    "host",
		Descr:   &descr,
		Address: []string{"192.168.1.1", "192.168.1.2", "192.168.1.3"},
		Detail:  []string{"Host 1", "Host 2", "Host 3"},
	}

	alias, err := newClient.Firewall.UpdateFirewallAlias(context.Background(), 1, updatedAlias)
	require.NoError(t, err)
	require.NotNil(t, alias)

	alias, err = newClient.Firewall.UpdateFirewallAlias(context.Background(), 1, updatedAlias)
	require.Error(t, err)
	require.Nil(t, alias)

	alias, err = newClient.Firewall.UpdateFirewallAlias(context.Background(), 1, updatedAlias)
	require.Error(t, err)
	require.Nil(t, alias)
}

func TestFirewallService_DeleteFirewallAlias(t *testing.T) {
	data := mustReadFileString(t, "testdata/singlefirewallalias.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	alias, err := newClient.Firewall.DeleteFirewallAlias(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, alias)

	alias, err = newClient.Firewall.DeleteFirewallAlias(context.Background(), 1)
	require.Error(t, err)
	require.Nil(t, alias)

	alias, err = newClient.Firewall.DeleteFirewallAlias(context.Background(), 1)
	require.Error(t, err)
	require.Nil(t, alias)
}
