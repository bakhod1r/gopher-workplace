package groupbybug

import (
	"reflect"
	"testing"
)

func parity(n int) int { return n % 2 }

func TestGroupByCollectsAll(t *testing.T) {
	got := GroupBy([]int{1, 2, 3, 4}, parity)
	if !reflect.DeepEqual(got[1], []int{1, 3}) {
		t.Errorf("group 1 = %v, want [1 3]", got[1])
	}
	if !reflect.DeepEqual(got[0], []int{2, 4}) {
		t.Errorf("group 0 = %v, want [2 4]", got[0])
	}
}

func TestGroupBySingle(t *testing.T) {
	got := GroupBy([]int{1}, parity)
	if !reflect.DeepEqual(got[1], []int{1}) {
		t.Errorf("group 1 = %v, want [1]", got[1])
	}
}

func TestGroupByEmpty(t *testing.T) {
	if got := GroupBy([]int{}, parity); len(got) != 0 {
		t.Errorf("GroupBy = %v, want empty", got)
	}
}
