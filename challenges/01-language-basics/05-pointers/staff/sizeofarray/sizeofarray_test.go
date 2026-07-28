package sizeofarray

import "testing"

func TestTotalSize(t *testing.T) {
	var a [4]int32
	if got := TotalSize(&a); got != 16 {
		t.Errorf("=%d want 16 (measured one element?)", got)
	}
}
