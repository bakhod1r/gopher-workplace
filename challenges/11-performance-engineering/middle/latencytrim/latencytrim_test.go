package latencytrim

import (
	"math"
	"reflect"
	"testing"
)

func near(got, want float64) bool { return math.Abs(got-want) < 1e-9 }

func TestTrimmedMean(t *testing.T) {
	// n=5, 20% trims floor(1) from each end: [1 100] go, [2 3 4] stay.
	if got := TrimmedMean([]float64{1, 2, 3, 4, 100}, 20); !near(got, 3) {
		t.Errorf("TrimmedMean = %v, want 3", got)
	}
}

func TestTrimmedMeanSortsFirst(t *testing.T) {
	if got := TrimmedMean([]float64{100, 3, 1, 4, 2}, 20); !near(got, 3) {
		t.Errorf("TrimmedMean = %v, want 3", got)
	}
}

func TestTrimmedMeanResistsOutliers(t *testing.T) {
	s := []float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 10000}
	plain := Mean(s)
	trimmed := TrimmedMean(s, 10)
	if !near(trimmed, 1) {
		t.Errorf("TrimmedMean = %v, want 1", trimmed)
	}
	if plain < 100 {
		t.Errorf("Mean = %v, expected the outlier to dominate", plain)
	}
}

func TestTrimmedMeanZeroPctIsThePlainMean(t *testing.T) {
	s := []float64{1, 2, 3, 4, 100}
	if !near(TrimmedMean(s, 0), Mean(s)) {
		t.Errorf("TrimmedMean(0) = %v, want %v", TrimmedMean(s, 0), Mean(s))
	}
}

func TestTrimmedMeanGuards(t *testing.T) {
	if got := TrimmedMean(nil, 10); got != 0 {
		t.Errorf("TrimmedMean(nil) = %v, want 0", got)
	}
	if got := TrimmedMean([]float64{1, 2}, -5); !near(got, 1.5) {
		t.Errorf("TrimmedMean(-5) = %v, want 1.5", got)
	}
	// Clamped below 50%, so at least one sample always survives.
	if got := TrimmedMean([]float64{1, 2, 3}, 99); got == 0 {
		t.Error("TrimmedMean(99) trimmed everything; pct must be clamped below 50")
	}
}

func TestTrimmedMeanDoesNotModifyInput(t *testing.T) {
	in := []float64{100, 3, 1, 4, 2}
	before := append([]float64(nil), in...)
	TrimmedMean(in, 20)
	if !reflect.DeepEqual(in, before) {
		t.Errorf("input changed: %v, want %v", in, before)
	}
}

func TestMean(t *testing.T) {
	if got := Mean([]float64{1, 2, 3, 4, 100}); !near(got, 22) {
		t.Errorf("Mean = %v, want 22", got)
	}
	if got := Mean(nil); got != 0 {
		t.Errorf("Mean(nil) = %v, want 0", got)
	}
}
