package flatten

import (
	"reflect"
	"testing"
)

func TestFlatten(t *testing.T) {
	got := Flatten([]int{1, 2}, []int{3}, []int{4, 5})
	if !reflect.DeepEqual(got, []int{1, 2, 3, 4, 5}) {
		t.Errorf("=%v want [1 2 3 4 5]", got)
	}
	if got := Flatten(); len(got) != 0 {
		t.Errorf("no groups should be empty")
	}
}
