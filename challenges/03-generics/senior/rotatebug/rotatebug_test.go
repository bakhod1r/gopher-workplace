package rotatebug

import (
	"reflect"
	"testing"
)

func TestRotateNegative(t *testing.T) {
	if got := Rotate([]int{1, 2, 3}, -1); !reflect.DeepEqual(got, []int{3, 1, 2}) {
		t.Errorf("Rotate(-1) = %v, want [3 1 2]", got)
	}
	if got := Rotate([]int{1, 2, 3}, -4); !reflect.DeepEqual(got, []int{3, 1, 2}) {
		t.Errorf("Rotate(-4) = %v, want [3 1 2]", got)
	}
}

func TestRotatePositive(t *testing.T) {
	if got := Rotate([]int{1, 2, 3}, 1); !reflect.DeepEqual(got, []int{2, 3, 1}) {
		t.Errorf("Rotate(1) = %v, want [2 3 1]", got)
	}
	if got := Rotate([]int{1, 2, 3}, 4); !reflect.DeepEqual(got, []int{2, 3, 1}) {
		t.Errorf("Rotate(4) = %v, want [2 3 1]", got)
	}
	if got := Rotate([]int{1, 2, 3}, 0); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("Rotate(0) = %v, want [1 2 3]", got)
	}
}

func TestRotateEmpty(t *testing.T) {
	if got := Rotate([]int{}, -3); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("Rotate(empty) = %v, want []", got)
	}
}
