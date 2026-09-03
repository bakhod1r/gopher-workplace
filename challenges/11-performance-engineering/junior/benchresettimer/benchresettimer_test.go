package benchresettimer

import "testing"

func TestMeasuredExcludesSetup(t *testing.T) {
	if got := Measured(1000, 7, 3); got != 21 {
		t.Errorf("Measured(1000, 7, 3) = %d, want 21 (setup must not count)", got)
	}
	if got := Measured(0, 7, 3); got != 21 {
		t.Errorf("Measured(0, 7, 3) = %d, want 21", got)
	}
}

func TestMeasuredNonPositiveN(t *testing.T) {
	for _, n := range []int64{0, -1} {
		if got := Measured(500, 7, n); got != 0 {
			t.Errorf("Measured(500, 7, %d) = %d, want 0", n, got)
		}
	}
}

func TestMeasuredScales(t *testing.T) {
	if got := Measured(1, 2, 1_000_000); got != 2_000_000 {
		t.Errorf("Measured = %d, want 2000000", got)
	}
}
