package unionprecisionbug

import (
	"math"
	"testing"
	"time"
)

func TestMeanFloats(t *testing.T) {
	if got := Mean([]float64{0.5, 0.5, 0.5, 0.5}); math.Abs(got-0.5) > 1e-9 {
		t.Errorf("Mean = %v, want 0.5", got)
	}
}

func TestMeanInts(t *testing.T) {
	if got := Mean([]int{1, 2, 3, 4}); math.Abs(got-2.5) > 1e-9 {
		t.Errorf("Mean = %v, want 2.5", got)
	}
}

func TestMeanEmpty(t *testing.T) {
	if got := Mean([]int64{}); got != 0 {
		t.Errorf("Mean = %v, want 0", got)
	}
}

func TestMeanScaleFloats(t *testing.T) {
	const n = 3000000
	xs := make([]float64, n)
	for i := range xs {
		xs[i] = 0.25
	}
	start := time.Now()
	got := Mean(xs)
	if el := time.Since(start); el > 5*time.Second {
		t.Fatalf("Mean over %d elements took %v, want under 5s", n, el)
	}
	if math.Abs(got-0.25) > 1e-9 {
		t.Errorf("Mean = %v, want 0.25", got)
	}
}
