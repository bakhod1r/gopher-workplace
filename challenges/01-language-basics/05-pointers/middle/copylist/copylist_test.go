package copylist

import "testing"

func build(vs ...int) *Node {
	var h *Node
	for i := len(vs) - 1; i >= 0; i-- {
		h = &Node{Val: vs[i], Next: h}
	}
	return h
}

func TestCopy(t *testing.T) {
	orig := build(1, 2, 3)
	cp := Copy(orig)
	cp.Val = 99
	cp.Next.Val = 88
	if orig.Val != 1 || orig.Next.Val != 2 {
		t.Errorf("original mutated: %d,%d", orig.Val, orig.Next.Val)
	}
	if cp.Next.Next.Val != 3 {
		t.Errorf("copy incomplete")
	}
}
