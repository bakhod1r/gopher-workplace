package latencyavg

import (
	"math"
	"testing"
)

func near(got, want float64) bool { return math.Abs(got-want) < 1e-9 }

func TestMean(t *testing.T) {
	if got := Mean([]float64{1, 2, 3}); !near(got, 2) {
		t.Errorf("Mean = %v, want 2", got)
	}
	if got := Mean([]float64{5}); !near(got, 5) {
		t.Errorf("Mean = %v, want 5", got)
	}
	if got := Mean(nil); got != 0 {
		t.Errorf("Mean(nil) = %v, want 0", got)
	}
}

func TestWeightedMean(t *testing.T) {
	if got := WeightedMean([]float64{10, 20}, []int{1, 3}); !near(got, 17.5) {
		t.Errorf("WeightedMean = %v, want 17.5", got)
	}
	if got := WeightedMean([]float64{10, 20}, []int{1, 1}); !near(got, 15) {
		t.Errorf("WeightedMean = %v, want 15", got)
	}
}

func TestWeightedMeanIsNotTheMeanOfMeans(t *testing.T) {
	// 1 slow request at 100ms, 999 fast ones at 1ms.
	got := WeightedMean([]float64{100, 1}, []int{1, 999})
	if !near(got, (100+999)/1000.0) {
		t.Errorf("WeightedMean = %v, want %v", got, (100+999)/1000.0)
	}
	if near(got, Mean([]float64{100, 1})) {
		t.Error("weighted mean equalled the unweighted mean of the two averages")
	}
}

func TestWeightedMeanGuards(t *testing.T) {
	if got := WeightedMean([]float64{10, 20}, []int{2}); !near(got, 10) {
		t.Errorf("WeightedMean with a short weights slice = %v, want 10", got)
	}
	if got := WeightedMean([]float64{10, 20}, []int{0, -1}); got != 0 {
		t.Errorf("WeightedMean with no positive weights = %v, want 0", got)
	}
	if got := WeightedMean(nil, nil); got != 0 {
		t.Errorf("WeightedMean(nil, nil) = %v, want 0", got)
	}
}
