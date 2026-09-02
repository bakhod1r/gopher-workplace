package mapsequalgen

import "testing"

func TestSameConfig(t *testing.T) {
	if !SameConfig(map[string]int{"a": 1}, map[string]int{"a": 1}) {
		t.Error("SameConfig({a:1}, {a:1}) = false, want true")
	}
	if SameConfig(map[string]int{"a": 1}, map[string]int{"a": 2}) {
		t.Error("SameConfig({a:1}, {a:2}) = true, want false")
	}
	if SameConfig(map[string]int{"a": 1}, map[string]int{"a": 1, "b": 2}) {
		t.Error("different sizes compared equal, want false")
	}
	if !SameConfig(map[string]int(nil), map[string]int{}) {
		t.Error("SameConfig(nil, {}) = false, want true")
	}
	if !SameConfig(map[int]string{1: "a"}, map[int]string{1: "a"}) {
		t.Error("SameConfig with int keys = false, want true")
	}
}
