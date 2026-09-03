package metricsbatcher

import (
	"errors"
	"reflect"
	"sync/atomic"
	"testing"
)

var errRejected = errors.New("collector rejected batch")

// collector rejects any batch containing a negative point.
func collector(calls *int64) func([]int) error {
	return func(batch []int) error {
		atomic.AddInt64(calls, 1)
		for _, p := range batch {
			if p < 0 {
				return errRejected
			}
		}
		return nil
	}
}

func TestFlushBatches(t *testing.T) {
	cases := []struct {
		name         string
		points       []int
		batchSize    int
		wantAccepted int
		wantRetry    []int
		wantCalls    int64
	}{
		{"all_accepted", []int{1, 2, 3, 4}, 2, 4, []int{}, 2},
		{"one_batch_rejected", []int{1, -1, 3, 4}, 2, 2, []int{1, -1}, 2},
		{"ragged_last_batch", []int{1, 2, 3, 4, 5}, 2, 5, []int{}, 3},
		{"retry_keeps_order", []int{-1, 2, 3, -4, 5}, 2, 1, []int{-1, 2, 3, -4}, 3},
		{"all_rejected", []int{-1, -2}, 1, 0, []int{-1, -2}, 2},
		{"batch_size_zero", []int{1, 2}, 0, 0, []int{1, 2}, 0},
		{"no_points", nil, 4, 0, []int{}, 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var calls int64
			accepted, retry := FlushBatches(tc.points, tc.batchSize, collector(&calls))
			if accepted != tc.wantAccepted {
				t.Errorf("accepted = %d, want %d", accepted, tc.wantAccepted)
			}
			if !reflect.DeepEqual(retry, tc.wantRetry) {
				t.Errorf("retry = %v, want %v", retry, tc.wantRetry)
			}
			if calls != tc.wantCalls {
				t.Errorf("flush called %d times, want %d", calls, tc.wantCalls)
			}
		})
	}
}
