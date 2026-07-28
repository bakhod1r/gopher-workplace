package sumtree

import "testing"

func TestSumTree(t *testing.T) {
	root := &Tree{Val: 1, Left: &Tree{Val: 2}, Right: &Tree{Val: 3}}
	if got := SumTree(root); got != 6 {
		t.Errorf("=%d want 6", got)
	}
	if SumTree(nil) != 0 {
		t.Errorf("nil 0")
	}
}
