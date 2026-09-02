package slicescompact

import (
	"reflect"
	"testing"
)

func TestSquash(t *testing.T) {
	if got := Squash([]int{1, 1, 2, 2, 1}); !reflect.DeepEqual(got, []int{1, 2, 1}) {
		t.Errorf("Squash = %v, want [1 2 1] (only adjacent runs collapse)", got)
	}
	if got := Squash([]int{1, 2}); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf("Squash = %v, want [1 2]", got)
	}
	if got := Squash([]string{"a", "a", "a"}); !reflect.DeepEqual(got, []string{"a"}) {
		t.Errorf("Squash = %v, want [a]", got)
	}
	if got := Squash([]int{}); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("Squash([]int{}) = %v, want []", got)
	}
}

func TestSquashDoesNotMutate(t *testing.T) {
	in := []int{1, 1, 2}
	Squash(in)
	if !reflect.DeepEqual(in, []int{1, 1, 2}) {
		t.Errorf("input mutated: %v, want [1 1 2]", in)
	}
}
