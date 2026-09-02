package lrucache

import "testing"

func TestPutGet(t *testing.T) {
	l := NewLRU(2)
	l.Put("a", "1")
	if v, ok := l.Get("a"); v != "1" || !ok {
		t.Errorf("Get = %q, %v", v, ok)
	}
	if _, ok := l.Get("zz"); ok {
		t.Error("Get on a missing key returned ok")
	}
}

func TestEvictsLeastRecentlyUsed(t *testing.T) {
	l := NewLRU(2)
	l.Put("a", "1")
	l.Put("b", "2")
	l.Get("a") // a is now the most recent
	l.Put("c", "3")

	if _, ok := l.Get("b"); ok {
		t.Error("b was least recently used and should have been evicted")
	}
	if _, ok := l.Get("a"); !ok {
		t.Error("a was touched and should survive")
	}
	if _, ok := l.Get("c"); !ok {
		t.Error("c was just inserted")
	}
	if l.Len() != 2 {
		t.Errorf("Len = %d, want 2", l.Len())
	}
}

func TestPutExistingTouches(t *testing.T) {
	l := NewLRU(2)
	l.Put("a", "1")
	l.Put("b", "2")
	l.Put("a", "9") // update and touch
	l.Put("c", "3")

	if _, ok := l.Get("b"); ok {
		t.Error("b should have been evicted")
	}
	if v, ok := l.Get("a"); v != "9" || !ok {
		t.Errorf("a = %q, %v; want \"9\", true", v, ok)
	}
	if l.Len() != 2 {
		t.Errorf("Len = %d, want 2", l.Len())
	}
}

func TestCeilingUnderChurn(t *testing.T) {
	l := NewLRU(8)
	for i := 0; i < 10000; i++ {
		k := string(rune('a' + i%26))
		l.Put(k, k)
		if l.Len() > 8 {
			t.Fatalf("Len = %d, exceeded Cap 8", l.Len())
		}
	}
}

func TestZeroCapacity(t *testing.T) {
	l := NewLRU(0)
	l.Put("a", "1")
	if l.Len() != 0 {
		t.Errorf("Len = %d, want 0", l.Len())
	}
	if _, ok := l.Get("a"); ok {
		t.Error("nothing should be cached")
	}
}

func BenchmarkPutGet(b *testing.B) {
	l := NewLRU(128)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		k := string(rune('a' + i%26))
		l.Put(k, k)
		l.Get(k)
	}
}
