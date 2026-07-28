package sizeofptr

import "testing"

func TestElemSize(t *testing.T) {
	var a [8]int32
	if got := ElemSize(&a); got != 4 {
		t.Errorf("=%d want 4 (int32), measured the pointer?", got)
	}
}
