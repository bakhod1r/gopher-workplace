package sumifgen

import "testing"

func TestSumIf(t *testing.T) {
	isEven := func(n int) bool { return n%2 == 0 }
	if got := SumIf([]int{1, 2, 3, 4}, isEven); got != 6 {
		t.Errorf("SumIf([]int{1, 2, 3, 4}, isEven) = %v, want 6", got)
	}
	if got := SumIf([]int{1, 3}, isEven); got != 0 {
		t.Errorf("SumIf([]int{1, 3}, isEven) = %v, want 0", got)
	}
	isPositive := func(f float64) bool { return f > 0 }
	if got := SumIf([]float64{1.5, -1}, isPositive); got != 1.5 {
		t.Errorf("SumIf([]float64{1.5, -1}, isPositive) = %v, want 1.5", got)
	}
	if got := SumIf([]int{}, isEven); got != 0 {
		t.Errorf("SumIf([]int{}, isEven) = %v, want 0", got)
	}
}
