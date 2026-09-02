package insertsortedgen

import (
	"reflect"
	"testing"
)

func TestInsertSorted(t *testing.T) {
	if got := InsertSorted([]int{1, 3}, 2); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("InsertSorted = %v, want [1 2 3]", got)
	}
	if got := InsertSorted([]int{2, 3}, 1); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("InsertSorted = %v, want [1 2 3]", got)
	}
	if got := InsertSorted([]int{1, 2}, 3); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("InsertSorted = %v, want [1 2 3]", got)
	}
	if got := InsertSorted([]int{}, 1); !reflect.DeepEqual(got, []int{1}) {
		t.Errorf("InsertSorted = %v, want [1]", got)
	}
}

func TestInsertSortedAfterEquals(t *testing.T) {
	got := InsertSorted([]int{1, 2, 2, 3}, 2)
	if !reflect.DeepEqual(got, []int{1, 2, 2, 2, 3}) {
		t.Errorf("InsertSorted = %v, want [1 2 2 2 3]", got)
	}
}

func TestInsertSortedDoesNotMutate(t *testing.T) {
	in := make([]int, 2, 8)
	in[0], in[1] = 1, 3
	InsertSorted(in, 2)
	if in[1] != 3 {
		t.Errorf("input mutated: %v, want [1 3]", in)
	}
}
