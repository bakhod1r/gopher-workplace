package boundedcrawl

import (
	"fmt"
	"reflect"
	"sync/atomic"
	"testing"
)

// fetcher records the peak number of concurrent fetches and returns 404 for
// URLs ending in "/gone", 200 otherwise.
func fetcher(live, peak *int64) func(string) int {
	return func(url string) int {
		n := atomic.AddInt64(live, 1)
		for {
			old := atomic.LoadInt64(peak)
			if n <= old || atomic.CompareAndSwapInt64(peak, old, n) {
				break
			}
		}
		defer atomic.AddInt64(live, -1)
		if len(url) >= 5 && url[len(url)-5:] == "/gone" {
			return 404
		}
		return 200
	}
}

func urls(n int) []string {
	out := make([]string, n)
	for i := range out {
		out[i] = fmt.Sprintf("/page/%d", i)
	}
	return out
}

func TestCrawlPages(t *testing.T) {
	cases := []struct {
		name        string
		urls        []string
		maxParallel int
		want        []int
	}{
		{"mixed_statuses", []string{"/", "/gone"}, 2, []int{200, 404}},
		{"limit_one_is_sequential", urls(6), 1, []int{200, 200, 200, 200, 200, 200}},
		{"limit_below_input", urls(9), 3, []int{200, 200, 200, 200, 200, 200, 200, 200, 200}},
		{"limit_above_input", urls(2), 50, []int{200, 200}},
		{"unlimited", urls(4), 0, []int{200, 200, 200, 200}},
		{"empty", nil, 4, []int{}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var live, peak int64
			got := CrawlPages(tc.urls, tc.maxParallel, fetcher(&live, &peak))
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("CrawlPages() = %v, want %v", got, tc.want)
			}
			if tc.maxParallel > 0 && peak > int64(tc.maxParallel) {
				t.Errorf("peak concurrency = %d, want <= %d", peak, tc.maxParallel)
			}
			if live != 0 {
				t.Errorf("%d fetches still in flight after return", live)
			}
		})
	}
}
