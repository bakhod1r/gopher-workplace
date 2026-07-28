package bsearchbug

import "testing"

func TestIndexOf(t *testing.T) {
	xs := []int{1, 3, 5, 7, 9}
	for i, v := range xs {
		if got := IndexOf(xs, v); got != i {
			t.Errorf("IndexOf(%d)=%d want %d", v, got, i)
		}
	}
	if got := IndexOf(xs, 4); got != -1 {
		t.Errorf("absent=%d want -1", got)
	}
}
