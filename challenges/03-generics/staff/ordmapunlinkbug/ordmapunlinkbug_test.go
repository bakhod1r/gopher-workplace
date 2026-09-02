package ordmapunlinkbug

import (
	"reflect"
	"testing"
	"time"
)

func TestDeleteMiddle(t *testing.T) {
	var m OrderedMap[string, int]
	m.Set("a", 1)
	m.Set("b", 2)
	m.Set("c", 3)
	if !m.Delete("b") {
		t.Fatal("Delete(b) = false, want true")
	}
	if got := m.Keys(); !reflect.DeepEqual(got, []string{"a", "c"}) {
		t.Errorf("Keys = %v, want [a c]", got)
	}
	if got := m.RevKeys(); !reflect.DeepEqual(got, []string{"c", "a"}) {
		t.Errorf("RevKeys = %v, want [c a]", got)
	}
}

func TestDeleteTailThenSet(t *testing.T) {
	var m OrderedMap[string, int]
	m.Set("a", 1)
	m.Set("b", 2)
	if !m.Delete("b") {
		t.Fatal("Delete(b) = false, want true")
	}
	m.Set("c", 3)
	if got := m.Keys(); !reflect.DeepEqual(got, []string{"a", "c"}) {
		t.Errorf("Keys = %v, want [a c]", got)
	}
	if m.Len() != len(m.Keys()) {
		t.Errorf("Len = %d but Keys has %d entries", m.Len(), len(m.Keys()))
	}
	if _, ok := m.Get("c"); !ok {
		t.Error("Get(c) missing after reinsert")
	}
}

func TestDeleteMissing(t *testing.T) {
	var m OrderedMap[string, int]
	m.Set("a", 1)
	if m.Delete("zz") {
		t.Error("Delete(zz) = true, want false")
	}
	if m.Len() != 1 {
		t.Errorf("Len = %d, want 1", m.Len())
	}
}

func TestDeleteScale(t *testing.T) {
	const n = 5000
	start := time.Now()
	var m OrderedMap[int, int]
	for i := 0; i < n; i++ {
		m.Set(i, i*i)
	}
	for i := 0; i < n; i += 2 {
		m.Delete(i)
	}
	ks := m.Keys()
	if len(ks) != n/2 {
		t.Fatalf("Keys has %d entries, want %d", len(ks), n/2)
	}
	if m.Len() != len(ks) {
		t.Fatalf("Len = %d but Keys has %d entries", m.Len(), len(ks))
	}
	for j, k := range ks {
		if k != 2*j+1 {
			t.Fatalf("Keys[%d] = %d, want %d", j, k, 2*j+1)
		}
	}
	rs := m.RevKeys()
	if len(rs) != len(ks) {
		t.Fatalf("RevKeys has %d entries, want %d", len(rs), len(ks))
	}
	for j, k := range rs {
		if k != ks[len(ks)-1-j] {
			t.Fatalf("RevKeys[%d] = %d, want %d", j, k, ks[len(ks)-1-j])
		}
	}
	if d := time.Since(start); d > 2*time.Second {
		t.Fatalf("5k insert/delete took %v, want under 2s", d)
	}
}
