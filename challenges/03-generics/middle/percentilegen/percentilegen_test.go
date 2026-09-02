package percentilegen

import (
	"reflect"
	"testing"
)

func TestPercentile(t *testing.T) {
	s := []float64{1, 2, 3, 4}
	if v, ok := Percentile(s, 50); v != 2 || !ok {
		t.Errorf("p50 = %v, %v, want 2, true", v, ok)
	}
	if v, ok := Percentile(s, 100); v != 4 || !ok {
		t.Errorf("p100 = %v, %v, want 4, true", v, ok)
	}
	if v, ok := Percentile(s, 25); v != 1 || !ok {
		t.Errorf("p25 = %v, %v, want 1, true", v, ok)
	}
}

func TestPercentileClamps(t *testing.T) {
	s := []float64{1, 2, 3, 4}
	if v, ok := Percentile(s, 0); v != 1 || !ok {
		t.Errorf("p0 = %v, %v, want 1, true", v, ok)
	}
	if v, ok := Percentile(s, -5); v != 1 || !ok {
		t.Errorf("p-5 = %v, %v, want 1, true", v, ok)
	}
	if v, ok := Percentile(s, 500); v != 4 || !ok {
		t.Errorf("p500 = %v, %v, want 4, true", v, ok)
	}
}

func TestPercentileEmptyAndPurity(t *testing.T) {
	if v, ok := Percentile([]float64{}, 95); v != 0 || ok {
		t.Errorf("Percentile(empty) = %v, %v, want 0, false", v, ok)
	}
	in := []float64{3, 1, 2}
	Percentile(in, 50)
	if !reflect.DeepEqual(in, []float64{3, 1, 2}) {
		t.Errorf("input mutated: %v", in)
	}
}
