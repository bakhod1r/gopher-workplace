package reverseinplace

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	xs := []int{1, 2, 3, 4}
	Reverse(xs)
	if !reflect.DeepEqual(xs, []int{4, 3, 2, 1}) {
		t.Errorf("in place got %v", xs)
	}
	if !reflect.DeepEqual(Reverse([]int{1, 2, 3}), []int{3, 2, 1}) {
		t.Error("odd length")
	}
	if len(Reverse(nil)) != 0 {
		t.Error("nil ok")
	}
}
