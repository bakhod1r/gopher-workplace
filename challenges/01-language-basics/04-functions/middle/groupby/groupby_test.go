package groupby

import (
	"reflect"
	"testing"
)

func TestGroupBy(t *testing.T) {
	mod := func(x int) int { return x % 2 }
	got := GroupBy([]int{1, 2, 3, 4, 5}, mod)
	want := map[int][]int{0: {2, 4}, 1: {1, 3, 5}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("=%v want %v", got, want)
	}
}
