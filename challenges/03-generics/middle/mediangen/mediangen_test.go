package mediangen

import (
	"reflect"
	"testing"
)

func TestMedian(t *testing.T) {
	if v, ok := Median([]float64{3, 1, 2}); v != 2 || !ok {
		t.Errorf("Median = %v, %v, want 2, true", v, ok)
	}
	if v, ok := Median([]float64{1, 2, 3, 4}); v != 2.5 || !ok {
		t.Errorf("Median = %v, %v, want 2.5, true", v, ok)
	}
	if v, ok := Median([]float64{5}); v != 5 || !ok {
		t.Errorf("Median = %v, %v, want 5, true", v, ok)
	}
	if v, ok := Median([]float64{}); v != 0 || ok {
		t.Errorf("Median(empty) = %v, %v, want 0, false", v, ok)
	}
}

func TestMedianDoesNotMutate(t *testing.T) {
	in := []float64{3, 1, 2}
	Median(in)
	if !reflect.DeepEqual(in, []float64{3, 1, 2}) {
		t.Errorf("input mutated: %v, want [3 1 2]", in)
	}
}
