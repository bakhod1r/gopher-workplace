package sumtreebug

import "testing"

func TestSumTree(t *testing.T) {
	root := &Tree{Val: 1, Left: &Tree{Val: 2}, Right: &Tree{Val: 3, Left: &Tree{Val: 4}}}
	if got := SumTree(root); got != 10 {
		t.Errorf("=%d want 10 (children skipped)", got)
	}
}
