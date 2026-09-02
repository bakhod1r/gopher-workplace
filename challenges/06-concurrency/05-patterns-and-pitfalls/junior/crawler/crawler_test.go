package crawler

import "testing"

func TestCrawlSite(t *testing.T) {
	fetch := func(u string) int { return len(u) }

	cases := []struct {
		name  string
		urls  []string
		limit int
		want  []int
	}{
		{"two_limit_two", []string{"/a", "/bb"}, 2, []int{2, 3}},
		{"order_preserved", []string{"/aaaa", "/b", "/ccc"}, 2, []int{5, 2, 4}},
		{"single_limit_one", []string{"/x"}, 1, []int{2}},
		{"limit_above_len", []string{"/a", "/b"}, 9, []int{2, 2}},
		{"empty_frontier", nil, 4, []int{}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := CrawlSite(tc.urls, tc.limit, fetch)
			if len(got) != len(tc.want) {
				t.Fatalf("CrawlSite(%v, %d) = %v, want %v", tc.urls, tc.limit, got, tc.want)
			}
			for i := range got {
				if got[i] != tc.want[i] {
					t.Fatalf("CrawlSite(%v, %d) = %v, want %v", tc.urls, tc.limit, got, tc.want)
				}
			}
		})
	}
}
