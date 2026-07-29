# continue as a filter guard

## Intuition

`continue` skips to the next iteration; the guard must name the elements to SKIP, so a `continue` on the wanted elements silently drops them.

## Approach

1. To sum positives, skip the non-positives.
2. The bug `continue`s on `v > 0`, skipping exactly the ones it should add.
3. Invert to `if v <= 0 { continue }`.

## Solution

```go
func SumPositive(xs []int) int {
	sum := 0
	for _, v := range xs {
		if v <= 0 {
			continue
		}
		sum += v
	}
	return sum
}
```

## Walkthrough

The inverted guard skips positives and sums negatives, giving a wrong total. Skipping `v <= 0` accumulates only the positives → 9.

## Pitfalls

- `continue` should guard the unwanted case (`if !keep { continue }`).
- Continuing on the values you meant to process drops them.
