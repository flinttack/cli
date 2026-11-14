// Package enginetest lets tests that need the analysis engine skip when it is absent.
package enginetest

import (
	"testing"

	"github.com/flinttack/cli/internal/engine"
)

// Require skips the test unless the engine executable can be located.
func Require(t testing.TB) {
	t.Helper()
	if !engine.Available() {
		t.Skip("flinttack-engine not available; set FLINTTACK_ENGINE to run this test")
	}
}
