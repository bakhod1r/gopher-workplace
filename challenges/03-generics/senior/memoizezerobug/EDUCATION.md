# The Cache That Never Fills

## Intuition

The lookup is there, but nothing is ever stored, so the map stays empty and every call falls through to `fn`.

## Approach

1. Create the cache in the outer function.
2. Return the hit if present.
3. Otherwise compute, store, and return.

## Solution

```go
func Memoize[K comparable, V any](fn func(K) V) func(K) V {
	cache := make(map[K]V)
	return func(k K) V {
		if v, ok := cache[k]; ok {
			return v
		}
		v := fn(k)
		cache[k] = v
		return v
	}
}
```

## Walkthrough

`f(1)` misses, calls `fn`, and returns — leaving the cache empty, so `f(1)` misses again.

## Pitfalls

- Recreating the map inside the returned closure.
- Assuming the closure is safe for concurrent use — it is not.
