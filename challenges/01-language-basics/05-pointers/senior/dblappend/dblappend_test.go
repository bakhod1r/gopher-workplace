package dblappend

import "testing"

func TestExtend(t *testing.T) {
	var xs []int
	Extend(&xs, 1, 2, 3)
	if len(xs) != 3 {
		t.Errorf("len=%d want 3", len(xs))
	}
}
