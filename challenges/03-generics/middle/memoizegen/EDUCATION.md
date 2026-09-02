# Memoize

## Intuition

Testing `cache[k] != zero` would recompute every cached zero value — and would need `V comparable`, needlessly narrowing the API.

## Approach

1. Allocate the cache before returning the closure.
2. Return the cached value when present.
3. Otherwise call `f`, store, and return.

## Solution

```go
func Memoize[K comparable, V any](f func(K) V) func(K) V {
	cache := make(map[K]V)
	return func(k K) V {
		if v, ok := cache[k]; ok {
			return v
		}
		v := f(k)
		cache[k] = v
		return v
	}
}
```

## Walkthrough

The second `m(1)` finds the key present and returns without calling `f` again.

## Pitfalls

- Allocating the cache inside the closure, so nothing is ever retained.
- Comparing against the zero value instead of using `ok`.
- Claiming concurrency safety without a mutex.
