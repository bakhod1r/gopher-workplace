package logcounts

import (
	"reflect"
	"strings"
	"testing"
)

func TestCountLevels(t *testing.T) {
	level := func(line string) string {
		fields := strings.Fields(line)
		if len(fields) == 0 {
			return ""
		}
		return fields[0]
	}

	cases := []struct {
		name   string
		chunks [][]string
		want   map[string]int
	}{
		{"same_level_two_chunks", [][]string{{"ERR a"}, {"ERR b"}}, map[string]int{"ERR": 2}},
		{"two_levels_one_chunk", [][]string{{"ERR a", "INFO b"}}, map[string]int{"ERR": 1, "INFO": 1}},
		{"three_chunks", [][]string{{"ERR a"}, {"INFO b", "INFO c"}, {"WARN d"}}, map[string]int{"ERR": 1, "INFO": 2, "WARN": 1}},
		{"empty_chunk", [][]string{{}, {"ERR a"}}, map[string]int{"ERR": 1}},
		{"no_chunks", nil, map[string]int{}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := CountLevels(tc.chunks, level)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("CountLevels(%v) = %v, want %v", tc.chunks, got, tc.want)
			}
		})
	}
}
