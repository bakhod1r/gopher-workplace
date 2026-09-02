// Package fanin — Gopher Workplace challenge.
package fanin

import "sync"

// Merger combines channels into one.
type Merger interface {
	Merge(done <-chan struct{}, ins ...<-chan int) <-chan int
}

// Fan merges channels, honouring a done signal.
type Fan struct{}

// Merge fans every input into one output channel.
//
// Each forwarding goroutine exits when its input closes or done is closed.
// The output channel is closed once every forwarder has exited.
//
// Examples:
//
//	merge two channels and drain => every value, then a closed output
func (Fan) Merge(done <-chan struct{}, ins ...<-chan int) <-chan int {
	// TODO(candidate): one goroutine per input, a WaitGroup, one closer.
	panic("not implemented")
}

var _ = sync.WaitGroup{}
