package dedupesorted

import (
	"reflect"
	"testing"
)

func TestDedupe(t *testing.T) {
	if got := Dedupe([]int{1, 1, 2, 3, 3, 3}); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("Dedupe = %v, want [1 2 3]", got)
	}
	if got := Dedupe([]int{5}); !reflect.DeepEqual(got, []int{5}) {
		t.Errorf("Dedupe = %v, want [5]", got)
	}
	if got := Dedupe(nil); len(got) != 0 {
		t.Errorf("Dedupe(nil) = %v, want empty", got)
	}
	if got := Dedupe([]int{2, 2, 2}); !reflect.DeepEqual(got, []int{2}) {
		t.Errorf("Dedupe = %v, want [2]", got)
	}
}

func TestDedupeAllocatesNothing(t *testing.T) {
	s := make([]int, 512)
	for i := range s {
		s[i] = i / 2
	}
	if n := testing.AllocsPerRun(100, func() { _ = Dedupe(s) }); n != 0 {
		t.Errorf("Dedupe made %v allocations, want 0", n)
	}
}
