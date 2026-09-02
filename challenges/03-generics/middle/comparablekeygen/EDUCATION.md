# comparable Versus any

## Intuition

The two functions do the same work and fail in different eras: one at compile time, one in production. That difference is the argument for constraints.

## Approach

1. Take the smaller length.
2. Allocate the map.
3. Assign each pair.

## Solution

```go
func Index[K comparable, V any](keys []K, vals []V) map[K]V {
	n := len(keys)
	if len(vals) < n {
		n = len(vals)
	}
	out := make(map[K]V, n)
	for i := 0; i < n; i++ {
		out[keys[i]] = vals[i]
	}
	return out
}

func IndexAny(keys []any, vals []any) map[any]any {
	n := len(keys)
	if len(vals) < n {
		n = len(vals)
	}
	out := make(map[any]any, n)
	for i := 0; i < n; i++ {
		out[keys[i]] = vals[i]
	}
	return out
}
```

## Walkthrough

`Index` with a `[]int` key argument fails to instantiate; `IndexAny` accepts it and panics on the first insert.

## Pitfalls

- Declaring `K any` and reintroducing the run-time panic.
- Ranging over `keys` and indexing `vals`, which panics when `vals` is shorter.
- Assuming the result size equals the number of pairs.
