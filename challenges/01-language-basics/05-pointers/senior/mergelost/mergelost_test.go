package mergelost

import "testing"

func build(vs ...int) *Node {
	var h *Node
	for i := len(vs) - 1; i >= 0; i-- {
		h = &Node{Val: vs[i], Next: h}
	}
	return h
}

func TestMerge(t *testing.T) {
	h := Merge(build(1, 2), build(3, 4, 5))
	got := []int{}
	for ; h != nil; h = h.Next {
		got = append(got, h.Val)
	}
	want := []int{1, 2, 3, 4, 5}
	if len(got) != len(want) {
		t.Fatalf("=%v want %v (remainder dropped)", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("=%v want %v", got, want)
		}
	}
}
