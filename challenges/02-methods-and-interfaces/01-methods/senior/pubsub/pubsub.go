// Package pubsub — Gopher Workplace challenge.
package pubsub

import "sync"

// PubSub manages async subscriptions.
type PubSub struct {
	mu   sync.RWMutex
	subs map[string][]chan string
}

// New creates a PubSub.
func New() *PubSub {
	return &PubSub{subs: make(map[string][]chan string)}
}

// Subscribe returns a channel that will receive messages for the topic.
func (ps *PubSub) Subscribe(topic string) <-chan string {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	ch := make(chan string, 10)
	ps.subs[topic] = append(ps.subs[topic], ch)
	return ch
}

// Publish sends msg to all channels subscribed to topic.
// Use RLock/RUnlock to read the map. Send without blocking (or assume buffers are large enough).
func (ps *PubSub) Publish(topic, msg string) {
	// TODO(candidate): RLock, iterate subs[topic], send msg, RUnlock.
	panic("not implemented")
}
