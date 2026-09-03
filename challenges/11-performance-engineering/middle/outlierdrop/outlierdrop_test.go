package outlierdrop

import (
	"reflect"
	"testing"
)

func TestQuartiles(t *testing.T) {
	q1, q3 := Quartiles([]float64{1, 2, 3, 4})
	if q1 != 1 || q3 != 3 {
		t.Errorf("Quartiles = %v, %v, want 1, 3", q1, q3)
	}
	q1, q3 = Quartiles([]float64{4, 3, 2, 1})
	if q1 != 1 || q3 != 3 {
		t.Errorf("Quartiles on unsorted input = %v, %v, want 1, 3", q1, q3)
	}
	q1, q3 = Quartiles(nil)
	if q1 != 0 || q3 != 0 {
		t.Errorf("Quartiles(nil) = %v, %v, want 0, 0", q1, q3)
	}
}

func TestFilterDropsTheOutlier(t *testing.T) {
	got := Filter([]float64{1, 2, 3, 4, 1000}, 1.5)
	if !reflect.DeepEqual(got, []float64{1, 2, 3, 4}) {
		t.Errorf("Filter = %v, want [1 2 3 4]", got)
	}
}

func TestFilterKeepsEverythingWhenTheDataIsTight(t *testing.T) {
	in := []float64{10, 11, 12, 13}
	got := Filter(in, 1.5)
	if !reflect.DeepEqual(got, in) {
		t.Errorf("Filter = %v, want %v", got, in)
	}
}

func TestFilterPreservesOrder(t *testing.T) {
	got := Filter([]float64{4, 1, 1000, 3, 2}, 1.5)
	if !reflect.DeepEqual(got, []float64{4, 1, 3, 2}) {
		t.Errorf("Filter = %v, want [4 1 3 2]", got)
	}
}

func TestFilterDoesNotModifyInput(t *testing.T) {
	in := []float64{4, 1, 1000, 3, 2}
	before := append([]float64(nil), in...)
	Filter(in, 1.5)
	if !reflect.DeepEqual(in, before) {
		t.Errorf("input changed: %v, want %v", in, before)
	}
}

func TestFilterGuards(t *testing.T) {
	got := Filter(nil, 1.5)
	if got == nil || len(got) != 0 {
		t.Errorf("Filter(nil) = %v, want empty non-nil slice", got)
	}
	// k = 0 keeps only the interquartile range itself.
	got = Filter([]float64{1, 2, 3, 4}, 0)
	if !reflect.DeepEqual(got, []float64{1, 2, 3}) {
		t.Errorf("Filter(k=0) = %v, want [1 2 3]", got)
	}
	if got := Filter([]float64{1, 2, 3, 4}, -5); !reflect.DeepEqual(got, []float64{1, 2, 3}) {
		t.Errorf("Filter(k=-5) = %v, want the k=0 result", got)
	}
}
