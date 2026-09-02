package mapscopygen

import "testing"

func TestMerge(t *testing.T) {
	base := map[string]int{"a": 1, "b": 1}
	over := map[string]int{"b": 2, "c": 3}
	got := Merge(base, over)
	if got["a"] != 1 || got["b"] != 2 || got["c"] != 3 {
		t.Errorf("Merge = %v, want {a:1 b:2 c:3}", got)
	}
	if len(got) != 3 {
		t.Errorf("len = %d, want 3", len(got))
	}
}

func TestMergeDoesNotMutateInputs(t *testing.T) {
	base := map[string]int{"a": 1}
	over := map[string]int{"a": 2}
	Merge(base, over)
	if base["a"] != 1 {
		t.Errorf("base mutated: %v, want {a:1}", base)
	}
	if over["a"] != 2 || len(over) != 1 {
		t.Errorf("override mutated: %v, want {a:2}", over)
	}
}

func TestMergeEmpty(t *testing.T) {
	got := Merge(map[string]int(nil), map[string]int(nil))
	if got == nil {
		t.Fatal("Merge(nil, nil) = nil, want an empty non-nil map")
	}
	if len(got) != 0 {
		t.Errorf("Merge(nil, nil) = %v, want {}", got)
	}
}
