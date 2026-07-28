package shortcircuit

import "testing"

func TestValueOr(t *testing.T) {
	if got := ValueOr(nil, 5); got != 5 {
		t.Errorf("nil pointer should give default 5, got %d", got)
	}
	n := 7
	if got := ValueOr(&n, 5); got != 7 {
		t.Errorf("=%d want 7", got)
	}
	z := -1
	if got := ValueOr(&z, 5); got != 5 {
		t.Errorf("negative should give default, got %d", got)
	}
}
