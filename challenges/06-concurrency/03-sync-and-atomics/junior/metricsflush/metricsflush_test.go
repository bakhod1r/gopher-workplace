package metricsflush

import (
	"sync"
	"testing"
)

func TestMeter(t *testing.T) {
	cases := []struct {
		name      string
		records   []int64
		wantDrain int64
	}{
		{"nothing_recorded", nil, 0},
		{"single", []int64{3}, 3},
		{"sum", []int64{3, 4}, 7},
		{"with_negative", []int64{10, -4}, 6},
		{"many", []int64{1, 1, 1, 1, 1}, 5},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var m Meter
			for _, v := range tc.records {
				m.Record(v)
			}
			if got := m.Pending(); got != tc.wantDrain {
				t.Errorf("Pending() = %d, want %d", got, tc.wantDrain)
			}
			if got := m.Drain(); got != tc.wantDrain {
				t.Errorf("Drain() = %d, want %d", got, tc.wantDrain)
			}
			if got := m.Pending(); got != 0 {
				t.Errorf("Pending() after Drain = %d, want 0", got)
			}
			if got := m.Drain(); got != 0 {
				t.Errorf("second Drain() = %d, want 0", got)
			}
		})
	}
}

func TestMeterConcurrentFlush(t *testing.T) {
	var m Meter
	const workers = 8
	const per = 500
	const flushers = 2

	drained := make(chan int64, workers*per)
	var wg sync.WaitGroup
	wg.Add(workers + flushers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < per; j++ {
				m.Record(1)
			}
		}()
	}
	for i := 0; i < flushers; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 200; j++ {
				drained <- m.Drain()
			}
		}()
	}
	wg.Wait()
	close(drained)

	var total int64
	for v := range drained {
		total += v
	}
	total += m.Drain()

	if want := int64(workers * per); total != want {
		t.Errorf("total drained = %d, want %d", total, want)
	}
}
