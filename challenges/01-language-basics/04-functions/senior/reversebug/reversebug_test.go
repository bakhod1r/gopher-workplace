package reversebug

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	xs := []int{1, 2, 3, 4}
	Reverse(xs)
	if !reflect.DeepEqual(xs, []int{4, 3, 2, 1}) {
		t.Errorf("=%v want [4 3 2 1]", xs)
	}
	ys := []int{1, 2, 3}
	Reverse(ys)
	if !reflect.DeepEqual(ys, []int{3, 2, 1}) {
		t.Errorf("=%v want [3 2 1]", ys)
	}
}
