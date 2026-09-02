# Count Failures

## Intuition

Counting is a full traversal, not a search. Every element is inspected and the running total survives the loop.

## Approach

1. Declare a counter before the loop.
2. Increment it for every non-nil entry.
3. Return the counter after the loop.

## Solution

```go
n := 0
for _, err := range errs {
	if err != nil {
		n++
	}
}
return n
```

## Walkthrough

For `[]error{nil, ErrX, ErrX}` the counter stays 0, becomes 1, then 2, and 2 is returned.

## Pitfalls

- Returning inside the loop — that stops at the first failure.
- Declaring the counter inside the loop, resetting it each pass.
- Counting the length of the slice instead of its non-nil entries.
