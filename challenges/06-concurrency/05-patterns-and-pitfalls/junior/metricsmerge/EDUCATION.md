# Fan-In: Merging Many Channels

## Intuition

A channel with many senders needs exactly one closer. The trick is a small
goroutine whose only job is "wait for every sender, then close" — it runs
concurrently with the consumer, so nothing deadlocks.

## Approach

1. Create the `merged` channel.
2. For each stream, `wg.Add(1)` and start a goroutine forwarding its samples into `merged`.
3. Start the closer goroutine, drain `merged` in the caller, then `sort.Ints`.

## Solution

```go
import (
	"sort"
	"sync"
)

func MergeMetrics(streams ...<-chan int) []int {
	merged := make(chan int)

	var wg sync.WaitGroup
	for _, s := range streams {
		wg.Add(1)
		go func(s <-chan int) {
			defer wg.Done()
			for sample := range s {
				merged <- sample
			}
		}(s)
	}

	go func() {
		wg.Wait()
		close(merged)
	}()

	var out []int
	for sample := range merged {
		out = append(out, sample)
	}
	sort.Ints(out)
	return out
}
```

## Walkthrough

With node streams `{3, 1}` and `{2}`, the two forwarders interleave into
`merged` however the scheduler decides. Once both ranges end, `wg.Wait()`
returns, `merged` closes, the caller's range ends, and sorting yields
`1, 2, 3`.

## Pitfalls

- Calling `wg.Wait()` in the calling goroutine before draining `merged`: forwarders block on send, so `Wait` never returns.
- Closing `merged` inside a forwarder — the other forwarders then panic sending on a closed channel.
- Forgetting to close at all, which hangs the aggregator's `range` forever.
