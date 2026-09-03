package lrucache

import (
	"sync"
	"testing"
)

func TestGetAndPut(t *testing.T) {
	c := NewLRU(2)
	c.Put("a", 1)
	if v, ok := c.Get("a"); !ok || v != 1 {
		t.Errorf("Get(a) = %d, %v, want 1, true", v, ok)
	}
	if _, ok := c.Get("missing"); ok {
		t.Error("Get(missing) reported ok, want false")
	}
}

func TestGetMarksAsRecentlyUsed(t *testing.T) {
	c := NewLRU(2)
	c.Put("a", 1)
	c.Put("b", 2)
	if _, ok := c.Get("a"); !ok {
		t.Fatal("a was not found")
	}
	c.Put("c", 3) // evicts the least recently used, which must be b
	if _, ok := c.Get("a"); !ok {
		t.Error("a was evicted, but Get(a) had just made it the newest")
	}
	if _, ok := c.Get("b"); ok {
		t.Error("b survived, want it evicted as the least recently used")
	}
}

func TestEvictsWithoutAnyGets(t *testing.T) {
	c := NewLRU(2)
	c.Put("a", 1)
	c.Put("b", 2)
	c.Put("c", 3)
	if _, ok := c.Get("a"); ok {
		t.Error("a survived, want it evicted")
	}
	if c.Len() != 2 {
		t.Errorf("Len = %d, want 2", c.Len())
	}
}

func TestStaysBounded(t *testing.T) {
	c := NewLRU(8)
	for i := 0; i < 2000; i++ {
		c.Put(string(rune('a'+i%26))+string(rune('a'+i/26%26)), i)
		c.Get(string(rune('a' + i%26)))
		if c.Len() > 8 {
			t.Fatalf("Len = %d, want at most 8", c.Len())
		}
	}
}

func TestConcurrentAccess(t *testing.T) {
	c := NewLRU(16)
	var wg sync.WaitGroup
	const workers = 8
	wg.Add(workers)
	for w := 0; w < workers; w++ {
		go func(w int) {
			defer wg.Done()
			for i := 0; i < 500; i++ {
				k := string(rune('a' + i%10))
				c.Put(k, i)
				c.Get(k)
			}
		}(w)
	}
	wg.Wait()
	if c.Len() > 16 {
		t.Errorf("Len = %d, want at most 16", c.Len())
	}
}
