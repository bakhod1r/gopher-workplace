package reverseloop

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	in := []int{1, 2, 3}
	got := Reverse(in)
	if !reflect.DeepEqual(got, []int{3, 2, 1}) {
		t.Errorf("=%v want [3 2 1]", got)
	}
	if !reflect.DeepEqual(in, []int{1, 2, 3}) {
		t.Errorf("input mutated: %v", in)
	}
}
