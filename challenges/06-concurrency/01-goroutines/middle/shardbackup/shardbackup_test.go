package shardbackup

import (
	"errors"
	"sync/atomic"
	"testing"
)

var errEmptyShard = errors.New("empty shard")

// uploader rejects zero-sized shards and counts every attempt.
func uploader(calls *int64) func(Shard) error {
	return func(sh Shard) error {
		atomic.AddInt64(calls, 1)
		if sh.Size == 0 {
			return errEmptyShard
		}
		return nil
	}
}

func TestUploadShards(t *testing.T) {
	cases := []struct {
		name      string
		shards    []Shard
		wantErrAt []int
		wantCalls int64
	}{
		{"all_land", []Shard{{"a", 1}, {"b", 2}, {"c", 3}}, nil, 3},
		{"one_empty", []Shard{{"a", 1}, {"b", 0}}, []int{1}, 2},
		{"all_empty", []Shard{{"a", 0}, {"b", 0}}, []int{0, 1}, 2},
		{"single", []Shard{{"a", 1}}, nil, 1},
		{"first_failure_does_not_stop_rest", []Shard{{"a", 0}, {"b", 1}, {"c", 0}}, []int{0, 2}, 3},
		{"empty_set", nil, nil, 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var calls int64
			got := UploadShards(tc.shards, uploader(&calls))
			if len(got) != len(tc.shards) {
				t.Fatalf("len(UploadShards()) = %d, want %d", len(got), len(tc.shards))
			}
			if calls != tc.wantCalls {
				t.Errorf("upload called %d times, want %d", calls, tc.wantCalls)
			}
			bad := map[int]bool{}
			for _, i := range tc.wantErrAt {
				bad[i] = true
			}
			for i, err := range got {
				if bad[i] && !errors.Is(err, errEmptyShard) {
					t.Errorf("errs[%d] = %v, want errEmptyShard", i, err)
				}
				if !bad[i] && err != nil {
					t.Errorf("errs[%d] = %v, want nil", i, err)
				}
			}
		})
	}
}
