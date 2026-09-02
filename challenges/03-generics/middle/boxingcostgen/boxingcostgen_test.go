package boxingcostgen

import "testing"

func TestSumTyped(t *testing.T) {
	if got := SumTyped([]int{1, 2, 3}); got != 6 {
		t.Errorf("SumTyped = %v, want 6", got)
	}
	if got := SumTyped([]float64{0.5, 0.5}); got != 1 {
		t.Errorf("SumTyped = %v, want 1", got)
	}
	if got := SumTyped([]int{}); got != 0 {
		t.Errorf("SumTyped(empty) = %v, want 0", got)
	}
}

func TestSumAny(t *testing.T) {
	if v, ok := SumAny([]any{1, 2}); v != 3 || !ok {
		t.Errorf("SumAny = %v, %v, want 3, true", v, ok)
	}
	if v, ok := SumAny([]any{1, "x"}); v != 0 || ok {
		t.Errorf("SumAny = %v, %v, want 0, false (must not panic)", v, ok)
	}
	if v, ok := SumAny([]any{}); v != 0 || !ok {
		t.Errorf("SumAny(empty) = %v, %v, want 0, true", v, ok)
	}
	if v, ok := SumAny([]any{1.5}); v != 0 || ok {
		t.Errorf("SumAny(float) = %v, %v, want 0, false", v, ok)
	}
}
