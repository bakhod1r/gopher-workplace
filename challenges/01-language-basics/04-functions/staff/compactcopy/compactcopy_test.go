package compactcopy

import (
	"reflect"
	"testing"
)

func TestDropFirst(t *testing.T) {
	got := DropFirst([]int{10, 20, 30, 40})
	if !reflect.DeepEqual(got, []int{20, 30, 40}) {
		t.Errorf("=%v want [20 30 40]", got)
	}
}
