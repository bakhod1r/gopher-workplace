package mapmerge

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	a := map[string]int{"x": 1, "y": 2}
	b := map[string]int{"y": 20, "z": 3}
	got := Merge(a, b)
	want := map[string]int{"x": 1, "y": 20, "z": 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Merge=%v; want %v", got, want)
	}
	// inputs unchanged
	if a["y"] != 2 {
		t.Error("input a mutated")
	}
}
