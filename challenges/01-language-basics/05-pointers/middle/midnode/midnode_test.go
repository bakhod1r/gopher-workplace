package midnode

import "testing"

func build(vs ...int) *Node {
	var h *Node
	for i := len(vs) - 1; i >= 0; i-- {
		h = &Node{Val: vs[i], Next: h}
	}
	return h
}

func TestMiddle(t *testing.T) {
	if m := Middle(build(1, 2, 3, 4, 5)); m == nil || m.Val != 3 {
		t.Fatalf("odd middle wrong")
	}
	if m := Middle(build(1, 2, 3, 4)); m == nil || m.Val != 3 {
		t.Fatalf("even middle should be 3")
	}
	if Middle(nil) != nil {
		t.Fatalf("nil")
	}
}
