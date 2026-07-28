package bumpptr

import "testing"

func TestBump(t *testing.T) {
	x := 41
	Bump(&x)
	if x != 42 {
		t.Errorf("x=%d want 42", x)
	}
}
