package pfsenseapi

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func mustReadFileString(t *testing.T, filename string) string {
	t.Helper()

	out, err := os.ReadFile(filename)

	require.NoError(t, err, "could not read file %q", filename)

	return string(out)
}
