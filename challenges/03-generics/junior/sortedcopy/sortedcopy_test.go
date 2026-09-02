package sortedcopy

import (
	"reflect"
	"testing"
)

func TestSorted(t *testing.T) {
	if got := Sorted([]int{3, 1, 2}); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("Sorted([]int{3, 1, 2}) = %v, want [1 2 3]", got)
	}
	if got := Sorted([]string{"b", "a"}); !reflect.DeepEqual(got, []string{"a", "b"}) {
		t.Errorf("Sorted([]string{...}) = %v, want [a b]", got)
	}
	if got := Sorted([]int{}); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("Sorted([]int{}) = %v, want []", got)
	}
}

func TestSortedDoesNotMutate(t *testing.T) {
	in := []int{3, 1, 2}
	Sorted(in)
	if !reflect.DeepEqual(in, []int{3, 1, 2}) {
		t.Errorf("input mutated: %v, want [3 1 2]", in)
	}
}
