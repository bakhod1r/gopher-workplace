package singleflight

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestDoReturnsValue(t *testing.T) {
	g := NewGroup()
	got := g.Do("k", LoaderFunc(func(key string) string { return "v:" + key }))
	if got != "v:k" {
		t.Errorf("Do = %q, want \"v:k\"", got)
	}
}

func TestConcurrentCallsCollapse(t *testing.T) {
	g := NewGroup()
	var loads int64
	release := make(chan struct{})

	loader := LoaderFunc(func(key string) string {
		atomic.AddInt64(&loads, 1)
		<-release // hold the flight open
		return "v"
	})

	const n = 100
	var wg sync.WaitGroup
	results := make([]string, n)
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			results[i] = g.Do("k", loader)
		}(i)
	}

	// Give the goroutines time to pile up on the same key.
	time.Sleep(50 * time.Millisecond)
	close(release)
	wg.Wait()

	if loads != 1 {
		t.Errorf("loader ran %d times, want 1", loads)
	}
	for i, r := range results {
		if r != "v" {
			t.Fatalf("results[%d] = %q, want \"v\"", i, r)
		}
	}
}

func TestDistinctKeysDoNotBlock(t *testing.T) {
	g := NewGroup()
	var loads int64

	loader := LoaderFunc(func(key string) string {
		atomic.AddInt64(&loads, 1)
		return key
	})

	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			k := string(rune('a' + i))
			if got := g.Do(k, loader); got != k {
				t.Errorf("Do(%q) = %q", k, got)
			}
		}(i)
	}
	wg.Wait()

	if loads != 2 {
		t.Errorf("loader ran %d times, want 2", loads)
	}
}

func TestSequentialCallsReload(t *testing.T) {
	g := NewGroup()
	var loads int64
	loader := LoaderFunc(func(key string) string {
		atomic.AddInt64(&loads, 1)
		return "v"
	})

	g.Do("k", loader)
	g.Do("k", loader)

	if loads != 2 {
		t.Errorf("loader ran %d times, want 2 (no caching between flights)", loads)
	}
}

func TestFlightIsCleanedUp(t *testing.T) {
	g := NewGroup()
	g.Do("k", LoaderFunc(func(string) string { return "v" }))

	g.mu.Lock()
	n := len(g.calls)
	g.mu.Unlock()
	if n != 0 {
		t.Errorf("%d in-flight entries left behind, want 0", n)
	}
}
