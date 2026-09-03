package smallmapcache

import "testing"

var sinkI int
var sinkB bool

func TestPutGet(t *testing.T) {
	c := &Cache{Cap: 2}
	c.Put("a", 1)
	c.Put("b", 2)
	if v, ok := c.Get("a"); v != 1 || !ok {
		t.Errorf("Get(a) = %d, %v, want 1, true", v, ok)
	}
	if v, ok := c.Get("zz"); v != 0 || ok {
		t.Errorf("Get(zz) = %d, %v, want 0, false", v, ok)
	}
}

func TestEvictsOldestInsert(t *testing.T) {
	c := &Cache{Cap: 2}
	c.Put("a", 1)
	c.Put("b", 2)
	c.Get("a") // a hit must not save "a" from eviction
	c.Put("c", 3)
	if _, ok := c.Get("a"); ok {
		t.Error("a survived; the oldest insert must be evicted")
	}
	if _, ok := c.Get("b"); !ok {
		t.Error("b was evicted, want it kept")
	}
	if _, ok := c.Get("c"); !ok {
		t.Error("c is missing")
	}
	if c.Len() != 2 {
		t.Errorf("Len = %d, want 2", c.Len())
	}
}

func TestOverwriteKeepsOrderAndSize(t *testing.T) {
	c := &Cache{Cap: 2}
	c.Put("a", 1)
	c.Put("b", 2)
	c.Put("a", 99)
	if c.Len() != 2 {
		t.Errorf("Len = %d, want 2 — an overwrite is not an insert", c.Len())
	}
	c.Put("c", 3)
	if _, ok := c.Get("a"); ok {
		t.Error("a survived; overwriting must not refresh its position")
	}
	if v, ok := c.Get("c"); v != 3 || !ok {
		t.Errorf("Get(c) = %d, %v, want 3, true", v, ok)
	}
}

func TestZeroCapStoresNothing(t *testing.T) {
	c := &Cache{Cap: 0}
	c.Put("a", 1)
	if _, ok := c.Get("a"); ok {
		t.Error("a zero-capacity cache stored an entry")
	}
	if c.Len() != 0 {
		t.Errorf("Len = %d, want 0", c.Len())
	}
}

func TestStats(t *testing.T) {
	c := &Cache{Cap: 4}
	c.Put("a", 1)
	c.Get("a")
	c.Get("a")
	c.Get("b")
	hits, misses := c.Stats()
	if hits != 2 || misses != 1 {
		t.Errorf("Stats = %d, %d, want 2, 1", hits, misses)
	}
}

func TestGetHitDoesNotAllocate(t *testing.T) {
	c := &Cache{Cap: 64}
	for i := 0; i < 64; i++ {
		c.Put(string(rune('a'+i%26))+string(rune('a'+i/26)), i)
	}
	key := "aa"
	allocs := testing.AllocsPerRun(100, func() { sinkI, sinkB = c.Get(key) })
	if allocs != 0 {
		t.Errorf("Get on a hit made %v allocations, want 0", allocs)
	}
}

func TestStaysBounded(t *testing.T) {
	c := &Cache{Cap: 8}
	for i := 0; i < 10_000; i++ {
		c.Put(string(rune('a'+i%26))+string(rune(i)), i)
		if c.Len() > 8 {
			t.Fatalf("Len = %d, want at most 8", c.Len())
		}
	}
}
