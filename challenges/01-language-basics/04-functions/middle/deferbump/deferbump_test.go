package deferbump

import "testing"

func TestCompute(t *testing.T) {
	if got := Compute(); got != 20 {
		t.Errorf("=%d want 20 (defer doubled the named result)", got)
	}
}
