# Broadcasting with a Closed Channel

## Intuition

Sending a value wakes one receiver. Closing a channel wakes all of them, and
keeps waking any that arrive later. That asymmetry is why signals — start,
stop, cancel — are always modelled as a close, never as a send.

## Approach

1. Preallocate `counts` with `make([]int, len(shards))`.
2. Per shard: `wg.Add(1)` and start a goroutine that blocks on `<-ready`, then writes `counts[i] = warm(shard)`.
3. `wg.Wait()`, `sort.Ints(counts)`, return.

## Solution

```go
import (
	"sort"
	"sync"
)

func WarmShards(ready <-chan struct{}, shards []string, warm func(string) int) []int {
	counts := make([]int, len(shards))

	var wg sync.WaitGroup
	for i, shard := range shards {
		wg.Add(1)
		go func(i int, shard string) {
			defer wg.Done()
			<-ready
			counts[i] = warm(shard)
		}(i, shard)
	}

	wg.Wait()
	sort.Ints(counts)
	return counts
}
```

## Walkthrough

All warmers are created and immediately park on `<-ready`. When the deploy
closes the channel, every parked receive returns at once and the warmers run
concurrently. `wg.Wait()` returns after the last one, and the counts are
sorted before being handed back.

## Pitfalls

- Sending `ready <- struct{}{}` once instead of closing: only a single warmer would ever start.
- Reading `counts` before `wg.Wait()` returns — the writes are not guaranteed visible yet.
- Appending to a shared slice instead of writing distinct indexes, which is a data race.
