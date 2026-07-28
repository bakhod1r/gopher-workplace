package concatlist

import "testing"

func build(vs ...int) *Node {
	var h *Node
	for i := len(vs) - 1; i >= 0; i-- {
		h = &Node{Val: vs[i], Next: h}
	}
	return h
}

func slice(h *Node) []int {
	var o []int
	for ; h != nil; h = h.Next {
		o = append(o, h.Val)
	}
	return o
}

func eq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestConcat(t *testing.T) {
	if got := slice(Concat(build(1, 2), build(3, 4))); !eq(got, []int{1, 2, 3, 4}) {
		t.Fatalf("=%v", got)
	}
	if got := slice(Concat(nil, build(9))); !eq(got, []int{9}) {
		t.Fatalf("nil a: %v", got)
	}
}
