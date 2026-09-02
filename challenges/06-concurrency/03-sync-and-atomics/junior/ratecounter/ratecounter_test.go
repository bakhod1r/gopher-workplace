package ratecounter

import (
	"sync"
	"testing"
)

func TestRateCounter(t *testing.T) {
	cases := []struct {
		name    string
		workers int
		per     int
		want    int
	}{
		{"single_handler", 1, 1, 1},
		{"no_traffic", 4, 0, 0},
		{"two_handlers", 2, 500, 1000},
		{"eight_handlers", 8, 250, 2000},
		{"wide_fanout", 50, 20, 1000},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var c RateCounter
			var wg sync.WaitGroup
			wg.Add(tc.workers)
			for i := 0; i < tc.workers; i++ {
				go func() {
					defer wg.Done()
					for j := 0; j < tc.per; j++ {
						c.Record()
					}
				}()
			}
			wg.Wait()
			if got := c.Hits(); got != tc.want {
				t.Errorf("Hits() = %d, want %d", got, tc.want)
			}
		})
	}
}
