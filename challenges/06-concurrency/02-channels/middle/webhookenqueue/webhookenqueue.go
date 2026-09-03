// Package webhookenqueue — Gopher Workplace challenge.
package webhookenqueue

// Delivery is one webhook attempt waiting to be sent to a customer endpoint.
type Delivery struct {
	ID       string
	Endpoint string
}

// EnqueueDeliveries offers every delivery in the batch to the outbound queue
// without ever blocking. The ids that made it onto the queue are returned as
// accepted; the ids that found the queue full are returned as dropped, so the
// caller can shed them to the dead-letter table instead of stalling the
// producing request.
//
// Both slices are non-nil and keep the batch order.
//
// Examples:
//
//	EnqueueDeliveries(queue cap 2, [a b c]) => [a b], [c]
//	EnqueueDeliveries(queue cap 4, [a])     => [a], []
//	EnqueueDeliveries(unbuffered queue, [a]) => [], [a]
func EnqueueDeliveries(queue chan<- Delivery, batch []Delivery) (accepted, dropped []string) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
