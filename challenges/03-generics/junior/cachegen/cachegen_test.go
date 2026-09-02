package cachegen

import "testing"

func TestCacheEviction(t *testing.T) {
	c := NewCache[string, int](2)
	c.Put("a", 1)
	c.Put("b", 2)
	c.Put("c", 3)
	if _, ok := c.Get("a"); ok {
		t.Error(`Get("a") reported ok, want false (oldest should be evicted)`)
	}
	if v, ok := c.Get("b"); v != 2 || !ok {
		t.Errorf(`Get("b") = %v, %v, want 2, true`, v, ok)
	}
	if v, ok := c.Get("c"); v != 3 || !ok {
		t.Errorf(`Get("c") = %v, %v, want 3, true`, v, ok)
	}
}

func TestCacheUpdateDoesNotEvict(t *testing.T) {
	c := NewCache[string, int](2)
	c.Put("a", 1)
	c.Put("a", 2)
	c.Put("b", 3)
	if v, ok := c.Get("a"); v != 2 || !ok {
		t.Errorf(`Get("a") = %v, %v, want 2, true`, v, ok)
	}
	if v, ok := c.Get("b"); v != 3 || !ok {
		t.Errorf(`Get("b") = %v, %v, want 3, true`, v, ok)
	}
}

func TestCacheZeroSize(t *testing.T) {
	c := NewCache[string, int](0)
	c.Put("a", 1)
	if _, ok := c.Get("a"); ok {
		t.Error("a zero-size cache stored a value, want nothing stored")
	}
}
