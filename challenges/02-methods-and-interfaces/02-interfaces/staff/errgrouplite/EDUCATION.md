# Error Group

## Intuition

This is `errgroup` in miniature. Two invariants carry it: the first error is recorded exactly once (so the cancel channel closes once), and the semaphore is acquired before spawning, which bounds goroutines as well as concurrency.

## Approach

1. `Go` calls `wg.Add(1)`, then blocks on the semaphore before starting the goroutine.
2. The goroutine defers both `wg.Done` and the semaphore release.
3. On error, `errOnce.Do` records it and closes `cancel`.
4. `Wait` waits on the WaitGroup and returns the recorded error.

## Solution

```go
func (g *Group) Go(t Task) {
	g.wg.Add(1)
	g.sem <- struct{}{}

	go func() {
		defer g.wg.Done()
		defer func() { <-g.sem }()

		if err := t.Run(g.cancel); err != nil {
			g.errOnce.Do(func() {
				g.err = err
				close(g.cancel)
			})
		}
	}()
}

func (g *Group) Wait() error {
	g.wg.Wait()
	return g.err
}
```

## Walkthrough

Reading `g.err` after `wg.Wait()` is safe without a lock: the WaitGroup establishes happens-before between the writing goroutine's `Done` and the waiter's return.

## Pitfalls

- `close(g.cancel)` without `sync.Once` — the second failure panics.
- Acquiring the semaphore inside the goroutine, which bounds concurrency but not the number of goroutines.
- Reading `g.err` before `Wait` returns, which is a data race.
