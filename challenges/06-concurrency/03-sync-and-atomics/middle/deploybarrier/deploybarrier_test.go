package deploybarrier

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestSingleParticipantPassesImmediately(t *testing.T) {
	cases := []struct {
		name  string
		n     int
		waits int
		want  int
	}{
		{"one_wait", 1, 1, 1},
		{"two_waits", 1, 2, 2},
		{"three_waits", 1, 3, 3},
		{"zero_means_one", 0, 1, 1},
		{"negative_means_one", -4, 2, 2},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			b := NewBarrier(tc.n)
			for range tc.waits {
				b.Wait()
			}
			if got := b.Phase(); got != tc.want {
				t.Errorf("Phase() = %d, want %d", got, tc.want)
			}
		})
	}
}

func TestNobodyPassesUntilEveryoneArrives(t *testing.T) {
	const regions = 8
	b := NewBarrier(regions)

	var passed atomic.Int64
	var wg sync.WaitGroup
	for range regions - 1 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			b.Wait()
			passed.Add(1)
		}()
	}

	if got := passed.Load(); got != 0 {
		t.Fatalf("%d participants passed before the last arrival", got)
	}

	b.Wait() // the last region arrives and releases everyone
	wg.Wait()

	if got := passed.Load(); got != regions-1 {
		t.Errorf("passed = %d, want %d", got, regions-1)
	}
	if got := b.Phase(); got != 1 {
		t.Errorf("Phase() = %d, want 1", got)
	}
}

func TestBarrierIsReusableAcrossPhases(t *testing.T) {
	const regions, phases = 6, 4
	b := NewBarrier(regions)

	var work atomic.Int64
	var wg sync.WaitGroup
	for range regions {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range phases {
				work.Add(1)
				b.Wait()
			}
		}()
	}
	wg.Wait()

	if got := work.Load(); got != regions*phases {
		t.Errorf("work units = %d, want %d", got, regions*phases)
	}
	if got := b.Phase(); got != phases {
		t.Errorf("Phase() = %d, want %d", got, phases)
	}
}
