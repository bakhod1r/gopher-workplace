package jobmetrics

import (
	"sync"
	"testing"
)

func TestJobMetricsAdd(t *testing.T) {
	cases := []struct {
		name   string
		deltas []int64
		reset  bool
		want   int64
	}{
		{"single_job", []int64{3}, false, 3},
		{"correction", []int64{2, -5}, false, -3},
		{"idle_pool", nil, false, 0},
		{"accumulated", []int64{1, 2, 3, 4}, false, 10},
		{"after_flush", []int64{7, 8}, true, 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var m JobMetrics
			for _, d := range tc.deltas {
				m.Add(d)
			}
			if tc.reset {
				m.Reset()
			}
			if got := m.Processed(); got != tc.want {
				t.Errorf("Processed() = %d, want %d", got, tc.want)
			}
		})
	}
}

func TestJobMetricsConcurrent(t *testing.T) {
	var m JobMetrics
	const workers = 8
	const per = 500
	var wg sync.WaitGroup
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < per; j++ {
				m.Add(1)
				m.Processed()
			}
		}()
	}
	wg.Wait()
	if got, want := m.Processed(), int64(workers*per); got != want {
		t.Errorf("Processed() = %d, want %d", got, want)
	}
}
