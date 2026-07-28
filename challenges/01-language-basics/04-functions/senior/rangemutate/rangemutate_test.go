package rangemutate

import (
	"reflect"
	"testing"
)

func TestDoubleAll(t *testing.T) {
	xs := []int{1, 2, 3}
	DoubleAll(xs)
	if !reflect.DeepEqual(xs, []int{2, 4, 6}) {
		t.Errorf("=%v want [2 4 6]", xs)
	}
}
