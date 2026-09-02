# Non-Blocking Send With default

## Intuition

`select` cases can be sends as well as receives. With a `default`, a send
that would block is skipped entirely — the classic "drop it rather than
stall" policy for telemetry and best-effort queues. On an unbuffered
channel with no waiting collector, the send is never ready, so `default`
always wins.

## Approach

1. Write a `select` with the case `buffer <- sample` and a `default`.
2. Return `true` from the send case.
3. Return `false` from `default`.

## Solution

```go
func TryRecord(buffer chan<- int, sample int) bool {
	select {
	case buffer <- sample:
		return true
	default:
		return false
	}
}
```

## Walkthrough

On an empty buffer of capacity 1 the send is ready and returns `true`; a
second call finds `len == cap`, so `default` fires and the sample is
dropped. On an unbuffered buffer with no collector, `default` fires
immediately.

## Pitfalls

- A bare `buffer <- sample` blocks when the buffer is full — that is what `default` avoids.
- Sending on a closed channel panics even inside a `select`; `default` does not protect you.
- Silently dropping samples needs to be a deliberate decision, and usually a counted one.
