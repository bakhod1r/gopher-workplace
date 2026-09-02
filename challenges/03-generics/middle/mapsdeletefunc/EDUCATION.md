# Delete Map Entries

## Intuition

Deleting during a map range is explicitly legal in Go, which is what lets the stdlib do this in one pass without a second collection.

## Approach

1. Clone the map, normalising nil.
2. Delete the matching entries.
3. Return the clone.

## Solution

```go
func Prune[K comparable, V any](m map[K]V, drop func(K, V) bool) map[K]V {
	out := maps.Clone(m)
	if out == nil {
		out = make(map[K]V)
	}
	maps.DeleteFunc(out, drop)
	return out
}
```

## Walkthrough

`Prune({a:1,b:2}, dropEven)` removes `b` from the copy while the original keeps both entries.

## Pitfalls

- Assigning the result of `maps.DeleteFunc`, which returns nothing.
- Deleting from the caller's map.
- Collecting keys first and deleting afterwards — unnecessary for maps.
