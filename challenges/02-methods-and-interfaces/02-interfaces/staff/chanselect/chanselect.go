// Package chanselect — Gopher Workplace challenge.
package chanselect

import "sort"

// Source exposes a receive-only channel.
type Source interface {
	Chan() <-chan int
}

// ChanSource wraps a channel.
type ChanSource struct {
	C <-chan int
}

// Chan returns the underlying channel.
func (c ChanSource) Chan() <-chan int { return c.C }

// TryRecv attempts a non-blocking receive.
//
// ready is false when nothing was available. When ready is true, ok reports
// whether the channel was still open.
//
// Examples:
//
//	empty open channel => 0, false, false
//	closed channel     => 0, false, true
func TryRecv(ch <-chan int) (v int, ok bool, ready bool) {
	// TODO(candidate): select with default.
	panic("not implemented")
}

// Drain reads both channels until both are closed and returns every value,
// sorted ascending.
func Drain(a, b <-chan int) []int {
	// TODO(candidate): select over both; nil out a channel once it closes.
	panic("not implemented")
}

// FirstReady blocks until either channel yields a value and reports which
// one it came from ("a" or "b"). It returns from = "" when both close first.
func FirstReady(a, b <-chan int) (v int, from string) {
	// TODO(candidate): select; disable closed channels.
	panic("not implemented")
}

var _ = sort.Ints
