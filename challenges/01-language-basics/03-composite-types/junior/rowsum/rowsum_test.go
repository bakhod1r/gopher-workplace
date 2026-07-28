package rowsum

import (
	"reflect"
	"testing"
)

func TestRowSums(t *testing.T) {
	grid := [][]int{{1, 2, 3}, {4, 5}, {}}
	got := RowSums(grid)
	want := []int{6, 9, 0}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("RowSums=%v; want %v", got, want)
	}
	if len(RowSums(nil)) != 0 {
		t.Error("nil grid -> empty")
	}
}
