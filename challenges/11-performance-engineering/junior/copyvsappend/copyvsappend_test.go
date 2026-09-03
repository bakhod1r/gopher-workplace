package copyvsappend

import (
	"reflect"
	"testing"
)

var sink []int

func TestMerge(t *testing.T) {
	if got := Merge([]int{1, 2}, []int{3}); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("Merge = %v, want [1 2 3]", got)
	}
	if got := Merge(nil, []int{3}); !reflect.DeepEqual(got, []int{3}) {
		t.Errorf("Merge = %v, want [3]", got)
	}
}

func TestMergeEmpty(t *testing.T) {
	got := Merge(nil, nil)
	if got == nil || len(got) != 0 {
		t.Errorf("Merge(nil, nil) = %v, want empty non-nil slice", got)
	}
}

func TestMergeSharesNothing(t *testing.T) {
	a := []int{1, 2}
	b := []int{3, 4}
	got := Merge(a, b)
	got[0] = 99
	got[3] = 99
	if a[0] != 1 || b[1] != 4 {
		t.Errorf("writing to the result changed an input: a = %v, b = %v", a, b)
	}
	got = append(got, 7)
	if len(b) != 2 || b[0] != 3 {
		t.Errorf("appending to the result changed b: %v", b)
	}
}

func TestMergeAllocatesOnce(t *testing.T) {
	a := make([]int, 500)
	b := make([]int, 500)
	allocs := testing.AllocsPerRun(50, func() { sink = Merge(a, b) })
	if allocs > 1 {
		t.Errorf("Merge made %v allocations, want at most 1 — both lengths are known", allocs)
	}
}
