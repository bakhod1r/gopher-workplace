package slicesinsertgen

import (
	"reflect"
	"testing"
)

func TestInsertAt(t *testing.T) {
	if got := InsertAt([]int{1, 3}, 1, 2); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("InsertAt = %v, want [1 2 3]", got)
	}
	if got := InsertAt([]int{2, 3}, 0, 1); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("InsertAt at 0 = %v, want [1 2 3]", got)
	}
	if got := InsertAt([]int{1}, 1, 2); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf("InsertAt at len = %v, want [1 2] (appending is legal)", got)
	}
}

func TestInsertAtOutOfRange(t *testing.T) {
	if got := InsertAt([]int{1}, 5, 2); !reflect.DeepEqual(got, []int{1}) {
		t.Errorf("InsertAt(5) = %v, want [1]", got)
	}
	if got := InsertAt([]int{1}, -1, 2); !reflect.DeepEqual(got, []int{1}) {
		t.Errorf("InsertAt(-1) = %v, want [1]", got)
	}
}

func TestInsertAtDoesNotMutate(t *testing.T) {
	in := make([]int, 2, 8)
	in[0], in[1] = 1, 3
	InsertAt(in, 1, 2)
	if in[1] != 3 {
		t.Errorf("input mutated: %v, want [1 3]", in)
	}
}
