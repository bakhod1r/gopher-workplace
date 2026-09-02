# Filtering in a Pipeline

## Intuition

A filter is a stage whose *send* is conditional. The *close* is not: the
alerting consumer relies on the close to know the batch is finished, whether
or not a single error record got through.

## Approach

1. Create the output channel and start a goroutine with `defer close(out)`.
2. Range over `records`.
3. Send `rec` only when `isError(rec)` returns true.

## Solution

```go
func ErrorFilter(records <-chan string, isError func(string) bool) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for rec := range records {
			if isError(rec) {
				out <- rec
			}
		}
	}()
	return out
}
```

## Walkthrough

With `ERR disk, INFO ok, ERR io`: the first record passes, the second is
dropped, the third passes. The input closes, the range ends, `close(out)`
runs — the alerting stage sees exactly two records and then a clean shutdown.

## Pitfalls

- Returning early when nothing matches, leaving the output channel open forever.
- Putting `close(out)` inside the loop: it closes after the first match and panics on the next send.
- Calling `isError` twice per record when it may be expensive — evaluate once into a variable if so.
