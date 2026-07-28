package appendtail

import "testing"

func slice(h *Node) []int {
	var o []int
	for ; h != nil; h = h.Next {
		o = append(o, h.Val)
	}
	return o
}

func TestAppend(t *testing.T) {
	var h *Node
	Append(&h, 1)
	Append(&h, 2)
	Append(&h, 3)
	got := slice(h)
	want := []int{1, 2, 3}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("=%v want %v", got, want)
		}
	}
}
