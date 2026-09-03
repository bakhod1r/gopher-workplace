// Package boundedcrawl — Gopher Workplace challenge.
package boundedcrawl

// CrawlPages fetches every URL and returns the status codes in input order,
// never running more than maxParallel fetches at the same time. The politeness
// limit protects the origin from being knocked over by its own crawler; a
// maxParallel of zero or less means "no limit".
//
// Examples:
//
//	CrawlPages([]string{"/", "/gone"}, 2, fetch)  => [200 404]
//	CrawlPages([]string{"/"}, 1, fetch)           => [200]
//	CrawlPages(nil, 4, fetch)                     => []
func CrawlPages(urls []string, maxParallel int, fetch func(url string) int) []int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
