package fillshared

import (
	"reflect"
	"testing"
)

func TestFill(t *testing.T) {
	s := []int{1, 2, 3}
	Fill(s, 7)
	if !reflect.DeepEqual(s, []int{7, 7, 7}) {
		t.Errorf("s = %v, want [7 7 7]", s)
	}
}

func TestFillWritesThroughAView(t *testing.T) {
	s := []int{1, 2, 3, 4}
	Fill(s[1:3], 0)
	if !reflect.DeepEqual(s, []int{1, 0, 0, 4}) {
		t.Errorf("s = %v, want [1 0 0 4]: only the view's range may be written", s)
	}
}

func TestFillEmpty(t *testing.T) {
	Fill(nil, 1)
	Fill([]int{}, 1)
}

func TestFillAllocatesNothing(t *testing.T) {
	s := make([]int, 256)
	if n := testing.AllocsPerRun(100, func() { Fill(s, 3) }); n != 0 {
		t.Errorf("Fill made %v allocations, want 0", n)
	}
}
