package copyshiftleft

import (
	"reflect"
	"testing"
)

func TestShiftLeft(t *testing.T) {
	xs := []int{1, 2, 3, 4}
	ShiftLeft(xs)
	if !reflect.DeepEqual(xs, []int{2, 3, 4, 0}) {
		t.Errorf("ShiftLeft=%v; want [2 3 4 0]", xs)
	}
}
