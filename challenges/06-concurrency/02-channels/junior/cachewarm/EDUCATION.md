# The Comma-Ok Receive

## Intuition

A bare `<-ch` cannot distinguish a sent `0` from a closed channel — both
look like `0`. The two-value form `v, ok := <-ch` adds that bit: `ok` is
`true` when the value came from a real send, `false` when the channel is
closed and empty.

## Approach

1. Use the two-value receive form.
2. Return both results unchanged.

## Solution

```go
func NextKey(feed <-chan int) (int, bool) {
	id, ok := <-feed
	return id, ok
}
```

## Walkthrough

On a feed holding key `5`, the receive takes `5` and sets `ok = true`. On a
closed empty feed the receive returns immediately with `0, false` — and keeps
doing so on every later tick.

## Pitfalls

- `id := <-feed` alone loses the closed/open distinction.
- A closed channel never blocks — a `for` loop over bare receives spins forever.
- `ok` is `true` for ids still buffered in a *closed* feed; it flips to `false` only once the buffer is empty.
