package sliceresetreuse

import (
	"reflect"
	"testing"
)

var sink []int

func TestResetKeepsCapacity(t *testing.T) {
	s := make([]int, 3, 64)
	got := Reset(s)
	if len(got) != 0 {
		t.Errorf("len = %d, want 0", len(got))
	}
	if cap(got) != 64 {
		t.Errorf("cap = %d, want 64 — the array must be reused", cap(got))
	}
}

func TestResetNil(t *testing.T) {
	got := Reset(nil)
	if got == nil || len(got) != 0 {
		t.Errorf("Reset(nil) = %v, want empty non-nil slice", got)
	}
}

func TestFillEvens(t *testing.T) {
	if got := FillEvens(nil, 3); !reflect.DeepEqual(got, []int{0, 2, 4}) {
		t.Errorf("FillEvens = %v, want [0 2 4]", got)
	}
	if got := FillEvens(nil, 0); len(got) != 0 {
		t.Errorf("FillEvens(nil, 0) = %v, want empty", got)
	}
}

func TestFillEvensOverwritesPreviousContent(t *testing.T) {
	buf := []int{9, 9, 9, 9, 9}
	got := FillEvens(buf, 2)
	if !reflect.DeepEqual(got, []int{0, 2}) {
		t.Errorf("FillEvens = %v, want [0 2]", got)
	}
}

func TestFillEvensReusesBuffer(t *testing.T) {
	buf := make([]int, 0, 1000)
	allocs := testing.AllocsPerRun(50, func() { sink = FillEvens(buf, 500) })
	if allocs != 0 {
		t.Errorf("FillEvens into a big enough buffer made %v allocations, want 0", allocs)
	}
}
