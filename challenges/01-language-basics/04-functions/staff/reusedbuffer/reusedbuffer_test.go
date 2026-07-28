package reusedbuffer

import (
	"reflect"
	"testing"
)

func TestReader(t *testing.T) {
	r := Reader()
	a := r(1, 2)
	b := r(3, 4)
	if !reflect.DeepEqual(a, []int{1, 2}) {
		t.Errorf("first result mutated: %v want [1 2]", a)
	}
	if !reflect.DeepEqual(b, []int{3, 4}) {
		t.Errorf("=%v want [3 4]", b)
	}
}
