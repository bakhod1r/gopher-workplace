package warmupcountedbug

import (
	"math"
	"testing"
)

func near(got, want float64) bool { return math.Abs(got-want) < 1e-9 }

func TestStableMeanDropsTheWarmup(t *testing.T) {
	if got := StableMean([]float64{100, 2, 4}, 1); !near(got, 3) {
		t.Errorf("StableMean = %v, want 3 — the first sample is the warmup", got)
	}
}

func TestStableMeanDropsSeveral(t *testing.T) {
	if got := StableMean([]float64{100, 90, 2, 4}, 2); !near(got, 3) {
		t.Errorf("StableMean = %v, want 3", got)
	}
}

func TestStableMeanZeroWarmup(t *testing.T) {
	if got := StableMean([]float64{100, 2, 4}, 0); !near(got, 106.0/3) {
		t.Errorf("StableMean = %v, want %v", got, 106.0/3)
	}
	if got := StableMean([]float64{100, 2, 4}, -5); !near(got, 106.0/3) {
		t.Errorf("StableMean = %v, want %v", got, 106.0/3)
	}
}

func TestStableMeanDroppingEverything(t *testing.T) {
	if got := StableMean([]float64{1, 2}, 2); got != 0 {
		t.Errorf("StableMean = %v, want 0", got)
	}
	if got := StableMean([]float64{1, 2}, 9); got != 0 {
		t.Errorf("StableMean = %v, want 0", got)
	}
}

func TestStableMeanKeepsTheSteadyState(t *testing.T) {
	// One slow first sample, then a stable 5.
	samples := []float64{500, 5, 5, 5, 5, 5}
	if got := StableMean(samples, 1); !near(got, 5) {
		t.Errorf("StableMean = %v, want 5", got)
	}
}
