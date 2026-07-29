# Binary search boundary updates

## Intuition

After testing `mid`, the half that excludes it is `[mid+1, hi]` or `[lo, mid-1]`; reusing `mid` in a bound breaks the shrink invariant and can loop.

## Approach

1. When the target is greater than `xs[mid]`, search the right half **above** mid.
2. The bug sets `lo = mid` (can loop forever); use `lo = mid + 1`.

## Solution

```go
func IndexOf(xs []int, target int) int {
	lo, hi := 0, len(xs)-1
	for lo <= hi {
		mid := (lo + hi) / 2
		switch {
		case xs[mid] == target:
			return mid
		case xs[mid] < target:
			lo = mid + 1
		default:
			hi = mid - 1
		}
	}
	return -1
}
```

## Walkthrough

`lo = mid` without advancing can stall when `lo == mid`. `lo = mid + 1` strictly shrinks the range, so the search terminates and finds index 3.

## Pitfalls

- `xs[mid] < target` ⇒ `lo = mid + 1`, never `lo = mid`.
- Every iteration must reduce `hi - lo`.
