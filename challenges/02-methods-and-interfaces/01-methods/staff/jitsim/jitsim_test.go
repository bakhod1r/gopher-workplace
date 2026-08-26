package jitsim

import "testing"

func TestJIT(t *testing.T) {
	j := &JIT{}
	if got := j.Compile(); got != "compiled" {
		t.Errorf("got %q", got)
	}
}
