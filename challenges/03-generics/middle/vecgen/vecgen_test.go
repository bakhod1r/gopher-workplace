package vecgen

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	if got := Add([]int{1, 2}, []int{3, 4}); !reflect.DeepEqual(got, []int{4, 6}) {
		t.Errorf("Add = %v, want [4 6]", got)
	}
	if got := Add([]float64{0.5}, []float64{0.5}); !reflect.DeepEqual(got, []float64{1}) {
		t.Errorf("Add = %v, want [1]", got)
	}
	if got := Add([]int{1}, []int{1, 2}); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("Add(mismatched) = %v, want []", got)
	}
	if got := Add([]int{}, []int{}); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("Add(empty) = %v, want []", got)
	}
}

func TestDot(t *testing.T) {
	if v, ok := Dot([]int{1, 2}, []int{3, 4}); v != 11 || !ok {
		t.Errorf("Dot = %v, %v, want 11, true", v, ok)
	}
	if v, ok := Dot([]float64{0.5, 2}, []float64{2, 0.5}); v != 2 || !ok {
		t.Errorf("Dot = %v, %v, want 2, true", v, ok)
	}
	if v, ok := Dot([]int{1}, []int{1, 2}); v != 0 || ok {
		t.Errorf("Dot(mismatched) = %v, %v, want 0, false", v, ok)
	}
	if v, ok := Dot([]int{}, []int{}); v != 0 || !ok {
		t.Errorf("Dot(empty) = %v, %v, want 0, true", v, ok)
	}
}
