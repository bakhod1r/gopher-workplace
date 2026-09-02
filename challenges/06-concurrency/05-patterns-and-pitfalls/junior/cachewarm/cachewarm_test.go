package cachewarm

import "testing"

func TestWarmShards(t *testing.T) {
	keyCount := func(shard string) int { return len(shard) }

	cases := []struct {
		name   string
		shards []string
		want   []int
	}{
		{"two_shards", []string{"a", "bb"}, []int{1, 2}},
		{"single_shard", []string{"xyz"}, []int{3}},
		{"sorted_output", []string{"cccc", "b", "aa"}, []int{1, 2, 4}},
		{"equal_shards", []string{"aa", "bb"}, []int{2, 2}},
		{"no_shards", nil, []int{}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ready := make(chan struct{})
			close(ready)

			got := WarmShards(ready, tc.shards, keyCount)
			if len(got) != len(tc.want) {
				t.Fatalf("WarmShards(%v) = %v, want %v", tc.shards, got, tc.want)
			}
			for i := range got {
				if got[i] != tc.want[i] {
					t.Fatalf("WarmShards(%v) = %v, want %v", tc.shards, got, tc.want)
				}
			}
		})
	}
}
