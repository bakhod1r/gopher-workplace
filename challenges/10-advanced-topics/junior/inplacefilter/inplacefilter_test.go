package inplacefilter

import (
	"reflect"
	"testing"
)

func TestKeepEven(t *testing.T) {
	if got := KeepEven([]int{1, 2, 3, 4, 6}); !reflect.DeepEqual(got, []int{2, 4, 6}) {
		t.Errorf("KeepEven = %v, want [2 4 6]", got)
	}
	if got := KeepEven([]int{1, 3}); len(got) != 0 {
		t.Errorf("KeepEven = %v, want empty", got)
	}
	if got := KeepEven(nil); len(got) != 0 {
		t.Errorf("KeepEven(nil) = %v, want empty", got)
	}
	if got := KeepEven([]int{-2, -1, 0}); !reflect.DeepEqual(got, []int{-2, 0}) {
		t.Errorf("KeepEven = %v, want [-2 0]", got)
	}
}

func TestKeepEvenAllocatesNothing(t *testing.T) {
	s := make([]int, 256)
	for i := range s {
		s[i] = i
	}
	if n := testing.AllocsPerRun(100, func() { _ = KeepEven(s) }); n != 0 {
		t.Errorf("KeepEven made %v allocations, want 0: filter in place", n)
	}
}
