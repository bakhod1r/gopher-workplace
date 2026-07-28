package nthfromend

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
	if got := NthFromEnd(l, 3); got == nil || got.Val != 3 {
		t.Fatalf("n=3 want 3")
	}
	if NthFromEnd(l, 9) != nil {
		t.Fatalf("too far -> nil")
	}
}
