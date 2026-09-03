package amdahlbound

import (
	"math"
	"testing"
)

func near(got, want float64) bool { return math.Abs(got-want) < 1e-9 }

func TestMaxSpeedup(t *testing.T) {
	cases := []struct{ p, s, want float64 }{
		{0.5, 2, 1 / (0.5 + 0.25)},
		{0.9, 10, 1 / (0.1 + 0.09)},
		{1, 4, 4},
		{0, 100, 1},
		{0.5, 1, 1},
	}
	for _, c := range cases {
		if got := MaxSpeedup(c.p, c.s); !near(got, c.want) {
			t.Errorf("MaxSpeedup(%v, %v) = %v, want %v", c.p, c.s, got, c.want)
		}
	}
}

func TestMaxSpeedupInvalidInputs(t *testing.T) {
	for _, c := range []struct{ p, s float64 }{{-0.1, 2}, {1.5, 2}, {0.5, 0.5}, {0.5, 0}} {
		if got := MaxSpeedup(c.p, c.s); got != 1 {
			t.Errorf("MaxSpeedup(%v, %v) = %v, want 1", c.p, c.s, got)
		}
	}
}

func TestTinyFractionBarelyHelps(t *testing.T) {
	// 10x on 2% of the runtime is worth under 2%.
	got := MaxSpeedup(0.02, 10)
	if got > 1.02 {
		t.Errorf("MaxSpeedup(0.02, 10) = %v, want at most 1.02", got)
	}
}

func TestCeiling(t *testing.T) {
	if got := Ceiling(0.9); !near(got, 10) {
		t.Errorf("Ceiling(0.9) = %v, want 10", got)
	}
	if got := Ceiling(0.5); !near(got, 2) {
		t.Errorf("Ceiling(0.5) = %v, want 2", got)
	}
	if got := Ceiling(0); !near(got, 1) {
		t.Errorf("Ceiling(0) = %v, want 1", got)
	}
	if got := Ceiling(1); !math.IsInf(got, 1) {
		t.Errorf("Ceiling(1) = %v, want +Inf", got)
	}
}
