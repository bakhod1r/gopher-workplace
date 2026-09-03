package drainwait

import (
	"sync"
	"testing"
)

func TestInflightCount(t *testing.T) {
	cases := []struct {
		name   string
		starts int
		dones  int
		want   int
	}{
		{"idle", 0, 0, 0},
		{"one_open", 1, 0, 1},
		{"one_closed", 1, 1, 0},
		{"three_open", 3, 0, 3},
		{"three_two_closed", 3, 2, 1},
		{"balanced", 5, 5, 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			d := NewDrain()
			for range tc.starts {
				d.Start()
			}
			for range tc.dones {
				d.Done()
			}
			if got := d.Inflight(); got != tc.want {
				t.Errorf("Inflight() = %d, want %d", got, tc.want)
			}
		})
	}
}

func TestWaitReturnsImmediatelyWhenIdle(t *testing.T) {
	d := NewDrain()
	done := make(chan struct{})
	go func() {
		d.Wait()
		close(done)
	}()
	<-done
}

func TestWaitBlocksUntilDrained(t *testing.T) {
	d := NewDrain()
	const requests = 32

	for range requests {
		d.Start()
	}

	waited := make(chan struct{})
	go func() {
		d.Wait()
		close(waited)
	}()

	var wg sync.WaitGroup
	for range requests {
		wg.Add(1)
		go func() {
			defer wg.Done()
			d.Done()
		}()
	}
	wg.Wait()

	<-waited // hangs (and the test times out) if Done never wakes the waiter

	if got := d.Inflight(); got != 0 {
		t.Errorf("Inflight() after drain = %d, want 0", got)
	}
}

func TestManyWaitersAllWakeUp(t *testing.T) {
	d := NewDrain()
	d.Start()

	const waiters = 8
	var wg sync.WaitGroup
	for range waiters {
		wg.Add(1)
		go func() {
			defer wg.Done()
			d.Wait()
		}()
	}

	d.Done()
	wg.Wait() // hangs if only one waiter is woken
}
