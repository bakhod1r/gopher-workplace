package transpose

import (
	"reflect"
	"testing"
)

func TestTranspose(t *testing.T) {
	got := Transpose([][]int{{1, 2, 3}, {4, 5, 6}})
	want := [][]int{{1, 4}, {2, 5}, {3, 6}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Transpose=%v; want %v", got, want)
	}
	if len(Transpose(nil)) != 0 {
		t.Error("nil -> empty")
	}
}
