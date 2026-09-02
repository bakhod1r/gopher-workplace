package nansafegen

import (
	"math"
	"testing"
)

func TestMinIgnoringNaN(t *testing.T) {
	nan := math.NaN()
	if v, ok := MinIgnoringNaN([]float64{3, nan, 1}); v != 1 || !ok {
		t.Errorf("MinIgnoringNaN = %v, %v, want 1, true", v, ok)
	}
	if v, ok := MinIgnoringNaN([]float64{nan, 2}); v != 2 || !ok {
		t.Errorf("MinIgnoringNaN = %v, %v, want 2, true (a leading NaN must not poison the seed)", v, ok)
	}
	if v, ok := MinIgnoringNaN([]float64{5, 4, 6}); v != 4 || !ok {
		t.Errorf("MinIgnoringNaN = %v, %v, want 4, true", v, ok)
	}
}

func TestMinIgnoringNaNNoRealValues(t *testing.T) {
	nan := math.NaN()
	if v, ok := MinIgnoringNaN([]float64{nan, nan}); v != 0 || ok {
		t.Errorf("MinIgnoringNaN(all NaN) = %v, %v, want 0, false", v, ok)
	}
	if v, ok := MinIgnoringNaN([]float64{}); v != 0 || ok {
		t.Errorf("MinIgnoringNaN(empty) = %v, %v, want 0, false", v, ok)
	}
}
