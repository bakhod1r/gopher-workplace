package partition

import (
	"reflect"
	"testing"
)

func TestPartition(t *testing.T) {
	e, o := Partition([]int{1, 2, 3, 4, 5, 6})
	if !reflect.DeepEqual(e, []int{2, 4, 6}) || !reflect.DeepEqual(o, []int{1, 3, 5}) {
		t.Errorf("Partition=%v,%v; want [2 4 6],[1 3 5]", e, o)
	}
	e, o = Partition(nil)
	if len(e) != 0 || len(o) != 0 {
		t.Error("nil -> empty,empty")
	}
}
