package slicesindexfunc

import "testing"

func TestFirstNegative(t *testing.T) {
	if got := FirstNegative([]int{1, -2, -3}); got != 1 {
		t.Errorf("FirstNegative([]int{1, -2, -3}) = %d, want 1", got)
	}
	if got := FirstNegative([]int{-1}); got != 0 {
		t.Errorf("FirstNegative([]int{-1}) = %d, want 0", got)
	}
	if got := FirstNegative([]int{1, 2}); got != -1 {
		t.Errorf("FirstNegative([]int{1, 2}) = %d, want -1", got)
	}
	if got := FirstNegative([]int{}); got != -1 {
		t.Errorf("FirstNegative([]int{}) = %d, want -1", got)
	}
	if got := FirstNegative([]int{0}); got != -1 {
		t.Errorf("FirstNegative([]int{0}) = %d, want -1 (zero is not negative)", got)
	}
}
