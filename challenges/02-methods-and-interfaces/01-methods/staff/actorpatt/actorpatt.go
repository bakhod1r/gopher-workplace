// Package actorpatt — Gopher Workplace challenge.
package actorpatt

// CounterActor processes messages serially via a channel.
type CounterActor struct {
	count int
	msgs  chan func(*int)
}

// New returns a started actor.
func New() *CounterActor {
	a := &CounterActor{msgs: make(chan func(*int), 10)}
	go a.run()
	return a
}

func (a *CounterActor) run() {
	for fn := range a.msgs {
		fn(&a.count)
	}
}

// Add sends a message to increment the counter asynchronously.
func (a *CounterActor) Add(n int) {
	// TODO(candidate): send a func(*int) that adds n to the counter
	panic("not implemented")
}
