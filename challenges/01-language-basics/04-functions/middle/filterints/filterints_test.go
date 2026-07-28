package filterints

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	even := func(x int) bool { return x%2 == 0 }
	got := Filter([]int{1, 2, 3, 4, 5, 6}, even)
	if !reflect.DeepEqual(got, []int{2, 4, 6}) {
		t.Errorf("=%v want [2 4 6]", got)
	}
}
