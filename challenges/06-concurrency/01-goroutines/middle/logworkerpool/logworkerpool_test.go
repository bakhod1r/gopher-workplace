package logworkerpool

import (
	"errors"
	"fmt"
	"strings"
	"sync/atomic"
	"testing"
)

var errRejected = errors.New("ingest rejected line")

// shipper rejects lines containing "bad" and counts both calls and the peak
// number of workers running at once.
func shipper(calls, live, peak *int64) func(string) error {
	return func(line string) error {
		n := atomic.AddInt64(live, 1)
		for {
			old := atomic.LoadInt64(peak)
			if n <= old || atomic.CompareAndSwapInt64(peak, old, n) {
				break
			}
		}
		defer atomic.AddInt64(live, -1)
		atomic.AddInt64(calls, 1)
		if strings.Contains(line, "bad") {
			return errRejected
		}
		return nil
	}
}

func lines(n int) []string {
	out := make([]string, n)
	for i := range out {
		out[i] = fmt.Sprintf("line-%d", i)
	}
	return out
}

func TestShipLines(t *testing.T) {
	cases := []struct {
		name      string
		lines     []string
		workers   int
		wantErrAt []int
	}{
		{"all_accepted", lines(6), 3, nil},
		{"one_rejected", []string{"ok", "bad-1", "ok"}, 2, []int{1}},
		{"more_workers_than_lines", []string{"ok", "bad-1"}, 16, []int{1}},
		{"single_worker", lines(5), 1, nil},
		{"non_positive_workers", []string{"ok", "bad-1"}, 0, []int{1}},
		{"all_rejected", []string{"bad-a", "bad-b"}, 2, []int{0, 1}},
		{"no_lines", nil, 4, nil},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var calls, live, peak int64
			got := ShipLines(tc.lines, tc.workers, shipper(&calls, &live, &peak))
			if len(got) != len(tc.lines) {
				t.Fatalf("len = %d, want %d", len(got), len(tc.lines))
			}
			if int(calls) != len(tc.lines) {
				t.Errorf("ship called %d times, want %d", calls, len(tc.lines))
			}
			limit := int64(tc.workers)
			if limit < 1 {
				limit = 1
			}
			if peak > limit {
				t.Errorf("peak concurrency = %d, want <= %d", peak, limit)
			}
			if live != 0 {
				t.Errorf("%d ships still running after return", live)
			}
			bad := map[int]bool{}
			for _, i := range tc.wantErrAt {
				bad[i] = true
			}
			for i, err := range got {
				if bad[i] != (err != nil) {
					t.Errorf("errs[%d] = %v, wantErr = %v", i, err, bad[i])
				}
			}
		})
	}
}
