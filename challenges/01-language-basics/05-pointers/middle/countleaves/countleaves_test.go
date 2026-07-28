package countleaves

import "testing"

func TestCountLeaves(t *testing.T) {
	root := &Tree{Val: 1,
		Left:  &Tree{Val: 2, Left: &Tree{Val: 4}, Right: &Tree{Val: 5}},
		Right: &Tree{Val: 3}}
	if got := CountLeaves(root); got != 3 {
		t.Errorf("=%d want 3", got)
	}
	if CountLeaves(nil) != 0 {
		t.Errorf("nil 0")
	}
}
