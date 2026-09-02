# Contains By Predicate

## Intuition

Choosing `ContainsFunc` over `IndexFunc(...) >= 0` states the intent: the caller wants existence, not a position.

## Approach

1. Return `slices.ContainsFunc` with a predicate testing `e.TTL <= 0`.

## Solution

```go
func AnyExpired(entries []Entry) bool {
	return slices.ContainsFunc(entries, func(e Entry) bool { return e.TTL <= 0 })
}
```

## Walkthrough

`AnyExpired([{a 5} {b 0}])` rejects the first entry and returns `true` at the second.

## Pitfalls

- Using `slices.Contains`, which needs a comparable value rather than a predicate.
- Scanning the whole slice and accumulating a bool, losing the early exit.
- Testing `e.TTL < 0` and missing entries that expired exactly now.
