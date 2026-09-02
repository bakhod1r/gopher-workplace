# The Or-Done Pattern

## Intuition

A goroutine leaks when it is blocked on an operation nobody will ever complete.
Or-done removes both such operations: the receive and the send each become a
`select` that `done` can win, so cancellation is always reachable.

## Approach

1. Create `out` and start a goroutine with `defer close(out)`.
2. Outer `select`: receive a tick (returning when `ok` is false) or observe `done`.
3. Inner `select`: send the tick on `out`, or observe `done` and return.

## Solution

```go
func LiveTicks(done <-chan struct{}, ticks <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			select {
			case tick, ok := <-ticks:
				if !ok {
					return
				}
				select {
				case out <- tick:
				case <-done:
					return
				}
			case <-done:
				return
			}
		}
	}()
	return out
}
```

## Walkthrough

With `done` open and ticks 1, 2, 3: each loop receives a tick and sends it on,
then the closed feed makes `ok` false and the goroutine returns, closing `out`.
If the dashboard closes `done` mid-stream, whichever select is blocked takes
the `done` case instead and the goroutine exits rather than leaking.

## Pitfalls

- Guarding only the receive: the goroutine then blocks forever on `out <- tick` once the consumer stops reading.
- Returning without closing `out`, which hangs any consumer still ranging over it.
- Assuming `done` and `ticks` cannot both be ready — `select` picks at random, which is fine here because both paths are correct.
