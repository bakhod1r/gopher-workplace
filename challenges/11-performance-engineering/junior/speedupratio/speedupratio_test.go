package speedupratio

import (
	"math"
	"testing"
)

func near(got, want float64) bool { return math.Abs(got-want) < 1e-9 }

func TestSpeedup(t *testing.T) {
	cases := []struct{ base, cand, want float64 }{
		{100, 25, 4},
		{100, 100, 1},
		{100, 200, 0.5},
		{0, 25, 0},
		{100, 0, 0},
		{-1, 25, 0},
	}
	for _, c := range cases {
		if got := Speedup(c.base, c.cand); !near(got, c.want) {
			t.Errorf("Speedup(%v, %v) = %v, want %v", c.base, c.cand, got, c.want)
		}
	}
}

func TestPercentChange(t *testing.T) {
	cases := []struct{ base, cand, want float64 }{
		{100, 80, -20},
		{100, 125, 25},
		{100, 100, 0},
		{100, 0, -100},
		{0, 80, 0},
	}
	for _, c := range cases {
		if got := PercentChange(c.base, c.cand); !near(got, c.want) {
			t.Errorf("PercentChange(%v, %v) = %v, want %v", c.base, c.cand, got, c.want)
		}
	}
}

func TestSpeedupAndPercentAgree(t *testing.T) {
	// A 2x speedup is a 50% reduction, not a 200% one.
	if got := Speedup(100, 50); !near(got, 2) {
		t.Errorf("Speedup = %v, want 2", got)
	}
	if got := PercentChange(100, 50); !near(got, -50) {
		t.Errorf("PercentChange = %v, want -50", got)
	}
}
