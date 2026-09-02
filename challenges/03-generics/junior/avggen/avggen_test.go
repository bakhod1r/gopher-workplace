package avggen

import "testing"

func TestAverage(t *testing.T) {
	if got := Average([]int{1, 2, 3}); got != 2 {
		t.Errorf("Average([]int{1, 2, 3}) = %v, want 2", got)
	}
	if got := Average([]int{1, 2}); got != 1.5 {
		t.Errorf("Average([]int{1, 2}) = %v, want 1.5 (do not sum in T)", got)
	}
	if got := Average([]float64{1, 2}); got != 1.5 {
		t.Errorf("Average([]float64{1, 2}) = %v, want 1.5", got)
	}
	if got := Average([]int{}); got != 0 {
		t.Errorf("Average([]int{}) = %v, want 0", got)
	}
}
