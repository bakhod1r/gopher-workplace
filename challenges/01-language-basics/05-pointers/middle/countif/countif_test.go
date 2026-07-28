package countif

import "testing"

func build(vs ...int) *Node {
	var h *Node
	for i := len(vs) - 1; i >= 0; i-- {
		h = &Node{Val: vs[i], Next: h}
	}
	return h
}

func TestCountIf(t *testing.T) {
	even := func(n int) bool { return n%2 == 0 }
	if got := CountIf(build(1, 2, 3, 4, 5, 6), even); got != 3 {
		t.Errorf("=%d want 3", got)
	}
	if CountIf(nil, even) != 0 {
		t.Errorf("nil should be 0")
	}
}
