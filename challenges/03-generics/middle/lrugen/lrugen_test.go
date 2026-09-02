package lrugen

import "testing"

func TestLRUPromotesOnGet(t *testing.T) {
	c := NewLRU[string, int](2)
	c.Put("a", 1)
	c.Put("b", 2)
	if _, ok := c.Get("a"); !ok {
		t.Fatal(`Get("a") = false, want true`)
	}
	c.Put("c", 3)
	if _, ok := c.Get("b"); ok {
		t.Error(`Get("b") = true, want false (b was least recently used)`)
	}
	if v, ok := c.Get("a"); !ok || v != 1 {
		t.Errorf(`Get("a") = %v, %v, want 1, true`, v, ok)
	}
}

func TestLRUMiss(t *testing.T) {
	c := NewLRU[string, int](2)
	if v, ok := c.Get("nope"); v != 0 || ok {
		t.Errorf("Get(miss) = %v, %v, want 0, false", v, ok)
	}
}

func TestLRUZeroSize(t *testing.T) {
	c := NewLRU[string, int](0)
	c.Put("a", 1)
	if _, ok := c.Get("a"); ok {
		t.Error("a zero-size cache stored a value")
	}
}
