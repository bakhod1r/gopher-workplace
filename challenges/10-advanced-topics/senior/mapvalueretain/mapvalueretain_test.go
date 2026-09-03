package mapvalueretain

import (
	"bytes"
	"testing"
)

func TestIndexStoresTheBytes(t *testing.T) {
	m := map[string][]byte{}
	batch := []byte("hello world")
	Index(m, "a", batch, 0, 5)
	if !bytes.Equal(m["a"], []byte("hello")) {
		t.Errorf("m[a] = %q, want \"hello\"", m["a"])
	}
}

func TestIndexSurvivesBatchReuse(t *testing.T) {
	m := map[string][]byte{}
	batch := make([]byte, 16)
	copy(batch, "first-value")
	Index(m, "a", batch, 0, 5)
	copy(batch, "SECOND-VALUE")
	if !bytes.Equal(m["a"], []byte("first")) {
		t.Errorf("m[a] = %q, want \"first\": the entry views the reused batch", m["a"])
	}
}

func TestIndexReleasesTheBatch(t *testing.T) {
	m := map[string][]byte{}
	batch := make([]byte, 1<<20)
	Index(m, "a", batch, 0, 8)
	if cap(m["a"]) > 64 {
		t.Errorf("cap = %d, want a right-sized copy: the entry still owns the batch's array", cap(m["a"]))
	}
}

func TestIndexBadRanges(t *testing.T) {
	m := map[string][]byte{}
	batch := []byte("abcd")
	for _, c := range [][2]int{{-1, 2}, {0, -1}, {3, 3}, {5, 1}} {
		Index(m, "k", batch, c[0], c[1])
		if _, ok := m["k"]; ok {
			t.Fatalf("off=%d n=%d stored an entry, want none", c[0], c[1])
		}
	}
}

func TestIndexNilMap(t *testing.T) {
	Index(nil, "a", []byte("x"), 0, 1)
}
