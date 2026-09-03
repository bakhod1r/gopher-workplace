package concatalias

import (
	"reflect"
	"testing"
)

func TestConcatJoins(t *testing.T) {
	got := Concat([]int{1}, []int{2, 3})
	if !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("Concat = %v, want [1 2 3]", got)
	}
}

func TestConcatSingleDoesNotAlias(t *testing.T) {
	src := []int{1, 2}
	got := Concat(src)
	got[0] = 99
	if src[0] != 1 {
		t.Errorf("input mutated: %v, want [1 2]", src)
	}
}

func TestConcatNoArgs(t *testing.T) {
	if got := Concat[int](); len(got) != 0 {
		t.Errorf("Concat = %v, want []", got)
	}
}
