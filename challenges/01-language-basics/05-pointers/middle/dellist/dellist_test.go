package dellist

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

func TestDelete(t *testing.T) {
	if got := slice(Delete(build(1, 2, 3), 2)); !eq(got, []int{1, 3}) {
		t.Fatalf("mid: %v", got)
	}
	if got := slice(Delete(build(1, 2, 3), 1)); !eq(got, []int{2, 3}) {
		t.Fatalf("head: %v", got)
	}
	if got := slice(Delete(build(1, 2, 3), 9)); !eq(got, []int{1, 2, 3}) {
		t.Fatalf("absent: %v", got)
	}
}
