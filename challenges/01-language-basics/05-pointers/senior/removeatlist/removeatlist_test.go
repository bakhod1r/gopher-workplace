package removeatlist

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

func TestRemoveAt(t *testing.T) {
	got := slice(RemoveAt(build(10, 20, 30, 40), 2))
	want := []int{10, 20, 40}
	if len(got) != 3 {
		t.Fatalf("=%v want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("=%v want %v", got, want)
		}
	}
}
