# Early Return Without Leaking

## Intuition

Returning early is only safe if the goroutines you leave behind can still
finish. Buffering the results channel to the number of senders guarantees
exactly that: every send succeeds immediately, whether or not anyone is left
to read it.

## Approach

1. Return false for an empty replica set.
2. `acks := make(chan bool, len(replicas))` and one goroutine per replica sending its result.
3. Loop `len(replicas)` times receiving; return true on the first `true`, otherwise false.

## Solution

```go
func FirstReplicaAck(replicas []string, write func(string) bool) bool {
	if len(replicas) == 0 {
		return false
	}

	acks := make(chan bool, len(replicas))
	for _, replica := range replicas {
		go func(replica string) {
			acks <- write(replica)
		}(replica)
	}

	for range replicas {
		if <-acks {
			return true
		}
	}
	return false
}
```

## Walkthrough

With `ok-1` and `bad-2`, suppose the ack from `ok-1` is read first: the
function returns true straight away. The goroutine for `bad-2` still has a
free buffer slot, so its send completes and the goroutine exits — nothing is
leaked, and the buffered value is simply garbage collected.

## Pitfalls

- Using an unbuffered channel: after the early return the remaining writers block on send and leak, one per request.
- Ranging over `acks` instead of counting — nobody closes it, so the range would block after the last result.
- Waiting for every replica before returning, which makes the write as slow as the slowest replica.
