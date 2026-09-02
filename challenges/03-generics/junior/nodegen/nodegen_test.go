package nodegen

import (
	"reflect"
	"testing"
)

func TestPrependAndToSlice(t *testing.T) {
	var head *Node[int]
	head = Prepend(head, 2)
	head = Prepend(head, 1)
	if got := ToSlice(head); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf("ToSlice = %v, want [1 2]", got)
	}
}

func TestToSliceNil(t *testing.T) {
	if got := ToSlice[int](nil); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("ToSlice(nil) = %v, want []", got)
	}
}

func TestPrependKeepsOldList(t *testing.T) {
	old := Prepend[string](nil, "b")
	_ = Prepend(old, "a")
	if got := ToSlice(old); !reflect.DeepEqual(got, []string{"b"}) {
		t.Errorf("old list changed: %v, want [b]", got)
	}
}
