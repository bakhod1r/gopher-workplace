package shiftleft

import (
	"reflect"
	"testing"
)

func TestShift(t *testing.T) {
	if got := Shift([]int{1, 2, 3, 4}, 2); !reflect.DeepEqual(got, []int{3, 4}) {
		t.Errorf("Shift = %v, want [3 4]", got)
	}
	if got := Shift([]int{1, 2, 3}, 0); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("Shift = %v, want [1 2 3]", got)
	}
	if got := Shift([]int{1, 2}, 9); len(got) != 0 {
		t.Errorf("Shift = %v, want empty", got)
	}
	if got := Shift([]int{1, 2}, -1); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf("Shift = %v, want [1 2]", got)
	}
	if got := Shift(nil, 1); len(got) != 0 {
		t.Errorf("Shift = %v, want empty", got)
	}
}

func TestShiftOverlapIsHandled(t *testing.T) {
	s := make([]int, 1000)
	for i := range s {
		s[i] = i
	}
	got := Shift(s, 1)
	for i, v := range got {
		if v != i+1 {
			t.Fatalf("got[%d] = %d, want %d: the overlapping ranges were not copied correctly", i, v, i+1)
		}
	}
}

func TestShiftAllocatesNothing(t *testing.T) {
	s := make([]int, 512)
	if n := testing.AllocsPerRun(100, func() { _ = Shift(s, 1) }); n != 0 {
		t.Errorf("Shift made %v allocations, want 0", n)
	}
}
