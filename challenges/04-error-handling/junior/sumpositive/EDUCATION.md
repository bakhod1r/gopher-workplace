# Sum Of Positives

## Intuition

When a computation aborts, its intermediate state is meaningless. Returning the zero value alongside the error makes it impossible for a caller to use half a sum.

## Approach

1. Accumulate in a local total.
2. Return `0, ErrNegativeValue` on the first negative entry.
3. Return `total, nil` after the loop.

## Solution

```go
total := 0
for _, n := range nums {
	if n < 0 {
		return 0, ErrNegativeValue
	}
	total += n
}
return total, nil
```

## Walkthrough

For `[]int{1, 2, -1}` the total reaches 3, then the negative entry aborts and `0, ErrNegativeValue` is returned — the 3 is thrown away.

## Pitfalls

- Returning the partial total with the error.
- Rejecting zero by testing `n <= 0`.
- Validating in a first loop and summing in a second when one pass suffices.
