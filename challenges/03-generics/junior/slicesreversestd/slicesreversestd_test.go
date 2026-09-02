package slicesreversestd

import (
	"reflect"
	"testing"
)

func TestReverseCopy(t *testing.T) {
	in := []int{1, 2, 3}
	got := ReverseCopy(in)
	if !reflect.DeepEqual(got, []int{3, 2, 1}) {
		t.Errorf("ReverseCopy = %v, want [3 2 1]", got)
	}
	if !reflect.DeepEqual(in, []int{1, 2, 3}) {
		t.Errorf("input mutated: %v, want [1 2 3]", in)
	}
}

func TestReverseCopyEdges(t *testing.T) {
	if got := ReverseCopy([]int{}); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("ReverseCopy([]int{}) = %v, want []", got)
	}
	got := ReverseCopy([]string(nil))
	if got == nil {
		t.Error("ReverseCopy(nil) = nil, want an empty non-nil slice")
	}
	if len(got) != 0 {
		t.Errorf("ReverseCopy(nil) = %v, want []", got)
	}
}
