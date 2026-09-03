package nanminmaxbug

import (
	"math"
	"testing"
	"time"
)

func TestSkipsLeadingNaN(t *testing.T) {
	mn, mx, ok := MinMaxSkipNaN([]float64{math.NaN(), 3, 1, 2})
	if !ok || mn != 1 || mx != 3 {
		t.Errorf("MinMaxSkipNaN = %v, %v, %v; want 1, 3, true", mn, mx, ok)
	}
}

func TestSkipsInteriorNaN(t *testing.T) {
	mn, mx, ok := MinMaxSkipNaN([]float64{5, math.NaN(), 2})
	if !ok || mn != 2 || mx != 5 {
		t.Errorf("MinMaxSkipNaN = %v, %v, %v; want 2, 5, true", mn, mx, ok)
	}
}

func TestAllNaN(t *testing.T) {
	if _, _, ok := MinMaxSkipNaN([]float64{math.NaN(), math.NaN()}); ok {
		t.Errorf("MinMaxSkipNaN ok = true, want false")
	}
}

func TestIntegerInstantiation(t *testing.T) {
	mn, mx, ok := MinMaxSkipNaN([]int{3, 1, 2})
	if !ok || mn != 1 || mx != 3 {
		t.Errorf("MinMaxSkipNaN = %v, %v, %v; want 1, 3, true", mn, mx, ok)
	}
}

func TestScaleWithNaN(t *testing.T) {
	const n = 2000000
	xs := make([]float64, n)
	for i := range xs {
		xs[i] = float64(i % 1000)
	}
	xs[0] = math.NaN()
	xs[n/2] = math.NaN()
	start := time.Now()
	mn, mx, ok := MinMaxSkipNaN(xs)
	if el := time.Since(start); el > 5*time.Second {
		t.Fatalf("MinMaxSkipNaN over %d elements took %v, want under 5s", n, el)
	}
	if !ok || mn != 0 || mx != 999 {
		t.Errorf("MinMaxSkipNaN = %v, %v, %v; want 0, 999, true", mn, mx, ok)
	}
}
