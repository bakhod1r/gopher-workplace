# Transactional Mail Semaphore

## Intuition

A semaphore is a bowl with `limit` tokens. Take one before you start work, drop it back when you finish. If the bowl is empty you block at the counter — which is exactly what a full buffered channel does to a sender.

## Approach

1. Return `[]string{}` for an empty input; otherwise allocate `out` with `len(messages)`.
2. Normalise a non-positive `limit` to the message count (unlimited).
3. Create `sem` with capacity `limit`.
4. For each message: `wg.Add(1)`, acquire a slot, then launch a goroutine that defers `wg.Done` and the slot release, and writes `out[i]`.
5. `wg.Wait()` and return `out`.

## Solution

```go
import "sync"

func SendAll(messages []string, limit int, send func(msg string) string) []string {
	out := make([]string, len(messages))
	if len(messages) == 0 {
		return []string{}
	}

	if limit < 1 {
		limit = len(messages)
	}
	sem := make(chan struct{}, limit)

	var wg sync.WaitGroup
	for i, m := range messages {
		wg.Add(1)
		sem <- struct{}{}
		go func(i int, m string) {
			defer wg.Done()
			defer func() { <-sem }()
			out[i] = send(m)
		}(i, m)
	}
	wg.Wait()
	return out
}
```

## Walkthrough

With 200 messages and limit 5, the loop takes a token, spawns a sender, and repeats. On the sixth iteration the buffer is full, so the loop blocks until a running send finishes and returns its token. Peak concurrency therefore never passes 5, and each goroutine writes only its own index, so results stay in input order.

## Pitfalls

- Acquiring the slot *inside* the goroutine: all N goroutines are created at once, and a huge input can exhaust memory even though the provider limit is honoured.
- Releasing the slot only on the success path — one panic and the semaphore leaks a token permanently.
- Appending to a shared slice instead of writing `out[i]`: that is both a data race and a scrambled order.
