// Package eventbusgen — Gopher Workplace challenge.
package eventbusgen

// Bus delivers events of T to subscribed handlers.
// Its zero value is ready to use. Not safe for concurrent use.
type Bus[T any] struct {
	handlers map[int]func(T)
	nextID   int
}

// Subscribe registers h and returns its subscription id.
func (b *Bus[T]) Subscribe(h func(T)) int {
	// TODO(candidate): store the handler and return its id.
	panic("not implemented")
}

// Unsubscribe removes the handler with the given id, reporting
// whether one was removed.
func (b *Bus[T]) Unsubscribe(id int) bool {
	// TODO(candidate): remove the handler.
	panic("not implemented")
}

// Publish delivers v to every subscriber and returns how many
// handlers ran.
func (b *Bus[T]) Publish(v T) int {
	// TODO(candidate): call every handler and count them.
	panic("not implemented")
}
