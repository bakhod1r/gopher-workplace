package latencypeak

import (
	"sync"
	"testing"
)

func TestPeakTracker(t *testing.T) {
	cases := []struct {
		name string
		vals []int64
		want int64
	}{
		{"no_requests", nil, 0},
		{"single_request", []int64{5}, 5},
		{"faster_second", []int64{5, 3}, 5},
		{"rising", []int64{1, 2, 9}, 9},
		{"clock_skew_negatives", []int64{-4, -9}, 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var p PeakTracker
			for _, v := range tc.vals {
				p.Observe(v)
			}
			if got := p.Peak(); got != tc.want {
				t.Errorf("Peak() = %d, want %d", got, tc.want)
			}
		})
	}
}

func TestPeakTrackerConcurrent(t *testing.T) {
	var p PeakTracker
	const handlers = 8
	const per = 200
	var wg sync.WaitGroup
	wg.Add(handlers)
	for i := 0; i < handlers; i++ {
		go func(i int) {
			defer wg.Done()
			for j := 0; j < per; j++ {
				p.Observe(int64(i*per + j))
				p.Peak()
			}
		}(i)
	}
	wg.Wait()
	if got, want := p.Peak(), int64(handlers*per-1); got != want {
		t.Errorf("Peak() = %d, want %d", got, want)
	}
}
