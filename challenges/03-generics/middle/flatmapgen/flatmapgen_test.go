package flatmapgen

import (
	"reflect"
	"testing"
)

func TestFlatMap(t *testing.T) {
	dup := func(n int) []int { return []int{n, n} }
	if got := FlatMap([]int{1, 2}, dup); !reflect.DeepEqual(got, []int{1, 1, 2, 2}) {
		t.Errorf("FlatMap = %v, want [1 1 2 2]", got)
	}
	none := func(n int) []int { return nil }
	if got := FlatMap([]int{1, 2}, none); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("FlatMap = %v, want []", got)
	}
	if got := FlatMap([]int{}, dup); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("FlatMap([]) = %v, want []", got)
	}
}

func TestFlatMapChangesType(t *testing.T) {
	split := func(n int) []string { return []string{"a", "b"} }
	got := FlatMap([]int{1}, split)
	if !reflect.DeepEqual(got, []string{"a", "b"}) {
		t.Errorf("FlatMap = %v, want [a b]", got)
	}
}
