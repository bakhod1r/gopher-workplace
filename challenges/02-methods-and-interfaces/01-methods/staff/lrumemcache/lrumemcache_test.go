package lrumemcache

import "testing"

func TestLRU(t *testing.T) {
	l := New(2)
	l.Put("a", 1)
	l.Put("b", 2)

	if v, ok := l.Get("a"); !ok || v != 1 {
		t.Error("expected a=1")
	}

	l.Put("c", 3) // Evicts "b"

	if _, ok := l.Get("b"); ok {
		t.Error("expected b evicted")
	}
	if v, ok := l.Get("c"); !ok || v != 3 {
		t.Error("expected c=3")
	}
}
