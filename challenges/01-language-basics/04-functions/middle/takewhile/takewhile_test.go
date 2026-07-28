package takewhile

import (
	"reflect"
	"testing"
)

func TestTakeWhile(t *testing.T) {
	pos := func(x int) bool { return x > 0 }
	got := TakeWhile([]int{1, 2, 3, -1, 4}, pos)
	if !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("=%v want [1 2 3]", got)
	}
	if got := TakeWhile([]int{-1, 2}, pos); len(got) != 0 {
		t.Errorf("=%v want empty", got)
	}
}
