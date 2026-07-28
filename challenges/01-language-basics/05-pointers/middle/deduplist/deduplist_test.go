package deduplist

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

func TestDedup(t *testing.T) {
	if got := slice(Dedup(build(1, 1, 2, 3, 3, 3))); !eq(got, []int{1, 2, 3}) {
		t.Fatalf("=%v", got)
	}
	if Dedup(nil) != nil {
		t.Fatalf("nil")
	}
}
