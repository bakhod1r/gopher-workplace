package ptrfieldsize

import "testing"

func TestSize(t *testing.T) {
	if got := Size(); got != 16 {
		t.Errorf("=%d want 16 (measured one pointer?)", got)
	}
}
