package flatten

import (
	"reflect"
	"testing"
)

func TestFlatten(t *testing.T) {
	got := Flatten([][]int{{1, 2}, {3}, {}, {4, 5}})
	want := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Flatten=%v; want %v", got, want)
	}
	if len(Flatten(nil)) != 0 {
		t.Error("nil -> empty")
	}
}
