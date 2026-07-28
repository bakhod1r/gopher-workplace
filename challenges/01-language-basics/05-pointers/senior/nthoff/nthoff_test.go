package nthoff

import "testing"

func build(vs ...int) *Node {
	var h *Node
	for i := len(vs) - 1; i >= 0; i-- {
		h = &Node{Val: vs[i], Next: h}
	}
	return h
}

func TestNthFromEnd(t *testing.T) {
	l := build(1, 2, 3, 4, 5)
	if got := NthFromEnd(l, 1); got == nil || got.Val != 5 {
		t.Fatalf("n=1 want 5")
	}
	if got := NthFromEnd(l, 2); got == nil || got.Val != 4 {
		t.Fatalf("n=2 want 4")
	}
}
