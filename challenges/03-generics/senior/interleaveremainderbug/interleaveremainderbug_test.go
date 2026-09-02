package interleaveremainderbug

import (
	"reflect"
	"testing"
)

func TestInterleaveLongerFirst(t *testing.T) {
	got := Interleave([]int{1, 2, 3}, []int{9})
	if !reflect.DeepEqual(got, []int{1, 9, 2, 3}) {
		t.Errorf("Interleave = %v, want [1 9 2 3]", got)
	}
}

func TestInterleaveLongerSecond(t *testing.T) {
	got := Interleave([]int{1}, []int{8, 9})
	if !reflect.DeepEqual(got, []int{1, 8, 9}) {
		t.Errorf("Interleave = %v, want [1 8 9]", got)
	}
}

func TestInterleaveEmpty(t *testing.T) {
	if got := Interleave([]int{}, []int{}); len(got) != 0 {
		t.Errorf("Interleave = %v, want []", got)
	}
}
