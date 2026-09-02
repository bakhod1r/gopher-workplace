package scaleintgen

import (
	"reflect"
	"testing"
)

func TestScale(t *testing.T) {
	if got := Scale([]int{1, 2, 4}, 100); !reflect.DeepEqual(got, []int{25, 50, 100}) {
		t.Errorf("Scale = %v, want [25 50 100]", got)
	}
	if got := Scale([]int{5}, 10); !reflect.DeepEqual(got, []int{10}) {
		t.Errorf("Scale = %v, want [10]", got)
	}
	if got := Scale([]int{1, 3}, 9); !reflect.DeepEqual(got, []int{3, 9}) {
		t.Errorf("Scale = %v, want [3 9] (multiply before dividing)", got)
	}
}

func TestScaleEdges(t *testing.T) {
	if got := Scale([]int{0, 0}, 10); !reflect.DeepEqual(got, []int{0, 0}) {
		t.Errorf("Scale(all zero) = %v, want [0 0]", got)
	}
	if got := Scale([]int{}, 10); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("Scale(empty) = %v, want []", got)
	}
	if got := Scale([]int{-1, -2}, 10); !reflect.DeepEqual(got, []int{-1, -2}) {
		t.Errorf("Scale(all negative) = %v, want the input unchanged", got)
	}
}
