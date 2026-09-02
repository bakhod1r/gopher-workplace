package backupupload

import "testing"

func TestTotalUploaded(t *testing.T) {
	size := func(shard string) int { return len(shard) }

	cases := []struct {
		name    string
		shards  []string
		workers int
		want    int
	}{
		{"two_shards", []string{"a", "bb"}, 2, 3},
		{"single_worker", []string{"abcd"}, 1, 4},
		{"idle_workers", []string{"ab"}, 5, 2},
		{"five_shards", []string{"a", "b", "c", "d", "e"}, 3, 5},
		{"nothing_to_back_up", nil, 3, 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := TotalUploaded(tc.shards, tc.workers, size); got != tc.want {
				t.Errorf("TotalUploaded(%v, %d) = %d, want %d", tc.shards, tc.workers, got, tc.want)
			}
		})
	}
}
