package treeheight

import "testing"

func TestHeight(t *testing.T) {
	if Height(nil) != 0 {
		t.Fatalf("nil height 0")
	}
	root := &Tree{Val: 1,
		Left:  &Tree{Val: 2, Left: &Tree{Val: 4}},
		Right: &Tree{Val: 3}}
	if got := Height(root); got != 3 {
		t.Fatalf("=%d want 3", got)
	}
}
