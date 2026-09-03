package profilediff

import (
	"reflect"
	"testing"
)

func TestDiff(t *testing.T) {
	got := Diff(map[string]int64{"a": 10}, map[string]int64{"a": 4})
	want := []Change{{"a", 10, 4, -6}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Diff = %v, want %v", got, want)
	}
}

func TestDiffCoversTheUnion(t *testing.T) {
	got := Diff(
		map[string]int64{"gone": 7, "same": 5},
		map[string]int64{"new": 9, "same": 5},
	)
	want := []Change{
		{"new", 0, 9, 9},
		{"gone", 7, 0, -7},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Diff = %v, want %v", got, want)
	}
}

func TestDiffOrdersByAbsoluteDelta(t *testing.T) {
	got := Diff(
		map[string]int64{"a": 100, "b": 10, "c": 50},
		map[string]int64{"a": 90, "b": 40, "c": 0},
	)
	want := []Change{
		{"c", 50, 0, -50},
		{"b", 10, 40, 30},
		{"a", 100, 90, -10},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Diff = %v, want %v", got, want)
	}
}

func TestDiffTieBrokenByName(t *testing.T) {
	got := Diff(
		map[string]int64{"z": 0, "a": 0, "m": 10},
		map[string]int64{"z": 5, "a": 5, "m": 5},
	)
	want := []Change{
		{"a", 0, 5, 5},
		{"m", 10, 5, -5},
		{"z", 0, 5, 5},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Diff = %v, want %v", got, want)
	}
}

func TestDiffDropsUnchanged(t *testing.T) {
	got := Diff(map[string]int64{"a": 5}, map[string]int64{"a": 5})
	if got == nil || len(got) != 0 {
		t.Errorf("Diff = %v, want empty non-nil slice", got)
	}
}
