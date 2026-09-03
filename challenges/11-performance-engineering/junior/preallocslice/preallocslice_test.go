package preallocslice

import (
	"reflect"
	"testing"
)

var sink []int

func TestSquares(t *testing.T) {
	if got := Squares(3); !reflect.DeepEqual(got, []int{1, 4, 9}) {
		t.Errorf("Squares(3) = %v, want [1 4 9]", got)
	}
	if got := Squares(1); !reflect.DeepEqual(got, []int{1}) {
		t.Errorf("Squares(1) = %v, want [1]", got)
	}
}

func TestSquaresEmpty(t *testing.T) {
	for _, n := range []int{0, -4} {
		got := Squares(n)
		if got == nil || len(got) != 0 {
			t.Errorf("Squares(%d) = %v, want empty non-nil slice", n, got)
		}
	}
}

func TestSquaresAllocatesOnce(t *testing.T) {
	allocs := testing.AllocsPerRun(50, func() { sink = Squares(1000) })
	if allocs > 1 {
		t.Errorf("Squares(1000) made %v allocations, want at most 1 — the length is known up front", allocs)
	}
}
