package compactmap

import (
	"reflect"
	"testing"
)

func TestCompactKeepsEntries(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	got := Compact(m)
	if !reflect.DeepEqual(got, m) {
		t.Errorf("Compact = %v, want %v", got, m)
	}
}

func TestCompactReturnsANewMap(t *testing.T) {
	m := map[string]int{"a": 1}
	got := Compact(m)
	got["b"] = 2
	if _, ok := m["b"]; ok {
		t.Error("Compact returned the original map: the buckets are not released")
	}
}

func TestCompactAfterMassDeletion(t *testing.T) {
	m := make(map[string]int, 1<<12)
	for i := 0; i < 1<<12; i++ {
		m[string(rune('a'+i%26))+string(rune('a'+i/26))] = i
	}
	for k := range m {
		if m[k]%8 != 0 {
			delete(m, k)
		}
	}
	got := Compact(m)
	if len(got) != len(m) {
		t.Errorf("len = %d, want %d", len(got), len(m))
	}
	for k, v := range m {
		if got[k] != v {
			t.Fatalf("got[%q] = %d, want %d", k, got[k], v)
		}
	}
}

func TestCompactNil(t *testing.T) {
	if got := Compact(nil); len(got) != 0 {
		t.Errorf("Compact(nil) = %v, want empty", got)
	}
}
