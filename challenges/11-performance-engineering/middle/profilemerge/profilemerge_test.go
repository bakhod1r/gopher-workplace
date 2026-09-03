package profilemerge

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	got := Merge([]Profile{
		{Flat: map[string]int64{"a": 1, "b": 5}, Samples: 3},
		{Flat: map[string]int64{"a": 2, "c": 7}, Samples: 4},
	})
	want := Profile{Flat: map[string]int64{"a": 3, "b": 5, "c": 7}, Samples: 7}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Merge = %v, want %v", got, want)
	}
}

func TestMergePreservesTheTotal(t *testing.T) {
	in := []Profile{
		{Flat: map[string]int64{"a": 1, "b": 5}, Samples: 3},
		{Flat: map[string]int64{"a": 2, "c": 7}, Samples: 4},
	}
	var want int64
	for _, p := range in {
		want += Total(p)
	}
	if got := Total(Merge(in)); got != want {
		t.Errorf("merged total = %d, want %d — merging must not lose or duplicate samples", got, want)
	}
}

func TestMergeDoesNotModifyInputs(t *testing.T) {
	a := Profile{Flat: map[string]int64{"a": 1}, Samples: 1}
	b := Profile{Flat: map[string]int64{"a": 2}, Samples: 1}
	Merge([]Profile{a, b})
	if a.Flat["a"] != 1 || b.Flat["a"] != 2 {
		t.Errorf("an input was modified: a = %v, b = %v", a.Flat, b.Flat)
	}
}

func TestMergeSkipsEmptyProfiles(t *testing.T) {
	got := Merge([]Profile{
		{Flat: nil, Samples: 2},
		{Flat: map[string]int64{"a": 1}, Samples: 1},
	})
	want := Profile{Flat: map[string]int64{"a": 1}, Samples: 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Merge = %v, want %v", got, want)
	}
}

func TestMergeNothing(t *testing.T) {
	got := Merge(nil)
	if got.Flat == nil || len(got.Flat) != 0 || got.Samples != 0 {
		t.Errorf("Merge(nil) = %v, want an empty non-nil map and 0 samples", got)
	}
}

func TestTotal(t *testing.T) {
	if got := Total(Profile{Flat: map[string]int64{"a": 1, "b": 2}}); got != 3 {
		t.Errorf("Total = %d, want 3", got)
	}
	if got := Total(Profile{}); got != 0 {
		t.Errorf("Total(empty) = %d, want 0", got)
	}
}
