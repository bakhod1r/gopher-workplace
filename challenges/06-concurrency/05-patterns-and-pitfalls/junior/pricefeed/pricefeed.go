// Package pricefeed — Gopher Workplace challenge.
package pricefeed

// PriceFeed streams an endless sequence of quotes starting at base and
// increasing by one. It stops and closes its channel as soon as done is
// closed, so the goroutine never outlives its consumer.
//
// Examples:
//
//	PriceFeed(done, 100)  => yields 100, 101, 102, ... until done closes
//	PriceFeed(done, 0)    => yields 0, 1, 2, ... until done closes
//	close(done)           => the feed channel closes
func PriceFeed(done <-chan struct{}, base int) <-chan int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
