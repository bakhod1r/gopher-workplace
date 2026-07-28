package negate

import "testing"

func TestNegate(t *testing.T) {
	x := 5
	Negate(&x)
	if x != -5 {
		t.Errorf("x=%d want -5", x)
	}
	Negate(&x)
	if x != 5 {
		t.Errorf("x=%d want 5", x)
	}
}
