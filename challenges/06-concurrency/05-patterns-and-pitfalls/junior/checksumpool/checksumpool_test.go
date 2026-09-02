package checksumpool

import (
	"reflect"
	"testing"
)

func TestChecksumFiles(t *testing.T) {
	sum := func(file string) int { return len(file) }

	cases := []struct {
		name    string
		files   []string
		workers int
		want    map[string]int
	}{
		{"single_file", []string{"a"}, 2, map[string]int{"a": 1}},
		{"two_files", []string{"a", "bb"}, 2, map[string]int{"a": 1, "bb": 2}},
		{"one_worker", []string{"a", "bb", "ccc"}, 1, map[string]int{"a": 1, "bb": 2, "ccc": 3}},
		{"duplicate_files", []string{"ab", "ab"}, 3, map[string]int{"ab": 2}},
		{"no_files", nil, 3, map[string]int{}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := ChecksumFiles(tc.files, tc.workers, sum)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("ChecksumFiles(%v, %d) = %v, want %v", tc.files, tc.workers, got, tc.want)
			}
		})
	}
}
