# Labeled break and continue

## Intuition

Labels let `break`/`continue` target an enclosing loop, avoiding sentinel flags in nested iteration.

## Approach

1. Label the outer loop.
2. On the first matching pair, set the returns and `break Outer`.

## Solution

```go
func FindPairSum(xs []int, target int) (i, j int, ok bool) {
Outer:
	for i = 0; i < len(xs); i++ {
		for j = i + 1; j < len(xs); j++ {
			if xs[i]+xs[j] == target {
				ok = true
				break Outer
			}
		}
	}
	return
}
```

## Walkthrough

For target 7 in `[1 2 3 4]`, indices 2 and 3 sum to 7; the labeled break exits both loops at once.

## Pitfalls

- Only the FIRST matching pair should be returned; break immediately.
- Without the label, `break` exits just the inner loop.
