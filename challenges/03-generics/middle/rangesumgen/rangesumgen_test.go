package rangesumgen

import "testing"

func TestSumRange(t *testing.T) {
	if got := SumRange(1, 4); got != 10 {
		t.Errorf("SumRange(1, 4) = %v, want 10", got)
	}
	if got := SumRange(3, 3); got != 3 {
		t.Errorf("SumRange(3, 3) = %v, want 3", got)
	}
	if got := SumRange(0, 3); got != 6 {
		t.Errorf("SumRange(0, 3) = %v, want 6", got)
	}
	if got := SumRange(int64(1), int64(100)); got != 5050 {
		t.Errorf("SumRange(1, 100) = %v, want 5050", got)
	}
}

func TestSumRangeEmpty(t *testing.T) {
	if got := SumRange(5, 1); got != 0 {
		t.Errorf("SumRange(5, 1) = %v, want 0", got)
	}
}

func TestSumRangeNegatives(t *testing.T) {
	if got := SumRange(-2, 2); got != 0 {
		t.Errorf("SumRange(-2, 2) = %v, want 0", got)
	}
	if got := SumRange(-3, -1); got != -6 {
		t.Errorf("SumRange(-3, -1) = %v, want -6", got)
	}
}
