package uploadwindow

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestInUse(t *testing.T) {
	cases := []struct {
		name     string
		limit    int
		acquires int
		releases int
		want     int
	}{
		{"fresh", 3, 0, 0, 0},
		{"one_held", 3, 1, 0, 1},
		{"balanced", 3, 2, 2, 0},
		{"partial_release", 4, 3, 1, 2},
		{"full_window", 2, 2, 0, 2},
		{"limit_zero_means_one", 0, 1, 0, 1},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			w := NewWindow(tc.limit)
			for range tc.acquires {
				w.Acquire()
			}
			for range tc.releases {
				w.Release()
			}
			if got := w.InUse(); got != tc.want {
				t.Errorf("InUse() = %d, want %d", got, tc.want)
			}
		})
	}
}

func TestNeverExceedsLimit(t *testing.T) {
	const uploads, limit = 300, 4
	w := NewWindow(limit)

	var live, peak atomic.Int64
	var wg sync.WaitGroup
	for range uploads {
		wg.Add(1)
		go func() {
			defer wg.Done()
			w.Acquire()
			cur := live.Add(1)
			for {
				old := peak.Load()
				if cur <= old || peak.CompareAndSwap(old, cur) {
					break
				}
			}
			live.Add(-1)
			w.Release()
		}()
	}
	wg.Wait()

	if got := peak.Load(); got > limit {
		t.Errorf("peak concurrency = %d, want <= %d", got, limit)
	}
	if got := w.InUse(); got != 0 {
		t.Errorf("InUse() after drain = %d, want 0", got)
	}
}

func TestBlockedAcquirersAllProceed(t *testing.T) {
	w := NewWindow(1)
	w.Acquire() // window is now full

	const waiters = 16
	var wg sync.WaitGroup
	for range waiters {
		wg.Add(1)
		go func() {
			defer wg.Done()
			w.Acquire()
			w.Release()
		}()
	}

	w.Release() // hands the permit to the first waiter
	wg.Wait()   // hangs if a waiter is never woken

	if got := w.InUse(); got != 0 {
		t.Errorf("InUse() = %d, want 0", got)
	}
}
