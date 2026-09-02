# The Generator Pattern

## Intuition

A generator turns "a batch I already have" into "a stream consumers can pull
from". Consumers `range` over the result, which only terminates if somebody
closes the channel — and that somebody is the producing goroutine.

## Approach

1. Make an unbuffered `chan string`.
2. Launch a goroutine that sends every key, with `defer close(out)`.
3. Return the channel right away so the caller can start draining while the goroutine feeds it.

## Solution

```go
func UploadFeed(keys []string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for _, k := range keys {
			out <- k
		}
	}()
	return out
}
```

## Walkthrough

For a batch of `a.jpg, b.jpg`: `UploadFeed` returns instantly. The consumer's
`range` blocks on the first receive; the goroutine's `out <- "a.jpg"` unblocks
it. That handshake repeats, then `close(out)` runs and the range ends.

## Pitfalls

- Sending on the channel before starting the goroutine — an unbuffered send with no receiver deadlocks the caller.
- Forgetting `close(out)`: the consumer's `range` blocks forever and the upload batch never completes.
- Letting the consumer close the channel — whoever sends must close, or a later send panics.
