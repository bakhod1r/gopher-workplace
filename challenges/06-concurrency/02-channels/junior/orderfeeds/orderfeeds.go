// Package orderfeeds — Gopher Workplace challenge.
package orderfeeds

// MergeFeeds drains the primary feed completely, then the standby feed, and
// returns every order id in that order.
//
// Examples:
//
//	MergeFeeds(chan 1,2 | chan 3) => [1 2 3]
//	MergeFeeds(closed empty | chan 9) => [9]
//	MergeFeeds(chan 4 | closed empty) => [4]
func MergeFeeds(primary, standby <-chan int) []int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
