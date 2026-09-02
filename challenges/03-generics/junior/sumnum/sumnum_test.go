package sumnum

import "testing"

func TestSum(t *testing.T) {
	if got := Sum([]int{1, 2, 3}); got != 6 {
		t.Errorf("Sum([]int{1, 2, 3}) = %v, want 6", got)
	}
	if got := Sum([]float64{0.5, 0.5}); got != 1 {
		t.Errorf("Sum([]float64{0.5, 0.5}) = %v, want 1", got)
	}
	if got := Sum([]int64{10, -4}); got != 6 {
		t.Errorf("Sum([]int64{10, -4}) = %v, want 6", got)
	}
	if got := Sum([]int{}); got != 0 {
		t.Errorf("Sum([]int{}) = %v, want 0", got)
	}
}
