# The Self-Constraint Called Backwards

## Intuition

Because the constraint is written over `T` itself, the receiver and the argument have the same static type, so swapping them is invisible to the compiler. The loop asks "is the incumbent smaller than the candidate?" and then replaces the incumbent when the answer is yes — which selects the maximum.

## Approach

1. Return early for an empty slice.
2. Take the first element as the incumbent.
3. Replace it whenever the *candidate* is less than the incumbent.

## Solution

```go
func MinOf[T Lesser[T]](xs []T) (T, bool) {
	if len(xs) == 0 {
		var zero T
		return zero, false
	}
	best := xs[0]
	for _, v := range xs[1:] {
		if v.Less(best) {
			best = v
		}
	}
	return best, true
}
```

## Walkthrough

For `[{1,5} {1,2} {2,0}]` the loop keeps `{1,5}` against `{1,2}` and then adopts `{2,0}`, returning the largest version.

## Pitfalls

- Adding a `Greater` method to compensate instead of fixing the call.
- Assuming a `Less` that is not a strict weak ordering will still yield a stable minimum.
