// Package observerbus — Gopher Workplace challenge.
package observerbus

import "sync"

// Handler receives events.
type Handler interface {
	Handle(event string)
}

// HandlerFunc adapts a function to Handler.
type HandlerFunc func(string)

// Handle calls the underlying function.
func (f HandlerFunc) Handle(event string) { f(event) }

// Bus fans events out to subscribers.
type Bus struct {
	mu     sync.Mutex
	nextID int
	subs   map[int]Handler
}

// NewBus returns an empty bus.
func NewBus() *Bus {
	return &Bus{subs: make(map[int]Handler)}
}

// Subscribe registers h and returns a function that unsubscribes it.
// The returned function is safe to call more than once.
func (b *Bus) Subscribe(h Handler) func() {
	// TODO(candidate): register, return an idempotent unsubscribe.
	panic("not implemented")
}

// Publish delivers the event to every current subscriber.
//
// A handler may unsubscribe itself during delivery.
//
// Examples:
//
//	two subscribers => both receive the event
func (b *Bus) Publish(event string) {
	// TODO(candidate): snapshot under the lock, then deliver unlocked.
	panic("not implemented")
}

// Count returns how many subscribers are registered.
func (b *Bus) Count() int {
	// TODO(candidate): read under the lock.
	panic("not implemented")
}
