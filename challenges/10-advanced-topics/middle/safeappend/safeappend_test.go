package safeappend

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	if got := Add([]int{1, 2}, 3); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("Add = %v, want [1 2 3]", got)
	}
	if got := Add(nil, 1); !reflect.DeepEqual(got, []int{1}) {
		t.Errorf("Add = %v, want [1]", got)
	}
}

func TestAddDoesNotClobberTheTail(t *testing.T) {
	backing := []int{1, 2, 3, 4}
	head := backing[:2]
	Add(head, 99)
	if backing[2] != 3 {
		t.Errorf("backing = %v, want [1 2 3 4]: the append wrote past the head", backing)
	}
}

func TestAddResultIsUsable(t *testing.T) {
	backing := []int{1, 2, 3, 4}
	got := Add(backing[:2], 99)
	if !reflect.DeepEqual(got, []int{1, 2, 99}) {
		t.Errorf("Add = %v, want [1 2 99]", got)
	}
}
