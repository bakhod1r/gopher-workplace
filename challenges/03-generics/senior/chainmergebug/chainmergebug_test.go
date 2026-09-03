package chainmergebug

import (
	"reflect"
	"testing"
)

func TestMergeLaterWins(t *testing.T) {
	got := Merge(map[string]int{"a": 1}, map[string]int{"a": 2})
	if !reflect.DeepEqual(got, map[string]int{"a": 2}) {
		t.Errorf("Merge = %v, want map[a:2]", got)
	}
}

func TestMergeUnionOfKeys(t *testing.T) {
	got := Merge(map[string]int{"a": 1}, map[string]int{"b": 2})
	if !reflect.DeepEqual(got, map[string]int{"a": 1, "b": 2}) {
		t.Errorf("Merge = %v, want map[a:1 b:2]", got)
	}
}

func TestMergeNoArgs(t *testing.T) {
	if got := Merge[string, int](); len(got) != 0 {
		t.Errorf("Merge = %v, want empty", got)
	}
}
