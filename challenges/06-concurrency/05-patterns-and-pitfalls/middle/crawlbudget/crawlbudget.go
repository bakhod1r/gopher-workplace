// Package crawlbudget — Gopher Workplace challenge.
package crawlbudget

import "context"

// Page is one crawled document.
type Page struct {
	URL   string
	Bytes int
}

// CrawlPages fetches every URL with at most limit fetches in flight at once.
// A buffered channel is used as a counting semaphore: acquiring a slot is a
// send, releasing it is a receive, and the crawl can never exceed the politeness
// budget the site owner agreed to.
//
// Pages come back in the order the URLs were given. The first fetch error
// cancels the crawl and is returned with a nil slice.
//
// Examples:
//
//	CrawlPages(live ctx, 6 urls, 2, fetch)        => 6 pages in order
//	CrawlPages(live ctx, urls with "bad", 3, get) => errFetch
//	CrawlPages(cancelled ctx, urls, 2, fetch)     => context.Canceled
func CrawlPages(ctx context.Context, urls []string, limit int, fetch func(context.Context, string) (Page, error)) ([]Page, error) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
