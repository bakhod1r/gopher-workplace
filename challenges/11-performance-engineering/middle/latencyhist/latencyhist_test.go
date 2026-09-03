package latencyhist

import (
	"reflect"
	"testing"
)

func TestHistogram(t *testing.T) {
	got := Histogram([]float64{0, 5, 15}, 10, 1)
	if !reflect.DeepEqual(got, []int64{2, 1}) {
		t.Errorf("Histogram = %v, want [2 1]", got)
	}
}

func TestHistogramBucketBoundaries(t *testing.T) {
	// Buckets are [0,10), [10,20), [20,30), overflow at 30 and above.
	got := Histogram([]float64{0, 9.99, 10, 19.99, 20, 29.99, 30, 1000}, 10, 3)
	if !reflect.DeepEqual(got, []int64{2, 2, 2, 2}) {
		t.Errorf("Histogram = %v, want [2 2 2 2]", got)
	}
}

func TestHistogramDropsNegative(t *testing.T) {
	got := Histogram([]float64{-1, -0.001, 5}, 10, 2)
	if !reflect.DeepEqual(got, []int64{1, 0, 0}) {
		t.Errorf("Histogram = %v, want [1 0 0]", got)
	}
}

func TestHistogramGuards(t *testing.T) {
	for _, c := range []struct {
		width float64
		n     int
	}{{0, 3}, {-1, 3}, {10, 0}, {10, -2}} {
		got := Histogram([]float64{1}, c.width, c.n)
		if got == nil || len(got) != 0 {
			t.Errorf("Histogram(_, %v, %d) = %v, want empty non-nil slice", c.width, c.n, got)
		}
	}
}

func TestBusiest(t *testing.T) {
	if got := Busiest([]int64{1, 5, 5}); got != 1 {
		t.Errorf("Busiest = %d, want 1", got)
	}
	if got := Busiest([]int64{0, 0}); got != -1 {
		t.Errorf("Busiest = %d, want -1", got)
	}
	if got := Busiest(nil); got != -1 {
		t.Errorf("Busiest(nil) = %d, want -1", got)
	}
	if got := Busiest([]int64{9}); got != 0 {
		t.Errorf("Busiest = %d, want 0", got)
	}
}
