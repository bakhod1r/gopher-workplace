package slicetypegen

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	isEven := func(n int) bool { return n%2 == 0 }
	got := Slice[int]{1, 2, 3, 4}.Filter(isEven)
	if !reflect.DeepEqual(got, Slice[int]{2, 4}) {
		t.Errorf("Filter = %v, want [2 4]", got)
	}
	if got.Len() != 2 {
		t.Errorf("Len() = %d, want 2", got.Len())
	}
	empty := Slice[int]{}.Filter(isEven)
	if empty == nil || len(empty) != 0 {
		t.Errorf("Filter on empty = %v, want an empty non-nil Slice", empty)
	}
}

func TestFilterChaining(t *testing.T) {
	n := Slice[string]{"", "a", "bb"}.Filter(func(s string) bool { return s != "" }).Len()
	if n != 2 {
		t.Errorf("chained Len() = %d, want 2", n)
	}
}

func TestFilterDoesNotMutate(t *testing.T) {
	in := Slice[int]{1, 2, 3}
	in.Filter(func(n int) bool { return n == 2 })
	if !reflect.DeepEqual(in, Slice[int]{1, 2, 3}) {
		t.Errorf("receiver mutated: %v, want [1 2 3]", in)
	}
}
