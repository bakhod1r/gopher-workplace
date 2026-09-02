package filemanifest

import "testing"

func TestFileSizes(t *testing.T) {
	size := func(path string) int { return len(path) }

	cases := []struct {
		name  string
		paths []string
		want  []int
	}{
		{"two_files", []string{"a", "bb"}, []int{1, 2}},
		{"single_file", []string{"abc"}, []int{3}},
		{"order_preserved", []string{"aaaa", "b", "cc"}, []int{4, 1, 2}},
		{"duplicate_paths", []string{"ab", "ab"}, []int{2, 2}},
		{"empty_manifest", nil, []int{}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := FileSizes(tc.paths, size)
			if len(got) != len(tc.want) {
				t.Fatalf("FileSizes(%v) = %v, want %v", tc.paths, got, tc.want)
			}
			for i := range got {
				if got[i] != tc.want[i] {
					t.Fatalf("FileSizes(%v) = %v, want %v", tc.paths, got, tc.want)
				}
			}
		})
	}
}
