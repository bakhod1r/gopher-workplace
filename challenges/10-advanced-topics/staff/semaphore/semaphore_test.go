package semaphore

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestAcquireUpToCapacity(t *testing.T) {
	done := make(chan struct{})
	defer close(done)
	s := NewSem(2)
	if !s.Acquire(done) || !s.Acquire(done) {
		t.Fatal("the first two acquires must succeed")
	}
	if s.Held() != 2 {
		t.Errorf("Held = %d, want 2", s.Held())
	}
}

func TestAcquireBlocksWhenFull(t *testing.T) {
	done := make(chan struct{})
	s := NewSem(1)
	if !s.Acquire(done) {
		t.Fatal("the first acquire must succeed")
	}
	got := make(chan bool, 1)
	go func() { got <- s.Acquire(done) }()
	select {
	case <-got:
		t.Fatal("the second acquire returned while the semaphore was full")
	case <-time.After(50 * time.Millisecond):
	}
	close(done)
	if ok := <-got; ok {
		t.Error("the cancelled acquire reported true, want false")
	}
}

func TestReleaseUnblocksAWaiter(t *testing.T) {
	done := make(chan struct{})
	defer close(done)
	s := NewSem(1)
	s.Acquire(done)
	got := make(chan bool, 1)
	go func() { got <- s.Acquire(done) }()
	time.Sleep(20 * time.Millisecond)
	s.Release()
	select {
	case ok := <-got:
		if !ok {
			t.Error("the waiter reported false, want true")
		}
	case <-time.After(2 * time.Second):
		t.Error("the waiter was never released")
	}
}

func TestConcurrencyIsBounded(t *testing.T) {
	done := make(chan struct{})
	defer close(done)
	const limit = 4
	s := NewSem(limit)
	var inFlight, peak atomic.Int64
	var wg sync.WaitGroup
	const workers = 32
	wg.Add(workers)
	for w := 0; w < workers; w++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 50; i++ {
				if !s.Acquire(done) {
					return
				}
				n := inFlight.Add(1)
				for {
					p := peak.Load()
					if n <= p || peak.CompareAndSwap(p, n) {
						break
					}
				}
				inFlight.Add(-1)
				s.Release()
			}
		}()
	}
	wg.Wait()
	if peak.Load() > limit {
		t.Errorf("peak concurrency was %d, want at most %d", peak.Load(), limit)
	}
}
