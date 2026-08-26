package gcswp

import "testing"

func TestGCSweep(t *testing.T) {
	h := &Heap{Objects: []bool{true, false, true}}
	if got := h.Sweep(); got != 1 {
		t.Errorf("got %d", got)
	}
}
