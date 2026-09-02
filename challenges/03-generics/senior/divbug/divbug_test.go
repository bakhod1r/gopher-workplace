package divbug

import "testing"

func TestSumRangeOddCounts(t *testing.T) {
	if got := SumRange(1, 2); got != 3 {
		t.Errorf("SumRange(1, 2) = %v, want 3", got)
	}
	if got := SumRange(1, 4); got != 10 {
		t.Errorf("SumRange(1, 4) = %v, want 10", got)
	}
	if got := SumRange(2, 5); got != 14 {
		t.Errorf("SumRange(2, 5) = %v, want 14", got)
	}
}

func TestSumRangeAgainstLoop(t *testing.T) {
	for lo := -3; lo <= 3; lo++ {
		for hi := lo; hi <= lo+6; hi++ {
			want := 0
			for i := lo; i <= hi; i++ {
				want += i
			}
			if got := SumRange(lo, hi); got != want {
				t.Fatalf("SumRange(%d, %d) = %v, want %v", lo, hi, got, want)
			}
		}
	}
}

func TestSumRangeEmpty(t *testing.T) {
	if got := SumRange(5, 1); got != 0 {
		t.Errorf("SumRange(5, 1) = %v, want 0", got)
	}
}
