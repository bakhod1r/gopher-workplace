# Stateful closures for caching

## Intuition

A closure holding a map memoises across calls; the state is private and survives between invocations.

## Approach

1. Capture a `cache` map.
2. Return a closure that checks the cache before calling `f`, storing new results.

## Solution

```go
func Memoize(f func(int) int) func(int) int {
	cache := map[int]int{}
	return func(x int) int {
		if v, ok := cache[x]; ok {
			return v
		}
		v := f(x)
		cache[x] = v
		return v
	}
}
```

## Walkthrough

The first `m(2)` computes and caches; the second `m(2)` returns the cached value without calling `f` again.

## Pitfalls

- Use comma-ok to distinguish a stored 0 from a missing key.
- Only correct for pure functions; side-effecting `f` breaks caching semantics.
