package nilappend

import (
	"reflect"
	"testing"
)

func TestCollect(t *testing.T) {
	got := Collect(nil, []int{1, 2, 3})
	if !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("=%v want [1 2 3] (append to nil is fine)", got)
	}
	got = Collect([]int{0}, []int{9})
	if !reflect.DeepEqual(got, []int{0, 9}) {
		t.Errorf("=%v want [0 9]", got)
	}
}
