// Package crawler — Gopher Workplace challenge.
package crawler

// CrawlSite fetches every URL concurrently, with at most limit requests in
// flight, and returns the status codes in the same order as urls.
// limit is >= 1.
//
// Examples:
//
//	CrawlSite([]string{"/a", "/bb"}, 2, statusFn)  => []int{2, 3}
//	CrawlSite([]string{"/x"}, 1, statusFn)         => []int{2}
//	CrawlSite(nil, 4, statusFn)                    => []int{}
func CrawlSite(urls []string, limit int, fetch func(string) int) []int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
