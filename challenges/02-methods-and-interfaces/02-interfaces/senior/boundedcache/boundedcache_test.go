package boundedcache

import "testing"

func src(pairs map[string]string) *CountingSource {
	return &CountingSource{Data: pairs}
}

func TestCacheHit(t *testing.T) {
	s := src(map[string]string{"a": "1"})
	c := NewCache(s, 2)

	for i := 0; i < 5; i++ {
		if got := c.Get("a"); got != "1" {
			t.Fatalf("Get = %q, want \"1\"", got)
		}
	}
	if s.Calls != 1 {
		t.Errorf("source Calls = %d, want 1", s.Calls)
	}
}

func TestEvictionOrder(t *testing.T) {
	s := src(map[string]string{"a": "1", "b": "2", "c": "3"})
	c := NewCache(s, 2)

	c.Get("a")
	c.Get("b")
	c.Get("c")

	if c.Len() != 2 {
		t.Fatalf("Len = %d, want 2", c.Len())
	}
	if _, ok := c.entries["a"]; ok {
		t.Error("oldest key \"a\" should have been evicted")
	}
	if _, ok := c.entries["c"]; !ok {
		t.Error("newest key \"c\" should be cached")
	}

	before := s.Calls
	c.Get("a")
	if s.Calls != before+1 {
		t.Error("evicted key should be refetched")
	}
}

func TestCeilingHolds(t *testing.T) {
	data := make(map[string]string, 1000)
	keys := make([]string, 0, 1000)
	for i := 0; i < 1000; i++ {
		k := string(rune('a'+i%26)) + string(rune('a'+(i/26)%26)) + string(rune('a'+i/676))
		data[k] = k
		keys = append(keys, k)
	}

	s := src(data)
	c := NewCache(s, 10)
	for _, k := range keys {
		c.Get(k)
		if c.Len() > 10 {
			t.Fatalf("Len = %d, exceeded Max 10", c.Len())
		}
	}
	if c.Len() != 10 {
		t.Errorf("final Len = %d, want 10", c.Len())
	}
}

func TestZeroMaxCachesNothing(t *testing.T) {
	s := src(map[string]string{"a": "1"})
	c := NewCache(s, 0)
	c.Get("a")
	c.Get("a")
	if c.Len() != 0 {
		t.Errorf("Len = %d, want 0", c.Len())
	}
	if s.Calls != 2 {
		t.Errorf("Calls = %d, want 2", s.Calls)
	}
}

func BenchmarkGet(b *testing.B) {
	s := src(map[string]string{"a": "1"})
	c := NewCache(s, 16)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		c.Get("a")
	}
}
