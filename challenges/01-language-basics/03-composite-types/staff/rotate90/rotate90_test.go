package rotate90

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	m := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	want := [][]int{
		{7, 4, 1},
		{8, 5, 2},
		{9, 6, 3},
	}
	if got := Rotate(m); !reflect.DeepEqual(got, want) {
		t.Errorf("Rotate=%v; want %v", got, want)
	}
}
