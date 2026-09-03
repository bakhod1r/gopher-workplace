// Package pipeline — Gopher Workplace challenge.
package pipeline

import "sync"

// Stage returns a channel carrying each input doubled, computed by
// workers goroutines, and closed once the input drains or done is closed.
//
// Every goroutine must exit on done, whether it is blocked receiving or
// blocked sending.
//
// Examples:
//
//	Stage(done, in, 4) => a channel of doubled values
func Stage(done <-chan struct{}, in <-chan int, workers int) <-chan int {
	panic("not implemented")
}
