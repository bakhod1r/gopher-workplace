# A Stage That Stops When Told

## Intuition

Cancellation has to reach every place a goroutine can wait. Covering the receive and forgetting the send is the classic half-fix: the workers wake up, take a value, and block forever handing it on.

## Approach

1. Start `workers` goroutines tracked by a `WaitGroup`.
2. Each loops: `select` to receive or return on `done`; return when the input is closed.
3. `select` to send or return on `done`.
4. A separate goroutine `Wait`s and closes the output.

## Solution

```go
import "sync"

// Stage returns a channel carrying each input doubled, computed by
// workers goroutines, and closed once the input drains or done is closed.
//
// Every goroutine must exit on done, whether it is blocked receiving or
// blocked sending.
//
// Examples:
//
// 	Stage(done, in, 4) => a channel of doubled values
func Stage(done <-chan struct{}, in <-chan int, workers int) <-chan int {
	if workers < 1 {
		workers = 1
	}
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(workers)
	for w := 0; w < workers; w++ {
		go func() {
			defer wg.Done()
			for {
				var v int
				var ok bool
				select {
				case v, ok = <-in:
					if !ok {
						return
					}
				case <-done:
					return
				}
				select {
				case out <- v * 2:
				case <-done:
					return
				}
			}
		}()
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
```

## Walkthrough

With four workers and an abandoned consumer, closing `done` releases whichever `select` each worker is sitting in; the closer's `Wait` then returns and the output is closed.

## Pitfalls

- `for v := range in` with a plain send — the range is cancellable only by closing `in`, and the send is not cancellable at all.
- Closing `out` from a worker, which panics when the second one does it.
- Calling `wg.Wait()` inside `Stage`, which blocks the caller until the pipeline is fully consumed.
