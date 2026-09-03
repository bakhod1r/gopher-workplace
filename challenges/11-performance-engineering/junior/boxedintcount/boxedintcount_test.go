package boxedintcount

import (
	"reflect"
	"testing"
)

var sink []any

func TestBox(t *testing.T) {
	got := Box([]int{1, 2, 3})
	want := []any{1, 2, 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Box = %v, want %v", got, want)
	}
}

func TestBoxEmpty(t *testing.T) {
	got := Box(nil)
	if got == nil || len(got) != 0 {
		t.Errorf("Box(nil) = %v, want empty non-nil slice", got)
	}
}

func TestUnbox(t *testing.T) {
	got, skipped := Unbox([]any{1, "x", 3, nil})
	if !reflect.DeepEqual(got, []int{1, 3}) || skipped != 2 {
		t.Errorf("Unbox = %v, %d, want [1 3], 2", got, skipped)
	}
	got, skipped = Unbox(nil)
	if got == nil || len(got) != 0 || skipped != 0 {
		t.Errorf("Unbox(nil) = %v, %d, want empty non-nil slice and 0", got, skipped)
	}
}

func TestRoundTrip(t *testing.T) {
	in := []int{5, 6, 7}
	got, skipped := Unbox(Box(in))
	if !reflect.DeepEqual(got, in) || skipped != 0 {
		t.Errorf("round trip = %v, %d, want %v, 0", got, skipped, in)
	}
}

func TestBoxingSmallValuesIsFree(t *testing.T) {
	small := make([]int, 100)
	for i := range small {
		small[i] = i % 256
	}
	allocs := testing.AllocsPerRun(50, func() { sink = Box(small) })
	if allocs > 1 {
		t.Errorf("Box of 100 small ints made %v allocations, want at most 1 (the result slice)", allocs)
	}
}

func TestBoxingLargeValuesAllocatesPerElement(t *testing.T) {
	big := make([]int, 100)
	for i := range big {
		big[i] = 100000 + i
	}
	allocs := testing.AllocsPerRun(50, func() { sink = Box(big) })
	if allocs < 50 {
		t.Errorf("Box of 100 large ints made only %v allocations — each value must be boxed on the heap", allocs)
	}
}
