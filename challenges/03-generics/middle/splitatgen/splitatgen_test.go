package splitatgen

import (
	"reflect"
	"testing"
)

func TestSplitAt(t *testing.T) {
	l, r := SplitAt([]int{1, 2, 3}, 1)
	if !reflect.DeepEqual(l, []int{1}) || !reflect.DeepEqual(r, []int{2, 3}) {
		t.Errorf("SplitAt(1) = %v, %v, want [1], [2 3]", l, r)
	}
	l, r = SplitAt([]int{1, 2}, 0)
	if !reflect.DeepEqual(l, []int{}) || !reflect.DeepEqual(r, []int{1, 2}) {
		t.Errorf("SplitAt(0) = %v, %v, want [], [1 2]", l, r)
	}
	l, r = SplitAt([]int{1, 2}, 2)
	if !reflect.DeepEqual(l, []int{1, 2}) || !reflect.DeepEqual(r, []int{}) {
		t.Errorf("SplitAt(2) = %v, %v, want [1 2], []", l, r)
	}
}

func TestSplitAtClamps(t *testing.T) {
	l, r := SplitAt([]int{1, 2}, 9)
	if !reflect.DeepEqual(l, []int{1, 2}) || !reflect.DeepEqual(r, []int{}) {
		t.Errorf("SplitAt(9) = %v, %v, want [1 2], []", l, r)
	}
	l, r = SplitAt([]int{1, 2}, -1)
	if !reflect.DeepEqual(l, []int{}) || !reflect.DeepEqual(r, []int{1, 2}) {
		t.Errorf("SplitAt(-1) = %v, %v, want [], [1 2]", l, r)
	}
}

func TestSplitAtHalvesAreIndependent(t *testing.T) {
	in := []int{1, 2, 3, 4}
	l, r := SplitAt(in, 2)
	l = append(l, 99)
	if r[0] != 3 || in[2] != 3 {
		t.Errorf("halves share storage: l=%v r=%v in=%v", l, r, in)
	}
}
