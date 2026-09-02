package releasegate

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestGateState(t *testing.T) {
	cases := []struct {
		name  string
		opens int
		want  bool
	}{
		{"closed", 0, false},
		{"opened", 1, true},
		{"opened_twice", 2, true},
		{"opened_three_times", 3, true},
		{"still_closed", 0, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewGate()
			for i := 0; i < tc.opens; i++ {
				g.Open()
			}
			if got := g.IsOpen(); got != tc.want {
				t.Errorf("IsOpen() = %v, want %v", got, tc.want)
			}
			if tc.want {
				g.Wait()
			}
		})
	}
}

func TestGateReleasesAllWaiters(t *testing.T) {
	g := NewGate()
	const handlers = 16

	var passed atomic.Int64
	ready := make(chan struct{}, handlers)
	var wg sync.WaitGroup
	wg.Add(handlers)
	for i := 0; i < handlers; i++ {
		go func() {
			defer wg.Done()
			ready <- struct{}{}
			g.Wait()
			passed.Add(1)
		}()
	}
	for i := 0; i < handlers; i++ {
		<-ready
	}

	g.Open()
	wg.Wait()

	if got := passed.Load(); got != handlers {
		t.Errorf("handlers released = %d, want %d", got, handlers)
	}
	if !g.IsOpen() {
		t.Error("IsOpen() = false, want true")
	}
}
