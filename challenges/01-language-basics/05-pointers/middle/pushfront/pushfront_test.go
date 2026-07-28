package pushfront

import "testing"

func toSlice(h *Node) []int {
	var out []int
	for ; h != nil; h = h.Next {
		out = append(out, h.Val)
	}
	return out
}

func TestPushFront(t *testing.T) {
	h := PushFront(nil, 3)
	h = PushFront(h, 2)
	h = PushFront(h, 1)
	got := toSlice(h)
	want := []int{1, 2, 3}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("=%v want %v", got, want)
		}
	}
}
