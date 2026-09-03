package bucketcounts

import (
	"reflect"
	"testing"
)

var bounds = []float64{1, 5, 10}

func TestIndex(t *testing.T) {
	cases := []struct {
		v    float64
		want int
	}{
		{0.5, 0},
		{1, 0},
		{1.001, 1},
		{5, 1},
		{5.001, 2},
		{10, 2},
		{10.001, 3},
		{1e9, 3},
		{-4, 0},
	}
	for _, c := range cases {
		if got := Index(bounds, c.v); got != c.want {
			t.Errorf("Index(%v) = %d, want %d", c.v, got, c.want)
		}
	}
}

func TestIndexEmptyBounds(t *testing.T) {
	if got := Index(nil, 5); got != 0 {
		t.Errorf("Index(nil, 5) = %d, want 0", got)
	}
}

func TestCounts(t *testing.T) {
	got := Counts([]float64{1, 5}, []float64{0.5, 3, 100})
	if !reflect.DeepEqual(got, []int64{1, 1, 1}) {
		t.Errorf("Counts = %v, want [1 1 1]", got)
	}
}

func TestCountsBoundaryValuesGoInTheLowerBucket(t *testing.T) {
	got := Counts(bounds, []float64{1, 5, 10})
	if !reflect.DeepEqual(got, []int64{1, 1, 1, 0}) {
		t.Errorf("Counts = %v, want [1 1 1 0] — a value equal to a bound is inside it", got)
	}
}

func TestCountsEmptySamples(t *testing.T) {
	got := Counts(bounds, nil)
	if !reflect.DeepEqual(got, []int64{0, 0, 0, 0}) {
		t.Errorf("Counts = %v, want four zeros", got)
	}
}

func TestCumulative(t *testing.T) {
	if got := Cumulative([]int64{1, 2, 3}); !reflect.DeepEqual(got, []int64{1, 3, 6}) {
		t.Errorf("Cumulative = %v, want [1 3 6]", got)
	}
	got := Cumulative(nil)
	if got == nil || len(got) != 0 {
		t.Errorf("Cumulative(nil) = %v, want empty non-nil slice", got)
	}
}

func TestCumulativeDoesNotModifyInput(t *testing.T) {
	in := []int64{1, 2, 3}
	before := append([]int64(nil), in...)
	Cumulative(in)
	if !reflect.DeepEqual(in, before) {
		t.Errorf("input changed: %v, want %v", in, before)
	}
}

func TestScalesToManyBounds(t *testing.T) {
	big := make([]float64, 0, 1000)
	for i := 1; i <= 1000; i++ {
		big = append(big, float64(i))
	}
	if got := Index(big, 500.5); got != 500 {
		t.Errorf("Index = %d, want 500", got)
	}
	if got := Index(big, 1000.5); got != 1000 {
		t.Errorf("Index = %d, want 1000", got)
	}
}
