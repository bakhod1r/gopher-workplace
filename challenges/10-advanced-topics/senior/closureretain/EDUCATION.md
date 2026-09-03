# The Callback That Held The Whole Batch

## Intuition

Whatever a closure mentions, it keeps. Deferring the computation looks lazy and cheap, but it converts a temporary batch into state with the callback's lifetime.

## Approach

1. Compute the total before constructing the closure.
2. Return a closure that captures only the total.

## Solution

```go
// Record is one ingested item.
type Record struct {
	Size int
	Pad  [256]byte
}

// Summarize returns a function reporting the batch's total size.
//
// The returned function outlives the batch, so it must capture the answer
// rather than the data: a closure over the slice keeps every record alive
// for as long as the callback exists.
//
// Examples:
//
// 	f := Summarize(batch); f() => the total size
func Summarize(batch []Record) func() int {
	total := 0
	for _, r := range batch {
		total += r.Size
	}
	return func() int { return total }
}
```

## Walkthrough

8192 records of about 264 bytes is roughly 2 MiB. The lazy closure pins all of it; the eager version captures one int and the batch is collected as soon as `Summarize` returns.

## Pitfalls

- Copying the batch into the closure — same retention, plus a copy.
- Assuming laziness is free; it moves work later and lifetime longer.
