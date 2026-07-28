package clampptr

import "testing"

func TestClamp(t *testing.T) {
	x := 99
	Clamp(&x, 0, 10)
	if x != 10 {
		t.Errorf("x=%d want 10", x)
	}
	y := -4
	Clamp(&y, 0, 10)
	if y != 0 {
		t.Errorf("y=%d want 0", y)
	}
}
