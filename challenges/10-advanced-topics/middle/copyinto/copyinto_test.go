package copyinto

import (
	"reflect"
	"testing"
)

func TestCopyIntoShortDst(t *testing.T) {
	dst := make([]int, 2)
	if got := CopyInto(dst, []int{1, 2, 3}); got != 2 {
		t.Errorf("CopyInto = %d, want 2", got)
	}
	if !reflect.DeepEqual(dst, []int{1, 2}) {
		t.Errorf("dst = %v, want [1 2]", dst)
	}
}

func TestCopyIntoShortSrc(t *testing.T) {
	dst := []int{9, 9, 9}
	if got := CopyInto(dst, []int{1}); got != 1 {
		t.Errorf("CopyInto = %d, want 1", got)
	}
	if !reflect.DeepEqual(dst, []int{1, 9, 9}) {
		t.Errorf("dst = %v, want [1 9 9]", dst)
	}
}

func TestCopyIntoEmpty(t *testing.T) {
	if got := CopyInto(nil, []int{1}); got != 0 {
		t.Errorf("CopyInto = %d, want 0", got)
	}
	if got := CopyInto(make([]int, 3), nil); got != 0 {
		t.Errorf("CopyInto = %d, want 0", got)
	}
	dst := make([]int, 0, 8)
	if got := CopyInto(dst, []int{1, 2}); got != 0 {
		t.Errorf("CopyInto = %d, want 0: capacity is not length", got)
	}
}

func TestCopyIntoOverlapping(t *testing.T) {
	s := []int{1, 2, 3, 4}
	if got := CopyInto(s, s[1:]); got != 3 {
		t.Errorf("CopyInto = %d, want 3", got)
	}
	if !reflect.DeepEqual(s, []int{2, 3, 4, 4}) {
		t.Errorf("s = %v, want [2 3 4 4]", s)
	}
}

func TestCopyIntoAllocatesNothing(t *testing.T) {
	dst := make([]int, 128)
	src := make([]int, 256)
	if n := testing.AllocsPerRun(100, func() { _ = CopyInto(dst, src) }); n != 0 {
		t.Errorf("CopyInto made %v allocations, want 0", n)
	}
}
