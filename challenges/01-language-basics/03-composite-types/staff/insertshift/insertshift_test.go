package insertshift

import (
	"reflect"
	"testing"
)

func TestInsertAt(t *testing.T) {
	got := InsertAt([]int{1, 2, 3, 4}, 1, 9)
	if !reflect.DeepEqual(got, []int{1, 9, 2, 3, 4}) {
		t.Errorf("InsertAt=%v; want [1 9 2 3 4]", got)
	}
	got = InsertAt([]int{1, 2}, 2, 9)
	if !reflect.DeepEqual(got, []int{1, 2, 9}) {
		t.Errorf("append end=%v; want [1 2 9]", got)
	}
}
