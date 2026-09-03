package semaphoregate

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestAcquireRelease(t *testing.T) {
	s := New(2)
	if got := s.Available(); got != 2 {
		t.Errorf("Available = %d, want 2", got)
	}
	s.Acquire()
	if got := s.Available(); got != 1 {
		t.Errorf("Available = %d, want 1", got)
	}
	s.Release()
	if got := s.Available(); got != 2 {
		t.Errorf("Available = %d, want 2", got)
	}
}

func TestTryAcquire(t *testing.T) {
	s := New(1)
	if !s.TryAcquire() {
		t.Fatal("TryAcquire on a free semaphore = false, want true")
	}
	if s.TryAcquire() {
		t.Error("TryAcquire with no permits left = true, want false")
	}
	s.Release()
	if !s.TryAcquire() {
		t.Error("TryAcquire after Release = false, want true")
	}
}

func TestExtraReleaseIsDropped(t *testing.T) {
	s := New(2)
	s.Release()
	s.Release()
	s.Release()
	if got := s.Available(); got != 2 {
		t.Errorf("Available = %d, want 2 — an unmatched Release must not inflate the permit count", got)
	}
}

func TestNonPositiveGivesOnePermit(t *testing.T) {
	for _, n := range []int{0, -3} {
		s := New(n)
		if got := s.Available(); got != 1 {
			t.Errorf("New(%d).Available() = %d, want 1", n, got)
		}
	}
}

func TestSemaphoreBoundsConcurrency(t *testing.T) {
	const limit = 3
	s := New(limit)
	var live, peak atomic.Int64
	var wg sync.WaitGroup
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s.Acquire()
			defer s.Release()
			n := live.Add(1)
			for {
				p := peak.Load()
				if n <= p || peak.CompareAndSwap(p, n) {
					break
				}
			}
			live.Add(-1)
		}()
	}
	wg.Wait()
	if peak.Load() > limit {
		t.Errorf("peak concurrency = %d, want at most %d", peak.Load(), limit)
	}
	if got := s.Available(); got != limit {
		t.Errorf("Available = %d, want %d — every permit must come back", got, limit)
	}
}
