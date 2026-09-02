package cachelayer

import "testing"

func TestSlowSource(t *testing.T) {
	s := &SlowSource{Data: map[string]string{"a": "1"}}
	if v, ok := s.Get("a"); v != "1" || !ok {
		t.Errorf("Get = %q, %v", v, ok)
	}
	if v, ok := s.Get("zz"); v != "" || ok {
		t.Errorf("Get = %q, %v, want \"\", false", v, ok)
	}
	if s.Calls != 2 {
		t.Errorf("Calls = %d, want 2", s.Calls)
	}
}

func TestCacheHit(t *testing.T) {
	s := &SlowSource{Data: map[string]string{"a": "1"}}
	c := NewCache(s)

	for i := 0; i < 3; i++ {
		if v, ok := c.Get("a"); v != "1" || !ok {
			t.Fatalf("Get %d = %q, %v", i, v, ok)
		}
	}
	if s.Calls != 1 {
		t.Errorf("source Calls = %d, want 1", s.Calls)
	}
}

func TestCacheMiss(t *testing.T) {
	s := &SlowSource{Data: map[string]string{}}
	c := NewCache(s)

	if v, ok := c.Get("nope"); v != "" || ok {
		t.Errorf("Get = %q, %v, want \"\", false", v, ok)
	}
	c.Get("nope")
	if s.Calls != 1 {
		t.Errorf("source Calls = %d, want 1 (misses must be cached too)", s.Calls)
	}
}

func TestCacheIsSource(t *testing.T) {
	var src Source = NewCache(&SlowSource{Data: map[string]string{"k": "v"}})
	if v, ok := src.Get("k"); v != "v" || !ok {
		t.Errorf("Get = %q, %v", v, ok)
	}
}

func TestEmptyValueCached(t *testing.T) {
	s := &SlowSource{Data: map[string]string{"e": ""}}
	c := NewCache(s)
	if v, ok := c.Get("e"); v != "" || !ok {
		t.Errorf("Get = %q, %v; want \"\", true", v, ok)
	}
	c.Get("e")
	if s.Calls != 1 {
		t.Errorf("source Calls = %d, want 1", s.Calls)
	}
}
