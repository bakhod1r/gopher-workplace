package skipbug

import "testing"

func TestSumPositive(t *testing.T) {
	if got := SumPositive([]int{1, -2, 3, -4, 5}); got != 9 {
		t.Errorf("=%d want 9", got)
	}
	if got := SumPositive([]int{-1, -2}); got != 0 {
		t.Errorf("=%d want 0", got)
	}
}
