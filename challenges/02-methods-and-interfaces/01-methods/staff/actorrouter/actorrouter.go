// Package actorrouter — Gopher Workplace challenge.
package actorrouter

// Worker just processes an int.
type Worker struct {
	Inbox chan int
}

// Router routes messages to workers round-robin.
type Router struct {
	workers []*Worker
	idx     int
}

// Route sends msg to the next worker.
func (r *Router) Route(msg int) {
	// TODO(candidate): send msg to workers[idx], then increment idx (modulo len)
	panic("not implemented")
}
