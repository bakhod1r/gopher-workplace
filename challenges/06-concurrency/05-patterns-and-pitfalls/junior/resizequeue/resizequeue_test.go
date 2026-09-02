package resizequeue

import "testing"

func TestResizeQueue(t *testing.T) {
	resize := func(key string) string { return key + "-512" }

	cases := []struct {
		name    string
		uploads []string
		workers int
		want    []string
	}{
		{"three_uploads", []string{"a", "b", "c"}, 2, []string{"a-512", "b-512", "c-512"}},
		{"idle_workers", []string{"z"}, 4, []string{"z-512"}},
		{"single_worker", []string{"c", "b", "a"}, 1, []string{"a-512", "b-512", "c-512"}},
		{"duplicates", []string{"x", "x"}, 2, []string{"x-512", "x-512"}},
		{"empty_queue", nil, 3, nil},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := ResizeQueue(tc.uploads, tc.workers, resize)
			if len(got) != len(tc.want) {
				t.Fatalf("ResizeQueue(%v, %d) = %v, want %v", tc.uploads, tc.workers, got, tc.want)
			}
			for i := range got {
				if got[i] != tc.want[i] {
					t.Fatalf("ResizeQueue(%v, %d) = %v, want %v", tc.uploads, tc.workers, got, tc.want)
				}
			}
		})
	}
}
