package mapints

import (
	"reflect"
	"testing"
)

func TestMapInts(t *testing.T) {
	got := MapInts([]int{1, 2, 3}, func(x int) int { return x * x })
	if !reflect.DeepEqual(got, []int{1, 4, 9}) {
		t.Errorf("=%v want [1 4 9]", got)
	}
	if got := MapInts(nil, func(x int) int { return x }); len(got) != 0 {
		t.Errorf("nil should map to empty")
	}
}
