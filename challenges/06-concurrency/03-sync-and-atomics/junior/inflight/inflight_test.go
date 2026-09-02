package inflight

import (
	"sync"
	"testing"
)

func TestGauge(t *testing.T) {
	cases := []struct {
		name   string
		enters int
		exits  int
		want   int64
	}{
		{"idle", 0, 0, 0},
		{"one_in_flight", 1, 0, 1},
		{"completed", 1, 1, 0},
		{"two_in_flight", 2, 0, 2},
		{"partly_drained", 5, 3, 2},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var g Gauge
			for i := 0; i < tc.enters; i++ {
				g.Enter()
			}
			for i := 0; i < tc.exits; i++ {
				g.Exit()
			}
			if got := g.Current(); got != tc.want {
				t.Errorf("Current() = %d, want %d", got, tc.want)
			}
		})
	}
}

func TestGaugeConcurrent(t *testing.T) {
	var g Gauge
	const handlers = 16
	const per = 200
	var wg sync.WaitGroup
	wg.Add(handlers)
	for i := 0; i < handlers; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < per; j++ {
				g.Enter()
				if got := g.Current(); got < 1 {
					t.Errorf("Current() = %d while a request is in flight", got)
				}
				g.Exit()
			}
		}()
	}
	wg.Wait()

	if got := g.Current(); got != 0 {
		t.Errorf("Current() = %d after drain, want 0", got)
	}
}
