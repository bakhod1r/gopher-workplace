package pairbug

import "testing"

func TestSumDiffs(t *testing.T) {
	if got := SumDiffs([]int{1, 4, 9}); got != 8 {
		t.Errorf("=%d want 8", got)
	}
	if got := SumDiffs([]int{5}); got != 0 {
		t.Errorf("single element should be 0, got %d", got)
	}
}
