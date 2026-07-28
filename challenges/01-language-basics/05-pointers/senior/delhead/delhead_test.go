package delhead

import "testing"

func TestRemoveFirst(t *testing.T) {
	l := &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 3}}}
	h := RemoveFirst(l)
	if h == nil || h.Val != 2 {
		t.Fatalf("head=%v want 2", h)
	}
	if RemoveFirst(&Node{Val: 9}) != nil {
		t.Fatalf("single -> nil")
	}
}
