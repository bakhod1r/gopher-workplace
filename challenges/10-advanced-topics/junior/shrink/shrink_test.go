package shrink

import (
	"reflect"
	"testing"
)

func TestShrinkReleasesTheSpare(t *testing.T) {
	s := make([]int, 2, 64)
	s[0], s[1] = 1, 2
	out := Shrink(s)
	if !reflect.DeepEqual(out, []int{1, 2}) {
		t.Errorf("contents = %v, want [1 2]", out)
	}
	if cap(out) != 2 {
		t.Errorf("cap = %d, want 2", cap(out))
	}
	if &out[0] == &s[0] {
		t.Error("the result still points at the oversized array")
	}
}

func TestShrinkKeepsATightSlice(t *testing.T) {
	s := make([]int, 8, 10)
	out := Shrink(s)
	if &out[0] != &s[0] {
		t.Error("a tight slice was copied for nothing")
	}
	if n := testing.AllocsPerRun(100, func() { _ = Shrink(s) }); n != 0 {
		t.Errorf("Shrink made %v allocations on a tight slice, want 0", n)
	}
}

func TestShrinkEmpty(t *testing.T) {
	if got := Shrink(nil); len(got) != 0 {
		t.Errorf("Shrink(nil) = %v, want empty", got)
	}
}
