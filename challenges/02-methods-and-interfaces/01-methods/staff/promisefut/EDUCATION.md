# Future / Promise

## Intuition

There are two ways to hand a value to a waiting goroutine: send it, or publish
it and announce that it is ready. A send serves exactly one receiver, so a
buffered channel makes a future that only one caller can read. Closing a channel
is a *broadcast* — it releases everyone, forever — which is what a future needs.

The value itself then travels through an ordinary field, and the close provides
the memory-ordering guarantee that makes reading that field safe.

## Approach

1. `Complete`: write the field, then close the signal channel.
2. `Get`: receive from the channel (blocking until it is closed), then read the
   field.

## Solution

```go
func (f *Future) Complete(val int) {
	f.val = val
	close(f.done)
}

func (f *Future) Get() int {
	<-f.done
	return f.val
}
```

## Walkthrough

`<-f.done` on an open, empty channel blocks. When `Complete` closes it, every
blocked receive returns the zero value immediately, and so does every later
receive — which is why `Get` is repeatable and why 50 waiters all wake from a
single `close`.

The ordering matters for correctness, not just style: Go's memory model
guarantees that a close happens-before a receive that observes it. Writing
`f.val` before the close therefore makes the write visible to every waiter. The
race detector confirms it; reversing the two lines makes `go test -race` fail.

## Pitfalls

- **Closing before assigning.** A waiter can observe the close and read `f.val`
  concurrently with the write — a genuine data race, and it can return 0.
- **Sending instead of closing.** `f.ch <- val` wakes one receiver; the other 49
  block forever, and a second `Get` deadlocks.
- **Completing twice.** `close` of an already-closed channel panics. A
  production future guards with `sync.Once` or a mutex-protected flag.
- **`Get` reading the field without waiting.** Returns 0 before completion.

## Why not `sync.WaitGroup`?

A `WaitGroup` also broadcasts, but it counts work rather than carrying a
result, and it cannot be polled without blocking. The closed-channel idiom
composes with `select`, which is what lets `Get` grow a `context.Context` or a
timeout arm later.
