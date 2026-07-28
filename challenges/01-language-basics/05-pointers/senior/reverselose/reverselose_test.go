package reverselose

import "testing"

func TestReverse(t *testing.T) {
	l := &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 3}}}
	h := Reverse(l)
	got := []int{}
	for ; h != nil; h = h.Next {
		got = append(got, h.Val)
	}
	want := []int{3, 2, 1}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("=%v want %v", got, want)
		}
	}
}
