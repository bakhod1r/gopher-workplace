package sumlist

import "testing"

func build(vs ...int) *Node {
	var h *Node
	for i := len(vs) - 1; i >= 0; i-- {
		h = &Node{Val: vs[i], Next: h}
	}
	return h
}

func TestSumList(t *testing.T) {
	if SumList(nil) != 0 {
		t.Fatalf("nil")
	}
	if got := SumList(build(1, 2, 3, 4)); got != 10 {
		t.Fatalf("=%d want 10", got)
	}
}
