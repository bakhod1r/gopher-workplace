# Send-Only Channels and Ownership

## Intuition

A `chan<- int` parameter is a promise: this function will only put values
in. It may also `close` — closing is a send-side operation. The rule
"the sender closes" exists because a receiver closing would make other
senders panic.

## Approach

1. `range` the batch and send each id on `out`.
2. After the loop, `close(out)` once.

## Solution

```go
func ShipOrderIDs(out chan<- int, ids []int) {
	for _, id := range ids {
		out <- id
	}
	close(out)
}
```

## Walkthrough

With `[]int{101, 102}`: send `101`, send `102`, then `close(out)`. The
Kafka producer ranging over the same channel receives both ids and stops.

## Pitfalls

- `close(out)` inside the loop closes after the first id — the second send panics.
- Closing twice panics with "close of closed channel".
- You cannot receive from a `chan<- int`; the compiler rejects `<-out`.
