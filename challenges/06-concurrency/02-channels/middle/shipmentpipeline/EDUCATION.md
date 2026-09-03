# Shipment Label Pipeline

## Intuition

A pipeline is a chain of `for range in { ... out <- x }` loops. The only contract is that each stage closes the channel it writes to and never the one it reads from. Then closing the source unravels the whole chain.

## Approach

1. Start a source goroutine sending every order, `defer close(source)`.
2. Stage one: range `source`, forward orders that pass `keep`, `defer close(kept)`.
3. Stage two: range `kept`, send `render(o)`, `defer close(labels)`.
4. Range `labels` on the caller and append to the result slice.

## Solution

```go
func Labels(orders []string, keep func(order string) bool, render func(order string) string) []string {
	source := make(chan string)
	go func() {
		defer close(source)
		for _, o := range orders {
			source <- o
		}
	}()

	kept := make(chan string)
	go func() {
		defer close(kept)
		for o := range source {
			if keep(o) {
				kept <- o
			}
		}
	}()

	labels := make(chan string)
	go func() {
		defer close(labels)
		for o := range kept {
			labels <- render(o)
		}
	}()

	out := []string{}
	for l := range labels {
		out = append(out, l)
	}
	return out
}
```

## Walkthrough

For `[o1, bad, o2]`: the source sends `o1`; stage one keeps it and passes it on; stage two renders `LABEL-o1`. `bad` reaches stage one, fails `keep`, and is dropped there — stage two never sees it. When the source closes, each stage's `range` ends and closes its own output, so the caller's loop finishes with two labels.

## Pitfalls

- Closing an input channel from a downstream stage — only the writer may close.
- Returning `nil` instead of an empty slice when everything is dropped (the test compares with `[]string{}`).
- Adding parallelism inside a stage, which silently scrambles the output order.
