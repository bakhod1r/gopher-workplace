package nanbug

import (
	"math"
	"testing"
)

func TestLeadingNaN(t *testing.T) {
	nan := math.NaN()
	if v, ok := MinIgnoringNaN([]float64{nan, 2}); v != 2 || !ok {
		t.Errorf("MinIgnoringNaN = %v, %v, want 2, true (a leading NaN must not poison the seed)", v, ok)
	}
	if v, ok := MinIgnoringNaN([]float64{nan, 5, 3}); v != 3 || !ok {
		t.Errorf("MinIgnoringNaN = %v, %v, want 3, true", v, ok)
	}
}

func TestInnerNaN(t *testing.T) {
	nan := math.NaN()
	if v, ok := MinIgnoringNaN([]float64{3, nan, 1}); v != 1 || !ok {
		t.Errorf("MinIgnoringNaN = %v, %v, want 1, true", v, ok)
	}
}

func TestNoRealValues(t *testing.T) {
	nan := math.NaN()
	if v, ok := MinIgnoringNaN([]float64{nan, nan}); v != 0 || ok {
		t.Errorf("MinIgnoringNaN(all NaN) = %v, %v, want 0, false", v, ok)
	}
	if v, ok := MinIgnoringNaN([]float64{}); v != 0 || ok {
		t.Errorf("MinIgnoringNaN(empty) = %v, %v, want 0, false", v, ok)
	}
}
