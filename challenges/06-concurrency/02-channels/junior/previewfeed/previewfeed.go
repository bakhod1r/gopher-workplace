// Package previewfeed — Gopher Workplace challenge.
package previewfeed

// PreviewOrders receives at most limit order ids from the feed, stopping
// early if the feed closes first. It always returns a non-nil slice.
//
// Examples:
//
//	PreviewOrders(chan 1,2,3, 2) => [1 2]
//	PreviewOrders(chan 1, 5)    => [1]
//	PreviewOrders(chan 1,2, 0)  => []
func PreviewOrders(feed <-chan int, limit int) []int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
