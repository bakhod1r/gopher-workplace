package shardrouter

import (
	"reflect"
	"testing"
)

func TestShardIDs(t *testing.T) {
	cases := []struct {
		name   string
		keys   []string
		shards int
		want   []int
	}{
		{"four_shards", []string{"a", "b"}, 4, []int{1, 2}},
		{"single_shard", []string{"a", "b"}, 1, []int{0, 0}},
		{"empty_key", []string{""}, 4, []int{0}},
		{"multi_byte_key", []string{"ab", "abc"}, 4, []int{1, 2}},
		{"bad_shard_count", []string{"a"}, 0, []int(nil)},
		{"empty", []string{}, 4, []int{}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := ShardIDs(tc.keys, tc.shards); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("ShardIDs(%v) = %v, want %v", tc.keys, got, tc.want)
			}
		})
	}
}
