# Chaining Pipeline Stages

## Intuition

A pipeline is a chain of `range`-then-`close` stages. Termination flows
downstream: closing the source ends stage 2's loop, which closes its own
output, which ends the collector's loop. One close at the head shuts the
whole chain down cleanly.

## Approach

1. Start a goroutine that sends every unit count on `counts` and closes it.
2. Start a second goroutine that ranges `counts`, sends `n * 2` on `priced`, and closes `priced`.
3. In the collector, range `priced` and append `cents + 1`.
4. Return the slice.

## Solution

```go
func LineTotals(units []int) []int {
	counts := make(chan int)
	go func() {
		for _, n := range units {
			counts <- n
		}
		close(counts)
	}()

	priced := make(chan int)
	go func() {
		for n := range counts {
			priced <- n * 2
		}
		close(priced)
	}()

	out := []int{}
	for cents := range priced {
		out = append(out, cents+1)
	}
	return out
}
```

## Walkthrough

For `[1, 2]`: `counts` carries 1, 2; `priced` carries 2, 4; the collector
appends 3 and 5 — `[3 5]`. For an empty order, `counts` closes immediately,
`priced` closes right after, and the result is `[]`.

## Pitfalls

- A stage that closes its *input* instead of its output panics the upstream sender.
- Forgetting a stage's `close` hangs everything downstream.
- Unbuffered channels are fine here only because every stage runs in its own goroutine.
