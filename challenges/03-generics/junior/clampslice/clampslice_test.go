package clampslice

import (
	"reflect"
	"testing"
)

func TestClampAll(t *testing.T) {
	if got := ClampAll([]int{-1, 2, 9}, 0, 3); !reflect.DeepEqual(got, []int{0, 2, 3}) {
		t.Errorf("ClampAll([]int{-1, 2, 9}, 0, 3) = %v, want [0 2 3]", got)
	}
	if got := ClampAll([]float64{5}, 0, 1); !reflect.DeepEqual(got, []float64{1}) {
		t.Errorf("ClampAll([]float64{5}, 0, 1) = %v, want [1]", got)
	}
	if got := ClampAll([]int{}, 0, 3); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("ClampAll([]int{}, 0, 3) = %v, want []", got)
	}
}

func TestClampAllDoesNotMutate(t *testing.T) {
	in := []int{-1, 9}
	ClampAll(in, 0, 3)
	if !reflect.DeepEqual(in, []int{-1, 9}) {
		t.Errorf("input mutated: %v, want [-1 9]", in)
	}
}
