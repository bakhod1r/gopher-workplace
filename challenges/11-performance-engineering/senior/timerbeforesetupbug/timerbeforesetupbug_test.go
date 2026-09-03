package timerbeforesetupbug

import "testing"

func TestMeasuredExcludesSetup(t *testing.T) {
	if got := Measured(1000, 7, 3); got != 21 {
		t.Errorf("Measured(1000, 7, 3) = %d, want 21", got)
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

func TestPerOpIsStableAsNGrows(t *testing.T) {
	// The whole point of excluding setup: ns/op must not depend on b.N.
	first := PerOp(1_000_000, 7, 10)
	for _, n := range []int64{100, 10_000, 1_000_000} {
		if got := PerOp(1_000_000, 7, n); got != first {
			t.Fatalf("PerOp with n=%d = %d, want %d — setup time is leaking into ns/op", n, got, first)
		}
	}
	if first != 7 {
		t.Errorf("PerOp = %d, want 7", first)
	}
}
