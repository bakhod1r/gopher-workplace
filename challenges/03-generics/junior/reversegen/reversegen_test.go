package reversegen

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	if got := Reverse([]int{1, 2, 3}); !reflect.DeepEqual(got, []int{3, 2, 1}) {
		t.Errorf("Reverse([]int{1, 2, 3}) = %v, want [3 2 1]", got)
	}
	if got := Reverse([]string{"a", "b"}); !reflect.DeepEqual(got, []string{"b", "a"}) {
		t.Errorf("Reverse([]string{\"a\", \"b\"}) = %v, want [b a]", got)
	}
	if got := Reverse([]int{}); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("Reverse([]int{}) = %v, want []", got)
	}
}

func TestReverseDoesNotMutate(t *testing.T) {
	in := []int{1, 2, 3}
	Reverse(in)
	if !reflect.DeepEqual(in, []int{1, 2, 3}) {
		t.Errorf("input mutated: %v, want [1 2 3]", in)
	}
}
