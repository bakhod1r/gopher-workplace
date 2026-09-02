# A Semaphore from a Buffered Channel

## Intuition

A buffered channel of capacity N accepts N sends before blocking. That is
exactly a counting semaphore: sending takes a permit, receiving hands it back,
and the runtime does all the queuing for you.

## Approach

1. `sem := make(chan struct{}, limit)`.
2. Per recipient: `wg.Add(1)`, goroutine acquires, calls `send`, locks the mutex, adds, unlocks, releases.
3. `wg.Wait()` and return `total`.

## Solution

```go
import "sync"

func SendCampaign(recipients []string, limit int, send func(string) int) int {
	sem := make(chan struct{}, limit)

	var (
		wg    sync.WaitGroup
		mu    sync.Mutex
		total int
	)

	for _, addr := range recipients {
		wg.Add(1)
		go func(addr string) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			n := send(addr)

			mu.Lock()
			total += n
			mu.Unlock()
		}(addr)
	}

	wg.Wait()
	return total
}
```

## Walkthrough

With limit 2 and three recipients, the first two goroutines take both permits;
the third blocks on `sem <- struct{}{}` until one delivery finishes and
returns its permit. The mutex serialises the three `total +=` updates, so no
sum is lost.

## Pitfalls

- Adding to `total` without the mutex — `-race` flags it immediately and the total comes out wrong.
- Releasing the permit before `send` returns, which defeats the connection limit.
- Using an unbuffered channel: capacity 0 means the very first acquire blocks forever.
