package hyperlogl

import "testing"

func TestHLL(t *testing.T) {
	h := &HLL{}
	h.Add(1) // %5 = 1
	h.Add(4) // %5 = 4
	h.Add(2) // %5 = 2

	if h.maxZeros != 4 {
		t.Errorf("maxZeros = %d, want 4", h.maxZeros)
	}
}
