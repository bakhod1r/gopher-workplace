# The Non-Blocking Send

## Intuition

`queue <- d` has one failure mode that does not look like failure: it waits.
On a request path, waiting is the worst outcome — a customer's dead endpoint
becomes our latency. `select` with a `default` arm converts "wait" into "tell
me you couldn't", which is a decision the caller can act on.

`default` runs when *no other arm can proceed immediately*. For a send arm,
"can proceed" means there is buffer space, or a receiver already parked. That
is exactly the definition of "the queue has room".

## Approach

1. Start both result slices as empty, non-nil.
2. For each delivery, `select` on the send with a `default`.
3. Send path: append the id to `accepted`.
4. `default` path: append the id to `dropped`.
5. Never break early — later deliveries may still fit if a consumer drains.

## Solution

```go
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
	accepted = []string{}
	dropped = []string{}

	for _, d := range batch {
		select {
		case queue <- d:
			accepted = append(accepted, d.ID)
		default:
			dropped = append(dropped, d.ID)
		}
	}
	return accepted, dropped
}
```

## Walkthrough

- Capacity 2 with three deliveries: `a` and `b` take the two buffer slots, then
  the send arm for `c` cannot proceed and `default` runs.
- Capacity 3 already holding two items has exactly one slot: `a` fits, `b` does not.
- An unbuffered queue with no parked receiver can never accept a send that has
  no partner, so every delivery is dropped — correct, and worth understanding.
- An empty batch runs the loop zero times and returns two empty, non-nil slices.

## Pitfalls

- Returning on the first drop loses deliveries that would have fit once a
  consumer freed a slot.
- Leaving `accepted`/`dropped` as nil makes callers write `if len(x) > 0` guards
  around JSON encoding, and breaks `reflect.DeepEqual` against `[]string{}`.
- `default` in a *blocking* loop is a busy-wait; here the loop is bounded by the
  batch, so it spins zero times.
- Taking `chan Delivery` instead of `chan<- Delivery` gives this function the
  power to close a queue it does not own.
