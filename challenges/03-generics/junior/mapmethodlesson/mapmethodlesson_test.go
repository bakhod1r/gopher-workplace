package mapmethodlesson

import (
	"reflect"
	"strconv"
	"testing"
)

func TestMapSlice(t *testing.T) {
	double := func(n int) int { return n * 2 }
	if got := MapSlice(Slice[int]{1, 2}, double); !reflect.DeepEqual(got, Slice[int]{2, 4}) {
		t.Errorf("MapSlice = %v, want [2 4]", got)
	}
	if got := MapSlice(Slice[int]{1, 2}, strconv.Itoa); !reflect.DeepEqual(got, Slice[string]{"1", "2"}) {
		t.Errorf("MapSlice = %v, want [1 2]", got)
	}
	if got := MapSlice(Slice[int]{}, double); got == nil || len(got) != 0 {
		t.Errorf("MapSlice on empty = %v, want an empty non-nil Slice", got)
	}
}

func TestEach(t *testing.T) {
	var seen []int
	Slice[int]{1, 2, 3}.Each(func(n int) { seen = append(seen, n) })
	if !reflect.DeepEqual(seen, []int{1, 2, 3}) {
		t.Errorf("Each visited %v, want [1 2 3] in order", seen)
	}
	var none []int
	Slice[int]{}.Each(func(n int) { none = append(none, n) })
	if len(none) != 0 {
		t.Errorf("Each on empty called f %d times, want 0", len(none))
	}
}
