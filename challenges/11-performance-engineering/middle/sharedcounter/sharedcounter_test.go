package sharedcounter

import (
	"reflect"
	"sync"
	"testing"
)

func TestAddAndGet(t *testing.T) {
	var c Counter
	c.Add("a", 1)
	c.Add("a", 2)
	c.Add("b", -5)
	if got := c.Get("a"); got != 3 {
		t.Errorf("Get(a) = %d, want 3", got)
	}
	if got := c.Get("b"); got != -5 {
		t.Errorf("Get(b) = %d, want -5", got)
	}
	if got := c.Get("missing"); got != 0 {
		t.Errorf("Get(missing) = %d, want 0", got)
	}
}

func TestSnapshotIsACopy(t *testing.T) {
	var c Counter
	c.Add("a", 1)
	s := c.Snapshot()
	s["a"] = 999
	s["new"] = 1
	if got := c.Get("a"); got != 1 {
		t.Errorf("Get(a) = %d, want 1 — Snapshot must copy", got)
	}
	if got := c.Get("new"); got != 0 {
		t.Errorf("Get(new) = %d, want 0", got)
	}
	if want := (map[string]int64{"a": 1}); !reflect.DeepEqual(c.Snapshot(), want) {
		t.Errorf("Snapshot = %v, want %v", c.Snapshot(), want)
	}
}

func TestSnapshotOfEmptyCounter(t *testing.T) {
	var c Counter
	got := c.Snapshot()
	if got == nil || len(got) != 0 {
		t.Errorf("Snapshot = %v, want empty non-nil map", got)
	}
}

func TestConcurrentAddsLoseNothing(t *testing.T) {
	var c Counter
	const goroutines, each = 50, 1000
	var wg sync.WaitGroup
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < each; j++ {
				c.Add("shared", 1)
			}
		}()
	}
	wg.Wait()
	if got := c.Get("shared"); got != goroutines*each {
		t.Errorf("counter = %d, want %d — updates were lost", got, goroutines*each)
	}
}

func TestConcurrentMixedOperations(t *testing.T) {
	var c Counter
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			for j := 0; j < 500; j++ {
				c.Add("k", 1)
			}
		}()
		go func() {
			defer wg.Done()
			for j := 0; j < 500; j++ {
				_ = c.Get("k")
				_ = c.Snapshot()
			}
		}()
	}
	wg.Wait()
	if got := c.Get("k"); got != 20*500 {
		t.Errorf("counter = %d, want %d", got, 20*500)
	}
}
