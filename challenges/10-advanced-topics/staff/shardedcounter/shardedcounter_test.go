package shardedcounter

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

func TestAddAndTotal(t *testing.T) {
	c := NewCounter(4)
	c.Add("a", 1)
	c.Add("a", 2)
	c.Add("b", 5)
	got := c.Total()
	if got["a"] != 3 || got["b"] != 5 {
		t.Errorf("Total = %v, want map[a:3 b:5]", got)
	}
}

func TestAddNegativeAndZero(t *testing.T) {
	c := NewCounter(2)
	c.Add("k", 5)
	c.Add("k", -5)
	if got := c.Total(); got["k"] != 0 {
		t.Errorf("Total = %v, want map[k:0]", got)
	}
}

func TestSingleShard(t *testing.T) {
	c := NewCounter(0)
	c.Add("a", 1)
	c.Add("b", 2)
	got := c.Total()
	if got["a"] != 1 || got["b"] != 2 {
		t.Errorf("Total = %v, want map[a:1 b:2]", got)
	}
}

func TestConcurrentAdds(t *testing.T) {
	const (
		workers = 16
		perTask = 1000
		keys    = 8
	)
	c := NewCounter(16)
	var wg sync.WaitGroup
	wg.Add(workers)
	for w := 0; w < workers; w++ {
		go func() {
			defer wg.Done()
			for i := 0; i < perTask; i++ {
				c.Add(fmt.Sprintf("k%d", i%keys), 1)
			}
		}()
	}
	wg.Wait()
	got := c.Total()
	if len(got) != keys {
		t.Fatalf("Total has %d keys, want %d", len(got), keys)
	}
	want := int64(workers * perTask / keys)
	for k, v := range got {
		if v != want {
			t.Fatalf("Total[%q] = %d, want %d: increments were lost", k, v, want)
		}
	}
}

func TestSameKeyAlwaysHitsOneShard(t *testing.T) {
	c := NewCounter(8)
	first := c.shardFor("stable")
	for i := 0; i < 100; i++ {
		if c.shardFor("stable") != first {
			t.Fatal("shardFor is not deterministic for one counter")
		}
	}
}

func TestShardsDoNotShareALine(t *testing.T) {
	if got := unsafe.Sizeof(shard{}); got < 64 {
		t.Errorf("sizeof(shard) = %d, want at least 64: neighbouring shards share a cache line", got)
	}
}
