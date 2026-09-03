# Merge Channels Without Leaving Goroutines Behind

## Intuition

A goroutine blocked on a send to a channel nobody reads never returns and never gets collected. Every send in a fan-in has to be paired with an escape route, and the output can only be closed once all of them are gone.

## Approach

1. Start one forwarder per input, tracked by a `WaitGroup`.
2. Forward with `select` over the send and `<-done`.
3. In a separate goroutine, `Wait` then `close(out)`.

## Solution

```go
import "sync"

// Merge returns a channel carrying every value from ins, closed once all
// inputs are drained or done is closed.
//
// Every goroutine Merge starts must exit: an abandoned consumer must not
// leave forwarders blocked on a send forever.
//
// Examples:
//
// 	Merge(done, a, b) => a channel with everything from a and b
func Merge(done <-chan struct{}, ins ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(ins))
	for _, in := range ins {
		go func(in <-chan int) {
			defer wg.Done()
			for v := range in {
				select {
				case out <- v:
				case <-done:
					return
				}
			}
		}(in)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
```

## Walkthrough

With four inputs and a consumer that reads once and leaves, closing `done` lets all four forwarders return from their `select`; the closer's `Wait` then returns and `out` is closed.

## Pitfalls

- Calling `wg.Wait()` in `Merge` itself, which blocks until the output is fully consumed.
- Closing `out` in the forwarders, which panics on the second close.
- A bare `out <- v` with no `select`, which is exactly the leak.
