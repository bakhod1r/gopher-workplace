package setdiff

import (
	"reflect"
	"testing"
)

func TestDiff(t *testing.T) {
	got := Diff([]int{1, 2, 3, 3, 4}, []int{2, 4})
	want := []int{1, 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Diff=%v; want %v", got, want)
	}
	if len(Diff([]int{1}, []int{1})) != 0 {
		t.Error("all removed -> empty")
	}
}
