package mergedirection

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	base := map[string]int{"a": 1, "b": 2}
	over := map[string]int{"b": 20, "c": 3}
	got := Merge(base, over)
	want := map[string]int{"a": 1, "b": 20, "c": 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Merge=%v; want %v", got, want)
	}
}
