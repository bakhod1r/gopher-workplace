package mapsclonegen

import "testing"

func TestSnapshot(t *testing.T) {
	live := map[string]int{"a": 1}
	snap := Snapshot(live)
	if len(snap) != 1 || snap["a"] != 1 {
		t.Fatalf("Snapshot = %v, want {a:1}", snap)
	}
	snap["a"] = 99
	snap["b"] = 2
	if live["a"] != 1 {
		t.Errorf(`live["a"] = %v, want 1 (the copy must be independent)`, live["a"])
	}
	if len(live) != 1 {
		t.Errorf("live gained keys: %v", live)
	}
}

func TestSnapshotNil(t *testing.T) {
	got := Snapshot(map[string]int(nil))
	if got == nil {
		t.Fatal("Snapshot(nil) = nil, want an empty non-nil map")
	}
	got["a"] = 1
	if len(got) != 1 {
		t.Errorf("writing to the result failed: %v", got)
	}
}
