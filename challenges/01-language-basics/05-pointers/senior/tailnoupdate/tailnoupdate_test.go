package tailnoupdate

import "testing"

func TestAppend(t *testing.T) {
	var h *Node
	h = Append(h, 1)
	h = Append(h, 2)
	if h == nil || h.Val != 1 || h.Next == nil || h.Next.Val != 2 {
		t.Fatalf("append to empty failed: %v", h)
	}
}
