package revlist

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

func TestReverse(t *testing.T) {
	got := slice(Reverse(build(1, 2, 3)))
	want := []int{3, 2, 1}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("=%v want %v", got, want)
		}
	}
	if Reverse(nil) != nil {
		t.Fatalf("nil should reverse to nil")
	}
}
