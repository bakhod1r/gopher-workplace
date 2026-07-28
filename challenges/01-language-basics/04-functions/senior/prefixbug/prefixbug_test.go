package prefixbug

import "testing"

func TestRangeSum(t *testing.T) {
	xs := []int{1, 2, 3, 4, 5}
	if got := RangeSum(xs, 1, 4); got != 9 {
		t.Errorf("=%d want 9 (2+3+4)", got)
	}
	if got := RangeSum(xs, 0, 5); got != 15 {
		t.Errorf("=%d want 15", got)
	}
}
