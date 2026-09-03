package cishardrunner

import (
	"errors"
	"reflect"
	"sync/atomic"
	"testing"
)

var errShardCrashed = errors.New("shard runner crashed")

// runner passes one test per case and crashes on an empty shard.
func runner(calls *int64) func([]string) (int, error) {
	return func(shard []string) (int, error) {
		atomic.AddInt64(calls, 1)
		if len(shard) == 0 {
			return 7, errShardCrashed
		}
		return len(shard), nil
	}
}

func TestRunShards(t *testing.T) {
	cases := []struct {
		name         string
		shards       [][]string
		wantPassed   int
		wantFailures []int
	}{
		{"all_green", [][]string{{"a"}, {"b", "c"}}, 3, []int{}},
		{"one_crashed", [][]string{{"a"}, {}, {"c", "d"}}, 3, []int{1}},
		{"crashed_shard_contributes_no_passes", [][]string{{}}, 0, []int{0}},
		{"failures_sorted", [][]string{{}, {"a"}, {}, {"b"}}, 2, []int{0, 2}},
		{"all_crashed", [][]string{{}, {}}, 0, []int{0, 1}},
		{"no_shards", nil, 0, []int{}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var calls int64
			passed, failures := RunShards(tc.shards, runner(&calls))
			if passed != tc.wantPassed {
				t.Errorf("passed = %d, want %d", passed, tc.wantPassed)
			}
			if !reflect.DeepEqual(failures, tc.wantFailures) {
				t.Errorf("failures = %v, want %v", failures, tc.wantFailures)
			}
			if int(calls) != len(tc.shards) {
				t.Errorf("run called %d times, want %d", calls, len(tc.shards))
			}
		})
	}
}
