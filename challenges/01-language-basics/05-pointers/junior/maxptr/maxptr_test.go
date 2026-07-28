package maxptr

import "testing"

func TestMaxPtr(t *testing.T) {
	x, y := 3, 8
	if MaxPtr(&x, &y) != &y {
		t.Errorf("want &y")
	}
	if MaxPtr(&y, &x) != &y {
		t.Errorf("want &y")
	}
	if MaxPtr(&x, &x) != &x {
		t.Errorf("tie should be first")
	}
}
