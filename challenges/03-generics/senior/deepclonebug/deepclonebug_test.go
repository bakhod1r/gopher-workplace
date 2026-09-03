package deepclonebug

import "testing"

func TestSnapshotContent(t *testing.T) {
	m := map[string][]int{"a": {1, 2}, "b": {3}}
	got := Snapshot(m)
	if len(got) != 2 || len(got["a"]) != 2 || got["b"][0] != 3 {
		t.Fatalf("Snapshot = %v, want the same content", got)
	}
}

func TestSnapshotIsIndependent(t *testing.T) {
	m := map[string][]int{"a": {1, 2}}
	got := Snapshot(m)
	got["a"][0] = 99
	if m["a"][0] != 1 {
		t.Errorf("writing into the snapshot changed the source: %v", m)
	}
	m["a"][1] = 42
	if got["a"][1] != 2 {
		t.Errorf("writing into the source changed the snapshot: %v", got)
	}
}

func TestSnapshotNil(t *testing.T) {
	got := Snapshot(map[string][]int(nil))
	if got == nil {
		t.Fatal("Snapshot(nil) = nil, want an empty non-nil map")
	}
	if len(got) != 0 {
		t.Errorf("Snapshot(nil) = %v, want {}", got)
	}
}
