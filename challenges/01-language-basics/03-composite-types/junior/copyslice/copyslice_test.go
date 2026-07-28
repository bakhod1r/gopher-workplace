package copyslice

import (
	"reflect"
	"testing"
)

func TestClone(t *testing.T) {
	xs := []int{1, 2, 3}
	c := Clone(xs)
	if !reflect.DeepEqual(c, xs) {
		t.Errorf("Clone=%v; want %v", c, xs)
	}
	c[0] = 99
	if xs[0] != 1 {
		t.Error("clone shares backing array with original")
	}
	if Clone(nil) == nil {
		t.Error("clone of nil should be non-nil empty")
	}
}
