package setzero

import "testing"

func TestZero(t *testing.T) {
	x := 99
	Zero(&x)
	if x != 0 {
		t.Errorf("x=%d want 0", x)
	}
}
