package shutdownflag

import (
	"sync"
	"testing"
)

func TestShutdownFlag(t *testing.T) {
	cases := []struct {
		name     string
		requests int
		claims   int
		want     bool
	}{
		{"running", 0, 0, false},
		{"signal_received", 1, 0, true},
		{"duplicate_signal", 2, 0, true},
		{"claimed", 0, 1, true},
		{"signal_then_claim", 1, 1, true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var f ShutdownFlag
			for i := 0; i < tc.requests; i++ {
				f.Request()
			}
			for i := 0; i < tc.claims; i++ {
				f.ClaimShutdown()
			}
			if got := f.Requested(); got != tc.want {
				t.Errorf("Requested() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestShutdownFlagSingleDrainer(t *testing.T) {
	var f ShutdownFlag
	const workers = 16
	winners := make(chan bool, workers)
	var wg sync.WaitGroup
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			winners <- f.ClaimShutdown()
			f.Requested()
		}()
	}
	wg.Wait()
	close(winners)

	n := 0
	for w := range winners {
		if w {
			n++
		}
	}
	if n != 1 {
		t.Errorf("goroutines that claimed shutdown = %d, want 1", n)
	}
	if !f.Requested() {
		t.Error("Requested() = false, want true")
	}
}
