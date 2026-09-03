package prealloc

import (
	"reflect"
	"testing"
)

func TestSquares(t *testing.T) {
	if got := Squares(4); !reflect.DeepEqual(got, []int{0, 1, 4, 9}) {
		t.Errorf("Squares(4) = %v, want [0 1 4 9]", got)
	}
	if got := Squares(0); len(got) != 0 {
		t.Errorf("Squares(0) = %v, want empty", got)
	}
}

func TestSquaresAllocatesOnce(t *testing.T) {
	if n := testing.AllocsPerRun(100, func() { _ = Squares(64) }); n > 1 {
		t.Errorf("Squares made %v allocations, want 1: size the slice up front", n)
	}
}
