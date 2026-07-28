package deferwipe

import "testing"

func TestCompute(t *testing.T) {
	if got := Compute(6, 7); got != 42 {
		t.Errorf("=%d want 42 (defer wiped it?)", got)
	}
}
