# Draining a Channel

## Intuition

Draining is how you unblock a producer whose values you no longer need.
Every queued send completes as you receive, and the loop ends at close.
Writing `for range ch` without binding a variable says "the value does not
matter".

## Approach

1. Start `discarded` at 0.
2. `for range attempts { discarded++ }`.
3. Return `discarded`.

## Solution

```go
func DrainQueue(attempts <-chan int) int {
	discarded := 0
	for range attempts {
		discarded++
	}
	return discarded
}
```

## Walkthrough

For three queued attempts then a close: three iterations bring `discarded`
to 3, then the closed channel ends the loop and the worker can exit.

## Pitfalls

- Draining a channel the dispatcher never closes blocks the shutdown forever.
- `for _ = range ch` is the un-idiomatic spelling; `for range ch` is preferred and `gofmt`-clean.
- Draining discards deliveries — make sure they are already recorded as failed.
