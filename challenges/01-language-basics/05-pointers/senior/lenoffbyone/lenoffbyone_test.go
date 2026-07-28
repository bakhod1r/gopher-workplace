package lenoffbyone

import "testing"

func build(vs ...int) *Node {
	var h *Node
	for i := len(vs) - 1; i >= 0; i-- {
		h = &Node{Val: vs[i], Next: h}
	}
	return h
}

func TestLength(t *testing.T) {
	if Length(nil) != 0 {
		t.Errorf("nil should be 0")
	}
	if got := Length(build(1, 2, 3)); got != 3 {
		t.Errorf("=%d want 3", got)
	}
}
