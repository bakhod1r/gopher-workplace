# Match With A Predicate

## Intuition

`errors.Is` and `errors.As` are two fixed predicates over the same traversal. Exposing the traversal itself lets tooling ask questions the standard library never anticipated.

## Approach

1. Return false for nil.
2. Return true when `pred(err)` holds.
3. Recurse into joined branches, then into a wrapped child.

## Solution

```go
if err == nil {
	return false
}
if pred(err) {
	return true
}
if joined, ok := err.(interface{ Unwrap() []error }); ok {
	for _, e := range joined.Unwrap() {
		if Any(e, pred) {
			return true
		}
	}
	return false
}
if wrapped, ok := err.(interface{ Unwrap() error }); ok {
	return Any(wrapped.Unwrap(), pred)
}
return false
```

## Walkthrough

The nested case wraps a join inside a wrap inside a join; the predicate matches only at the deepest leaf, and every level is visited.

## Pitfalls

- Skipping the node itself and only testing children.
- Returning false as soon as one branch fails instead of trying the rest.
- Recursing infinitely on an error whose `Unwrap` returns itself.
