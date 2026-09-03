package percentilebug

import (
	"reflect"
	"testing"
)

func TestPercentileZero(t *testing.T) {
	s := []float64{4, 1, 3, 2}
	if v, ok := Percentile(s, 0); v != 1 || !ok {
		t.Errorf("p0 = %v, %v, want 1, true (must not panic)", v, ok)
	}
	if v, ok := Percentile(s, -10); v != 1 || !ok {
		t.Errorf("p-10 = %v, %v, want 1, true", v, ok)
	}
}

func TestPercentileNormal(t *testing.T) {
	s := []float64{1, 2, 3, 4}
	if v, ok := Percentile(s, 50); v != 2 || !ok {
		t.Errorf("p50 = %v, %v, want 2, true", v, ok)
	}
	if v, ok := Percentile(s, 100); v != 4 || !ok {
		t.Errorf("p100 = %v, %v, want 4, true", v, ok)
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
