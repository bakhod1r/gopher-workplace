package variancegen

import "testing"

func TestVariance(t *testing.T) {
	if got := Variance([]float64{2, 4}); got != 1 {
		t.Errorf("Variance([2 4]) = %v, want 1", got)
	}
	if got := Variance([]float64{5, 5, 5}); got != 0 {
		t.Errorf("Variance([5 5 5]) = %v, want 0", got)
	}
	if got := Variance([]float64{1, 2, 3, 4}); got != 1.25 {
		t.Errorf("Variance([1 2 3 4]) = %v, want 1.25 (population variance)", got)
	}
}

func TestVarianceSmallInputs(t *testing.T) {
	if got := Variance([]float64{1}); got != 0 {
		t.Errorf("Variance(one sample) = %v, want 0", got)
	}
	if got := Variance([]float64{}); got != 0 {
		t.Errorf("Variance(empty) = %v, want 0", got)
	}
	if got := Variance([]float32{2, 4}); got != 1 {
		t.Errorf("Variance(float32) = %v, want 1", got)
	}
}
