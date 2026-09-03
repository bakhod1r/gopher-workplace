package insertboundbug

import (
	"reflect"
	"testing"
)

func TestInsertAtEnd(t *testing.T) {
	if got := InsertAt([]int{1}, 1, 2); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf("InsertAt at len = %v, want [1 2]", got)
	}
	if got := InsertAt([]int{}, 0, 1); !reflect.DeepEqual(got, []int{1}) {
		t.Errorf("InsertAt into empty = %v, want [1]", got)
	}
}

func TestInsertAtMiddle(t *testing.T) {
	if got := InsertAt([]int{1, 3}, 1, 2); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("InsertAt = %v, want [1 2 3]", got)
	}
	if got := InsertAt([]int{2, 3}, 0, 1); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("InsertAt = %v, want [1 2 3]", got)
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
