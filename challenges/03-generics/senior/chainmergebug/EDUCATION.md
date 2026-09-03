# The Override That Never Applies

## Intuition

The presence check turns "later wins" into "first wins", so every override layer is discarded.

## Approach

1. Allocate the result.
2. Walk the maps in order.
3. Assign each key unconditionally.

## Solution

```go
func Merge[K comparable, V any](ms ...map[K]V) map[K]V {
	out := make(map[K]V)
	for _, m := range ms {
		for k, v := range m {
			out[k] = v
		}
	}
	return out
}
```

## Walkthrough

`Merge({a:1}, {a:2})` finds `a` already set from the first map and refuses to write `2`.

## Pitfalls

- Reversing the argument order to compensate instead of fixing the logic.
- Returning `ms[0]` when there is exactly one map — that aliases the caller's map.
