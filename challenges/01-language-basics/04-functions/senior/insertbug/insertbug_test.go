package insertbug

import (
	"reflect"
	"testing"
)

func TestInsertAt(t *testing.T) {
	got := InsertAt([]int{1, 2, 4}, 2, 3)
	if !reflect.DeepEqual(got, []int{1, 2, 3, 4}) {
		t.Errorf("=%v want [1 2 3 4]", got)
	}
	got = InsertAt([]int{2, 3}, 0, 1)
	if !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("=%v want [1 2 3]", got)
	}
}
