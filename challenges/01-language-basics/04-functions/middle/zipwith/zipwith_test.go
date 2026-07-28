package zipwith

import (
	"reflect"
	"testing"
)

func TestZipWith(t *testing.T) {
	add := func(x, y int) int { return x + y }
	got := ZipWith([]int{1, 2, 3}, []int{10, 20}, add)
	if !reflect.DeepEqual(got, []int{11, 22}) {
		t.Errorf("=%v want [11 22]", got)
	}
}
