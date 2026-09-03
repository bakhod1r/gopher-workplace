package appendto

import (
	"reflect"
	"testing"
)

func TestAppendSquares(t *testing.T) {
	if got := AppendSquares(nil, 4); !reflect.DeepEqual(got, []int{0, 1, 4, 9}) {
		t.Errorf("AppendSquares = %v, want [0 1 4 9]", got)
	}
	if got := AppendSquares([]int{7}, 2); !reflect.DeepEqual(got, []int{7, 0, 1}) {
		t.Errorf("AppendSquares = %v, want [7 0 1]", got)
	}
	if got := AppendSquares([]int{7}, 0); !reflect.DeepEqual(got, []int{7}) {
		t.Errorf("AppendSquares = %v, want [7]", got)
	}
	if got := AppendSquares(nil, -1); len(got) != 0 {
		t.Errorf("AppendSquares = %v, want empty", got)
	}
}

func TestAppendSquaresReusesCapacity(t *testing.T) {
	dst := make([]int, 0, 64)
	if n := testing.AllocsPerRun(100, func() { _ = AppendSquares(dst[:0], 64) }); n != 0 {
		t.Errorf("AppendSquares made %v allocations, want 0 when dst has room", n)
	}
}

func TestAppendSquaresDoesNotClobberBeyondLen(t *testing.T) {
	backing := make([]int, 4, 8)
	got := AppendSquares(backing[:1], 2)
	if !reflect.DeepEqual(got, []int{0, 0, 1}) {
		t.Errorf("got = %v, want [0 0 1]", got)
	}
}
