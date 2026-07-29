# Multiple closures over shared state

## Intuition

Closures defined in one scope share its variables, letting a factory return several functions that operate on the same hidden state.

## Approach

1. Capture a shared `sum`.
2. Return one closure that adds and one that reads.

## Solution

```go
func NewTracker() (add func(int), total func() int) {
	sum := 0
	add = func(n int) { sum += n }
	total = func() int { return sum }
	return
}
```

## Walkthrough

`add(3)` and `add(4)` both mutate the shared `sum`, so `total()` returns 7.

## Pitfalls

- Both returned functions capture the SAME `sum` by reference.
- The state is unreachable except through the returned closures.
