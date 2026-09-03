# Groups That Keep Only The Last

## Intuition

Assigning a fresh one-element slice drops the existing bucket, so only the last element of each group survives.

## Approach

1. Make the map.
2. Compute the key.
3. Append the element to that key's bucket.

## Solution

```go
func GroupBy[T any, K comparable](s []T, key func(T) K) map[K][]T {
	out := make(map[K][]T)
	for _, v := range s {
		k := key(v)
		out[k] = append(out[k], v)
	}
	return out
}
```

## Walkthrough

For `[1 3]` the bucket for odd is set to `[1]`, then replaced by `[3]`.

## Pitfalls

- Pre-initialising the bucket with `if _, ok := ...` — harmless but unnecessary.
- Assuming map iteration will report the groups in a stable order.
