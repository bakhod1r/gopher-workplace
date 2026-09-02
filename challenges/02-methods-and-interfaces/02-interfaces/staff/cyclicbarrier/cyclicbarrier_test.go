package cyclicbarrier

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestSingleParty(t *testing.T) {
	b := NewBarrier(1)
	if got := b.Await(); got != 0 {
		t.Errorf("Await = %d, want 0", got)
	}
	if got := b.Await(); got != 1 {
		t.Errorf("Await = %d, want 1", got)
	}
	if b.Round() != 2 {
		t.Errorf("Round = %d, want 2", b.Round())
	}
}

func TestReleasesAllParties(t *testing.T) {
	const parties = 3
	b := NewBarrier(parties)

	var wg sync.WaitGroup
	rounds := make([]int, parties)
	for i := 0; i < parties; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			rounds[i] = b.Await()
		}(i)
	}
	wg.Wait()

	for i, r := range rounds {
		if r != 0 {
			t.Errorf("party %d saw round %d, want 0", i, r)
		}
	}
	if b.Round() != 1 {
		t.Errorf("Round = %d, want 1", b.Round())
	}
}

func TestNoOneRunsAhead(t *testing.T) {
	const parties, rounds = 4, 50
	b := NewBarrier(parties)

	var inRound [rounds]int64
	var wg sync.WaitGroup

	for p := 0; p < parties; p++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for r := 0; r < rounds; r++ {
				got := b.Await()
				if got != r {
					t.Errorf("saw round %d, want %d", got, r)
					return
				}
				atomic.AddInt64(&inRound[r], 1)
			}
		}()
	}
	wg.Wait()

	for r := 0; r < rounds; r++ {
		if inRound[r] != parties {
			t.Fatalf("round %d had %d arrivals, want %d", r, inRound[r], parties)
		}
	}
	if b.Round() != rounds {
		t.Errorf("Round = %d, want %d", b.Round(), rounds)
	}
}

func TestManyRounds(t *testing.T) {
	const parties = 8
	b := NewBarrier(parties)

	var counter int64
	var wg sync.WaitGroup
	for p := 0; p < parties; p++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for r := 0; r < 100; r++ {
				atomic.AddInt64(&counter, 1)
				b.Await()
			}
		}()
	}
	wg.Wait()

	if counter != parties*100 {
		t.Errorf("counter = %d, want %d", counter, parties*100)
	}
	if b.Round() != 100 {
		t.Errorf("Round = %d, want 100", b.Round())
	}
}

func TestZeroPartiesClamped(t *testing.T) {
	b := NewBarrier(0)
	if b.Parties != 1 {
		t.Errorf("Parties = %d, want 1", b.Parties)
	}
	if got := b.Await(); got != 0 {
		t.Errorf("Await = %d, want 0", got)
	}
}

func TestIsWaiter(t *testing.T) {
	var w Waiter = NewBarrier(1)
	if got := w.Await(); got != 0 {
		t.Errorf("Await = %d, want 0", got)
	}
}
