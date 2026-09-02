package countpositive

import "testing"

func TestCountPositive(t *testing.T) {
	if got := CountPositive([]int{-1, 0, 2}); got != 1 {
		t.Errorf("CountPositive([]int{-1, 0, 2}) = %v, want 1", got)
	}
	if got := CountPositive([]float64{0.5, -0.5}); got != 1 {
		t.Errorf("CountPositive([]float64{0.5, -0.5}) = %v, want 1", got)
	}
	if got := CountPositive([]int{0, 0}); got != 0 {
		t.Errorf("CountPositive([]int{0, 0}) = %v, want 0 (zero is not positive)", got)
	}
	if got := CountPositive([]int{}); got != 0 {
		t.Errorf("CountPositive([]int{}) = %v, want 0", got)
	}
}
