# Drop Instead Of Blocking

## Intuition

A `select` with a `default` asks "can this proceed right now" instead of "wait until it can". That turns backpressure from a stall into a decision the caller gets to make.

## Approach

1. `select` with the send as the only case.
2. `default` returns false.

## Solution

```go
// Offer sends v on ch if it can be accepted immediately, and reports
// whether it was.
//
// A metrics pipeline must never block its caller: when the buffer is full,
// the sample is dropped.
//
// Examples:
//
// 	Offer(ch, 1) => true when ch has room
func Offer(ch chan<- int, v int) bool {
	select {
	case ch <- v:
		return true
	default:
		return false
	}
}
```

## Walkthrough

With a buffer of 4 and ten offers, the first four are accepted and the rest take the default branch immediately.

## Pitfalls

- Checking `len(ch) < cap(ch)` first — the state can change between the check and the send.
- Dropping silently in a context where the caller needed to know; the boolean is the point.
