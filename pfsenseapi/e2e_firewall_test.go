// +build e2e

package pfsenseapi

import (
	"context"
	"testing"

	"github.com/markphelps/optional"
	"github.com/stretchr/testify/require"
)

// TestE2E_FirewallRule_CRUD tests the full CRUD lifecycle of a firewall rule
func TestE2E_FirewallRule_CRUD(t *testing.T) {
	config, err := LoadE2EConfig()
	require.NoError(t, err, "Failed to load e2e config")

	client := NewClientWithLocalAuth(config.URL, config.Username, config.Password)
	ctx := context.Background()

	// Create a test firewall rule
	t.Log("Creating firewall rule...")
	ruleType := optional.NewString("pass")
	ipProto := optional.NewString("inet")
	proto := optional.NewString("tcp")
	src := optional.NewString("any")
	dst := optional.NewString("any")
	dstPort := optional.NewString("8443")
	descr := optional.NewString("E2E Test Rule - Safe to Delete")
	log := optional.NewBool(true)
	disabled := optional.NewBool(true) // Create disabled to avoid interfering with actual firewall
	top := optional.NewBool(false)

	newRule := FirewallRuleRequest{
		Type:            &ruleType,
		Interface:       []string{"lan"},
		IPProtocol:      &ipProto,
		Protocol:        &proto,
		Source:          &src,
		Destination:     &dst,
		DestinationPort: &dstPort,
		Descr:           &descr,
		Log:             &log,
		Disabled:        &disabled,
		Top:             &top,
	}

	createdRule, err := client.Firewall.CreateFirewallRule(ctx, newRule)
	require.NoError(t, err, "Failed to create firewall rule")
	require.NotNil(t, createdRule)
	require.Greater(t, createdRule.Id, 0)
	t.Logf("Created rule with ID: %d", createdRule.Id)

	// Ensure cleanup
	defer func() {
		t.Log("Cleaning up: deleting firewall rule...")
		_, err := client.Firewall.DeleteFirewallRule(ctx, createdRule.Id)
		if err != nil {
			t.Logf("Warning: Failed to delete rule %d: %v", createdRule.Id, err)
		} else {
			t.Logf("Successfully deleted rule %d", createdRule.Id)
		}
	}()

	// Read the rule back
	t.Log("Reading firewall rule...")
	readRule, err := client.Firewall.GetFirewallRule(ctx, createdRule.Id)
	require.NoError(t, err, "Failed to get firewall rule")
	require.NotNil(t, readRule)
	require.Equal(t, createdRule.Id, readRule.Id)

	descrValue, err := readRule.Descr.Get()
	require.NoError(t, err)
	require.Equal(t, "E2E Test Rule - Safe to Delete", descrValue)

	// Update the rule
	t.Log("Updating firewall rule...")
	updatedDescr := optional.NewString("E2E Test Rule - Updated - Safe to Delete")
	updatedDstPort := optional.NewString("9443")

	updateRule := FirewallRuleRequest{
		Type:            &ruleType,
		Interface:       []string{"lan"},
		IPProtocol:      &ipProto,
		Protocol:        &proto,
		Source:          &src,
		Destination:     &dst,
		DestinationPort: &updatedDstPort,
		Descr:           &updatedDescr,
		Log:             &log,
		Disabled:        &disabled,
		Top:             &top,
	}

	updatedRule, err := client.Firewall.UpdateFirewallRule(ctx, createdRule.Id, updateRule)
	require.NoError(t, err, "Failed to update firewall rule")
	require.NotNil(t, updatedRule)

	updatedDescrValue, err := updatedRule.Descr.Get()
	require.NoError(t, err)
	require.Equal(t, "E2E Test Rule - Updated - Safe to Delete", updatedDescrValue)

	updatedPortValue, err := updatedRule.DestinationPort.Get()
	require.NoError(t, err)
	require.Equal(t, "9443", updatedPortValue)

	t.Log("Firewall rule CRUD test completed successfully")
}

