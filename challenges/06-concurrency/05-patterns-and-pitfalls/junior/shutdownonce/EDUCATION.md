# Closing a Channel Exactly Once

## Intuition

Channel close is a broadcast, which is why several components want to call it —
and it is also a one-shot operation that panics on repeat. `sync.Once` bridges
the two: the first caller closes, the rest return having done nothing.

## Approach

1. Declare `var once sync.Once` and a `WaitGroup` in the enclosing function.
2. Start `closers` goroutines, each calling `once.Do(func() { close(quit) })`.
3. `wg.Wait()` so the function does not return before every requester is done.

## Solution

```go
import "sync"

func ShutdownOnce(quit chan struct{}, closers int) {
	var (
		once sync.Once
		wg   sync.WaitGroup
	)

	for i := 0; i < closers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(func() { close(quit) })
		}()
	}

	wg.Wait()
}
```

## Walkthrough

With ten closers, all ten call `once.Do`. One of them wins and runs the
closure; the other nine block until that close has completed and then carry on
without calling it. `quit` ends up closed once, and every waiter on it wakes.

## Pitfalls

- Creating the `Once` inside the goroutine: each goroutine gets its own, so every one of them closes and all but the first panic.
- Guarding with `if !closed { close(quit) }` and a plain bool — the check and the close are not atomic, so two goroutines can both pass the check.
- Returning before `wg.Wait()`, which can report the channel as still open.
