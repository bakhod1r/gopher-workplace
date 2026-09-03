# A Buffered Channel as a Ring Buffer

## Intuition

A buffered channel is a FIFO queue with a hard capacity, which is exactly what
a log tail needs. The only piece missing is the eviction policy: by default a
full channel blocks the sender, and here we want it to *evict* instead.

`select`/`default` detects fullness without blocking, and the recovery is a
receive followed by the send. That receive is what makes it drop-oldest — the
value that comes out is the line that has been in the ring the longest.

## Approach

1. `keep <= 0`: drain `lines` and return `[]string{}`.
2. `ring := make(chan string, keep)`.
3. For each line: try `ring <- line`; on `default`, `<-ring` then `ring <- line`.
4. `close(ring)` once the stream ends.
5. `range ring` into a slice and return it.

## Solution

```go
// TailBuildLog follows a build's log stream and returns only the last keep
// lines, in stream order. The log is unbounded and the CI page only renders a
// tail, so memory is capped by keep: a buffered channel of capacity keep acts
// as the ring, and the oldest line is discarded whenever a new one arrives at
// a full ring.
//
// Examples:
//
//	TailBuildLog(chan a,b,c, 2) => [b c]
//	TailBuildLog(chan a,b, 5)   => [a b]
//	TailBuildLog(chan a,b, 0)   => []
func TailBuildLog(lines <-chan string, keep int) []string {
	tail := []string{}
	if keep <= 0 {
		for range lines {
		}
		return tail
	}

	ring := make(chan string, keep)
	for line := range lines {
		select {
		case ring <- line:
		default:
			<-ring
			ring <- line
		}
	}
	close(ring)

	for line := range ring {
		tail = append(tail, line)
	}
	return tail
}
```

## Walkthrough

- `keep = 2`, lines `a b c`: `a` and `b` fill the ring. For `c` the send arm
  cannot proceed, so `default` receives `a` (the oldest) and then `c` fits.
  Draining gives `[b c]`.
- `keep = 5`, lines `a b`: the ring never fills, `default` never runs, and the
  drain returns everything in order.
- `keep = 1`, four lines: every line after the first evicts its predecessor, so
  only `d` survives.
- An empty log closes the ring immediately and the drain loop yields `[]`.
- Only one goroutine touches `ring`, so the non-blocking send and the recovery
  receive cannot race with anyone.

## Pitfalls

- Returning early on `keep <= 0` without draining leaves the log writer blocked
  forever on a stream with no reader.
- Doing the recovery in the other order — send first, then receive — is what a
  blocking `ring <- line` would do: it deadlocks a single-goroutine tailer.
- Forgetting `close(ring)` makes the final `range ring` block after the last
  buffered line.
- This single-goroutine trick is safe precisely because nobody else sends to
  `ring`; with a second sender, "full" could become "not full" between the
  `default` and the retry.
