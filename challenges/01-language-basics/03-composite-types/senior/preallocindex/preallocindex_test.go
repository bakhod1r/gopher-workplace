package preallocindex

import (
	"reflect"
	"testing"
)

func TestDoubled(t *testing.T) {
	got := Doubled([]int{1, 2, 3})
	if !reflect.DeepEqual(got, []int{2, 4, 6}) {
		t.Errorf("Doubled=%v; want [2 4 6]", got)
	}
	if len(Doubled(nil)) != 0 {
		t.Error("nil -> empty")
	}
}
