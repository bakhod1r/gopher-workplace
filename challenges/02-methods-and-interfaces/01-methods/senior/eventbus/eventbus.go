// Package eventbus — Gopher Workplace challenge.
package eventbus

// Bus manages event listeners.
type Bus struct {
	listeners map[string][]func(string)
}

// New creates an event bus.
func New() *Bus {
	return &Bus{listeners: make(map[string][]func(string))}
}

// On registers a listener for an event type.
func (b *Bus) On(eventType string, listener func(data string)) {
	// TODO(candidate): append listener to b.listeners[eventType]
	panic("not implemented")
}

// Emit calls all listeners for an event type synchronously.
func (b *Bus) Emit(eventType string, data string) {
	// TODO(candidate): call each listener in b.listeners[eventType] with data
	panic("not implemented")
}
