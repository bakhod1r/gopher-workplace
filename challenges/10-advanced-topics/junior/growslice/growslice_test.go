package growslice

import (
	"reflect"
	"testing"
)

func TestGrowCapacity(t *testing.T) {
	s := make([]int, 2, 2)
	s[0], s[1] = 1, 2
	out := Grow(s, 8)
	if len(out) != 2 {
		t.Errorf("len = %d, want 2: Grow must not change the length", len(out))
	}
	if cap(out) < 10 {
		t.Errorf("cap = %d, want at least 10", cap(out))
	}
	if !reflect.DeepEqual(out, []int{1, 2}) {
		t.Errorf("contents = %v, want [1 2]", out)
	}
}

func TestGrowIsANoOpWhenTheRoomExists(t *testing.T) {
	s := make([]int, 1, 32)
	if n := testing.AllocsPerRun(100, func() { _ = Grow(s, 4) }); n != 0 {
		t.Errorf("Grow made %v allocations, want 0 when the capacity already fits", n)
	}
}

func TestGrowNegative(t *testing.T) {
	s := make([]int, 1, 1)
	if out := Grow(s, -5); len(out) != 1 {
		t.Errorf("len = %d, want 1", len(out))
	}
}
