package compactbug

import (
	"reflect"
	"testing"
)

func TestDistinctUnsorted(t *testing.T) {
	if got := Distinct([]int{3, 1, 3}); !reflect.DeepEqual(got, []int{1, 3}) {
		t.Errorf("Distinct = %v, want [1 3]", got)
	}
	if got := Distinct([]string{"b", "a", "b", "a"}); !reflect.DeepEqual(got, []string{"a", "b"}) {
		t.Errorf("Distinct = %v, want [a b]", got)
	}
}

func TestDistinctAdjacent(t *testing.T) {
	if got := Distinct([]int{1, 1, 2}); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf("Distinct = %v, want [1 2]", got)
	}
	if got := Distinct([]int{}); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("Distinct(empty) = %v, want []", got)
	}
}

func TestDistinctDoesNotMutate(t *testing.T) {
	in := []int{3, 1, 3}
	Distinct(in)
	if !reflect.DeepEqual(in, []int{3, 1, 3}) {
		t.Errorf("input mutated: %v, want [3 1 3]", in)
	}
}
