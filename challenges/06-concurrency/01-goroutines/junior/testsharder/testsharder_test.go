package testsharder

import (
	"reflect"
	"testing"
)

func TestShardDurations(t *testing.T) {
	cases := []struct {
		name      string
		durations []int
		perShard  int
		want      []int
	}{
		{"even_shards", []int{10, 20, 30, 40}, 2, []int{30, 70}},
		{"ragged_tail", []int{10, 20, 30}, 2, []int{30, 30}},
		{"one_per_shard", []int{5, 6}, 1, []int{5, 6}},
		{"single_shard", []int{1, 2, 3}, 10, []int{6}},
		{"bad_size", []int{10}, 0, []int(nil)},
		{"empty", []int{}, 4, []int{}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := ShardDurations(tc.durations, tc.perShard); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("ShardDurations(%v) = %v, want %v", tc.durations, got, tc.want)
			}
		})
	}
}
