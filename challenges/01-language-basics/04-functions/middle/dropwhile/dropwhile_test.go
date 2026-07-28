package dropwhile

import (
	"reflect"
	"testing"
)

func TestDropWhile(t *testing.T) {
	pos := func(x int) bool { return x > 0 }
	got := DropWhile([]int{1, 2, -1, 3}, pos)
	if !reflect.DeepEqual(got, []int{-1, 3}) {
		t.Errorf("=%v want [-1 3]", got)
	}
	if got := DropWhile([]int{1, 2}, pos); len(got) != 0 {
		t.Errorf("all dropped should be empty")
	}
}
