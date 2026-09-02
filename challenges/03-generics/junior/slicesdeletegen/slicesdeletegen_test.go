package slicesdeletegen

import (
	"reflect"
	"testing"
)

func TestRemoveAt(t *testing.T) {
	if got := RemoveAt([]int{1, 2, 3}, 1); !reflect.DeepEqual(got, []int{1, 3}) {
		t.Errorf("RemoveAt = %v, want [1 3]", got)
	}
	if got := RemoveAt([]int{1, 2}, 0); !reflect.DeepEqual(got, []int{2}) {
		t.Errorf("RemoveAt(0) = %v, want [2]", got)
	}
	if got := RemoveAt([]int{1}, 0); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("RemoveAt = %v, want []", got)
	}
}

func TestRemoveAtOutOfRange(t *testing.T) {
	if got := RemoveAt([]int{1}, 5); !reflect.DeepEqual(got, []int{1}) {
		t.Errorf("RemoveAt(5) = %v, want [1]", got)
	}
	if got := RemoveAt([]int{1}, -1); !reflect.DeepEqual(got, []int{1}) {
		t.Errorf("RemoveAt(-1) = %v, want [1]", got)
	}
	if got := RemoveAt([]int{}, 0); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("RemoveAt on empty = %v, want []", got)
	}
}

func TestRemoveAtDoesNotMutate(t *testing.T) {
	in := []int{1, 2, 3}
	RemoveAt(in, 1)
	if !reflect.DeepEqual(in, []int{1, 2, 3}) {
		t.Errorf("input mutated: %v, want [1 2 3]", in)
	}
}
