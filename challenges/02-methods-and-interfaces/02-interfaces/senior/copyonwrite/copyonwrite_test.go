package copyonwrite

import (
	"strconv"
	"sync"
	"testing"
)

func TestStoreLoad(t *testing.T) {
	c := NewConfig()
	c.Store(Snapshot{"a": "1"})
	if got := c.Load()["a"]; got != "1" {
		t.Errorf("Load = %q, want \"1\"", got)
	}
}

func TestUpdateAddsKey(t *testing.T) {
	c := NewConfig()
	c.Update(MutatorFunc(func(s Snapshot) { s["a"] = "1" }))
	if got := c.Load()["a"]; got != "1" {
		t.Errorf("Load = %q, want \"1\"", got)
	}
}

func TestOldSnapshotUnchanged(t *testing.T) {
	c := NewConfig()
	c.Update(MutatorFunc(func(s Snapshot) { s["a"] = "1" }))

	old := c.Load()
	c.Update(MutatorFunc(func(s Snapshot) { s["b"] = "2" }))

	if _, ok := old["b"]; ok {
		t.Error("the previously loaded snapshot was mutated in place")
	}
	if len(old) != 1 {
		t.Errorf("old snapshot has %d keys, want 1", len(old))
	}
	if len(c.Load()) != 2 {
		t.Errorf("new snapshot has %d keys, want 2", len(c.Load()))
	}
}

func TestUpdateOverwrites(t *testing.T) {
	c := NewConfig()
	c.Update(MutatorFunc(func(s Snapshot) { s["a"] = "1" }))
	c.Update(MutatorFunc(func(s Snapshot) { s["a"] = "2" }))
	if got := c.Load()["a"]; got != "2" {
		t.Errorf("Load = %q, want \"2\"", got)
	}
}

func TestConcurrentReadersAndWriters(t *testing.T) {
	c := NewConfig()
	var wg sync.WaitGroup

	for w := 0; w < 8; w++ {
		wg.Add(1)
		go func(w int) {
			defer wg.Done()
			for i := 0; i < 100; i++ {
				k := strconv.Itoa(w*100 + i)
				c.Update(MutatorFunc(func(s Snapshot) { s[k] = k }))
			}
		}(w)
	}

	for r := 0; r < 8; r++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				s := c.Load()
				for k, v := range s {
					if k != v {
						t.Errorf("inconsistent snapshot: %q => %q", k, v)
						return
					}
				}
			}
		}()
	}

	wg.Wait()
	if got := len(c.Load()); got != 800 {
		t.Errorf("final snapshot has %d keys, want 800", got)
	}
}

func BenchmarkLoad(b *testing.B) {
	c := NewConfig()
	c.Update(MutatorFunc(func(s Snapshot) { s["a"] = "1" }))
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = c.Load()["a"]
		}
	})
}