// TestE2E_FirewallRules_List tests listing firewall rules
func TestE2E_FirewallRules_List(t *testing.T) {
	config, err := LoadE2EConfig()
	require.NoError(t, err, "Failed to load e2e config")

	client := NewClientWithLocalAuth(config.URL, config.Username, config.Password)
	ctx := context.Background()

	t.Log("Listing firewall rules...")
	rules, err := client.Firewall.ListFirewallRules(ctx)
	require.NoError(t, err, "Failed to list firewall rules")
	require.NotNil(t, rules)

	t.Logf("Found %d firewall rules", len(rules))

	// Verify structure of returned rules
	for i, rule := range rules {
		require.Greater(t, rule.Id, -1, "Rule %d should have a valid ID", i)
		t.Logf("Rule %d: ID=%d", i, rule.Id)
	}
}

// TestE2E_FirewallAlias_CRUD tests the full CRUD lifecycle of a firewall alias
func TestE2E_FirewallAlias_CRUD(t *testing.T) {
	config, err := LoadE2EConfig()
	require.NoError(t, err, "Failed to load e2e config")

	client := NewClientWithLocalAuth(config.URL, config.Username, config.Password)
	ctx := context.Background()

	// Create a test firewall alias
	t.Log("Creating firewall alias...")
	descr := optional.NewString("E2E Test Alias - Safe to Delete")

	newAlias := FirewallAliasRequest{
		Name:    "e2e_test_alias",
		Type:    "host",
		Descr:   &descr,
		Address: []string{"192.168.100.1", "192.168.100.2"},
		Detail:  []string{"Test Host 1", "Test Host 2"},
	}

	createdAlias, err := client.Firewall.CreateFirewallAlias(ctx, newAlias)
	require.NoError(t, err, "Failed to create firewall alias")
	require.NotNil(t, createdAlias)
	require.Greater(t, createdAlias.Id, -1)
	t.Logf("Created alias with ID: %d", createdAlias.Id)

	// Ensure cleanup
	defer func() {
		t.Log("Cleaning up: deleting firewall alias...")
		_, err := client.Firewall.DeleteFirewallAlias(ctx, createdAlias.Id)
		if err != nil {
			t.Logf("Warning: Failed to delete alias %d: %v", createdAlias.Id, err)
		} else {
			t.Logf("Successfully deleted alias %d", createdAlias.Id)
		}
	}()

	// Read the alias back
	t.Log("Reading firewall alias...")
	readAlias, err := client.Firewall.GetFirewallAlias(ctx, createdAlias.Id)
	require.NoError(t, err, "Failed to get firewall alias")
	require.NotNil(t, readAlias)
	require.Equal(t, createdAlias.Id, readAlias.Id)
	require.Equal(t, "e2e_test_alias", readAlias.Name)

	// Update the alias
	t.Log("Updating firewall alias...")
	updatedDescr := optional.NewString("E2E Test Alias - Updated - Safe to Delete")

	updateAlias := FirewallAliasRequest{
		Name:    "e2e_test_alias",
		Type:    "host",
		Descr:   &updatedDescr,
		Address: []string{"192.168.100.1", "192.168.100.2", "192.168.100.3"},
		Detail:  []string{"Test Host 1", "Test Host 2", "Test Host 3"},
	}

	updatedAlias, err := client.Firewall.UpdateFirewallAlias(ctx, createdAlias.Id, updateAlias)
	require.NoError(t, err, "Failed to update firewall alias")
	require.NotNil(t, updatedAlias)
	require.Len(t, updatedAlias.Address, 3)

	t.Log("Firewall alias CRUD test completed successfully")
}

// TestE2E_FirewallAliases_List tests listing firewall aliases
func TestE2E_FirewallAliases_List(t *testing.T) {
	config, err := LoadE2EConfig()
	require.NoError(t, err, "Failed to load e2e config")

	client := NewClientWithLocalAuth(config.URL, config.Username, config.Password)
	ctx := context.Background()

	t.Log("Listing firewall aliases...")
	aliases, err := client.Firewall.ListFirewallAliases(ctx)
	require.NoError(t, err, "Failed to list firewall aliases")
	require.NotNil(t, aliases)

	t.Logf("Found %d firewall aliases", len(aliases))

	// Verify structure of returned aliases
	for i, alias := range aliases {
		require.Greater(t, alias.Id, 0, "Alias %d should have a valid ID", i)
		require.NotEmpty(t, alias.Name, "Alias %d should have a name", i)
		t.Logf("Alias %d: ID=%d, Name=%s", i, alias.Id, alias.Name)
	}
}
