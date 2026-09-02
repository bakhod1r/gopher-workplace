package lruo1bug

import (
	"reflect"
	"testing"
	"time"
)

func TestLRUBasics(t *testing.T) {
	c := NewLRU[int, string](2)
	c.Put(1, "a")
	c.Put(2, "b")
	if v, ok := c.Get(1); !ok || v != "a" {
		t.Errorf("Get(1) = %q, %v, want a, true", v, ok)
	}
	if _, ok := c.Get(9); ok {
		t.Errorf("Get(9) reported a hit")
	}
	if c.Len() != 2 {
		t.Errorf("Len = %d, want 2", c.Len())
	}
}

func TestLRUEvictsLeastRecentlyUsed(t *testing.T) {
	c := NewLRU[int, string](2)
	c.Put(1, "a")
	c.Put(2, "b")
	c.Get(1)
	c.Put(3, "c")
	if got := c.Keys(); !reflect.DeepEqual(got, []int{3, 1}) {
		t.Errorf("Keys = %v, want [3 1]", got)
	}
	if _, ok := c.Get(2); ok {
		t.Errorf("key 2 survived eviction")
	}
}

func TestLRUGetIsConstantTime(t *testing.T) {
	const capacity = 16384
	const lookups = 150000
	const budget = 2 * time.Second

	c := NewLRU[int, int](capacity)
	for i := 0; i < capacity; i++ {
		c.Put(i, i)
	}

	start := time.Now()
	sum := 0
	for i := 0; i < lookups; i++ {
		if v, ok := c.Get((i * 7919) % capacity); ok {
			sum += v
		}
	}
	elapsed := time.Since(start)

	if sum == 0 {
		t.Fatal("no lookups hit the cache")
	}
	if c.Len() != capacity {
		t.Fatalf("Len = %d, want %d", c.Len(), capacity)
	}
	if elapsed > budget {
		t.Errorf("%d lookups on a %d-entry cache took %v, budget %v: Get is not constant time",
			lookups, capacity, elapsed, budget)
	}
}
