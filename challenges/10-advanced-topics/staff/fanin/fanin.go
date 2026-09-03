// Package fanin — Gopher Workplace challenge.
package fanin

import "sync"

// Merge returns a channel carrying every value from ins, closed once all
// inputs are drained or done is closed.
//
// Every goroutine Merge starts must exit: an abandoned consumer must not
// leave forwarders blocked on a send forever.
//
// Examples:
//
//	Merge(done, a, b) => a channel with everything from a and b
func Merge(done <-chan struct{}, ins ...<-chan int) <-chan int {
	panic("not implemented")
}
