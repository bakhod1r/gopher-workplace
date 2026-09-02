# Generic Type Alias

## Intuition

Aliases add vocabulary without adding types. That is the right choice when you want readability but still need to pass the value to anything expecting the underlying map.

## Approach

1. `NewIndex`: return `make(Index[K])`.
2. `Mark`: assign `struct{}{}`.
3. `Marked`: comma-ok lookup.

## Solution

```go
func NewIndex[K comparable]() Index[K] {
	return make(Index[K])
}

func Mark[K comparable](ix Index[K], k K) {
	ix[k] = struct{}{}
}

func Marked[K comparable](ix Index[K], k K) bool {
	_, ok := ix[k]
	return ok
}
```

## Walkthrough

`Mark(map[string]struct{}{}, "a")` compiles because the alias is the map type, not a wrapper around it.

## Pitfalls

- Writing `type Index[K comparable] map[K]struct{}` (no `=`), which defines a distinct type.
- Trying to declare methods on the alias.
- Forgetting `make`, so the first `Mark` panics on a nil map.
