package listslice

import (
	"reflect"
	"testing"
)

func build(vs ...int) *Node {
	var h *Node
	for i := len(vs) - 1; i >= 0; i-- {
		h = &Node{Val: vs[i], Next: h}
	}
	return h
}

func TestToSlice(t *testing.T) {
	if got := ToSlice(build(1, 2, 3)); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("=%v", got)
	}
	if got := ToSlice(nil); len(got) != 0 {
		t.Errorf("nil -> empty")
	}
}
