package sliceheadercopy

import (
	"reflect"
	"testing"
)

func TestFillIsVisibleToTheCaller(t *testing.T) {
	s := []int{1, 2, 3}
	Fill(s, 7)
	if !reflect.DeepEqual(s, []int{7, 7, 7}) {
		t.Errorf("s = %v, want [7 7 7]", s)
	}
}

func TestFillEmptyAndNil(t *testing.T) {
	Fill(nil, 1)
	Fill([]int{}, 1)
}

func TestFillDoesNotAllocate(t *testing.T) {
	s := make([]int, 1000)
	if allocs := testing.AllocsPerRun(50, func() { Fill(s, 3) }); allocs != 0 {
		t.Errorf("Fill made %v allocations, want 0", allocs)
	}
}

func TestAppendLocalDoesNotChangeCallerLength(t *testing.T) {
	s := make([]int, 3, 8)
	AppendLocal(s, 7)
	if len(s) != 3 {
		t.Errorf("len(s) = %d, want 3 — the callee's header is a copy", len(s))
	}
}

func TestAppendLocalStillWritesIntoSpareCapacity(t *testing.T) {
	s := make([]int, 3, 8)
	AppendLocal(s, 7)
	if got := s[:4][3]; got != 7 {
		t.Errorf("s[:4][3] = %d, want 7 — the array is shared even though the header is not", got)
	}
}
