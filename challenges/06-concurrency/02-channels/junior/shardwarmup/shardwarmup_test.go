package shardwarmup

import "testing"

func TestWaitForShards(t *testing.T) {
	cases := []struct {
		name string
		n    int
		want int
	}{
		{"three_shards", 3, 3},
		{"one_shard", 1, 1},
		{"no_shards", 0, 0},
		{"negative", -2, 0},
		{"many_shards", 16, 16},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := WaitForShards(tc.n); got != tc.want {
				t.Errorf("WaitForShards(%d) = %d, want %d", tc.n, got, tc.want)
			}
		})
	}
}
