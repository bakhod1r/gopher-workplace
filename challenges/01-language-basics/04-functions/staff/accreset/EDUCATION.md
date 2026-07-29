# Accumulator threading in recursion

## Intuition

Tail-recursive folds pass the running result forward each call; replacing the accumulator instead of adding to it loses all prior work.

## Approach

1. The accumulator must carry forward across recursive calls.
2. The bug passes `xs[0]` (resetting acc); pass `acc + xs[0]`.

## Solution

```go
func sumAcc(xs []int, acc int) int {
	if len(xs) == 0 {
		return acc
	}
	return sumAcc(xs[1:], acc+xs[0])
}

func Sum(xs []int) int {
	return sumAcc(xs, 0)
}
```

## Walkthrough

Passing only `xs[0]` discards the running total, so the result equals the last-but-one element logic, not the sum. Adding into `acc` accumulates correctly to 10.

## Pitfalls

- The recursive call must combine `acc` with the current element.
- Passing only `xs[0]` restarts the fold.
