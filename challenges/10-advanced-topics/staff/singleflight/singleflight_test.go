package singleflight

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestDoReturnsTheResult(t *testing.T) {
	var g Group
	if got := g.Do("a", func() int { return 42 }); got != 42 {
		t.Errorf("Do = %d, want 42", got)
	}
}

func TestDoRunsAgainAfterCompletion(t *testing.T) {
	var g Group
	var calls atomic.Int64
	fn := func() int { calls.Add(1); return 1 }
	g.Do("a", fn)
	g.Do("a", fn)
	if got := calls.Load(); got != 2 {
		t.Errorf("fn ran %d times, want 2: Do must not cache across completed calls", got)
	}
}

func TestDoDeduplicatesConcurrentCallers(t *testing.T) {
	var g Group
	var calls atomic.Int64
	release := make(chan struct{})
	fn := func() int {
		calls.Add(1)
		<-release
		return 7
	}

	const workers = 32
	var wg sync.WaitGroup
	results := make([]int, workers)
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func(i int) {
			defer wg.Done()
			results[i] = g.Do("hot", fn)
		}(i)
	}

	// let every goroutine reach Do before the fetch completes
	time.Sleep(50 * time.Millisecond)
	close(release)
	wg.Wait()

	if got := calls.Load(); got != 1 {
		t.Errorf("fn ran %d times, want 1: the callers were not deduplicated", got)
	}
	for i, v := range results {
		if v != 7 {
			t.Fatalf("worker %d got %d, want 7", i, v)
		}
	}
}

func TestDoDifferentKeysRunInParallel(t *testing.T) {
	var g Group
	var calls atomic.Int64
	var wg sync.WaitGroup
	const keys = 8
	wg.Add(keys)
	for i := 0; i < keys; i++ {
		go func(i int) {
			defer wg.Done()
			g.Do(string(rune('a'+i)), func() int { calls.Add(1); return i })
		}(i)
	}
	wg.Wait()
	if got := calls.Load(); got != keys {
		t.Errorf("fn ran %d times, want %d: distinct keys must not block each other", got, keys)
	}
}

func TestDoMixedKeysUnderLoad(t *testing.T) {
	var g Group
	var wg sync.WaitGroup
	const workers = 16
	wg.Add(workers)
	for w := 0; w < workers; w++ {
		go func(w int) {
			defer wg.Done()
			for i := 0; i < 100; i++ {
				key := string(rune('a' + i%4))
				want := i % 4
				if got := g.Do(key, func() int { return want }); got != want {
					panic("wrong value for key")
				}
			}
		}(w)
	}
	wg.Wait()
}
