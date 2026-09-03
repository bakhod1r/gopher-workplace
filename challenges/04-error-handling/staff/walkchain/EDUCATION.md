# Visit Every Error

## Intuition

Traversal order is part of the contract for anything that emits events. Pre-order means the aggregate is reported before the details it contains, which is what a reader expects.

## Approach

1. Return for a nil error.
2. Call `visit(err)`.
3. Recurse over joined branches, or into a wrapped child.

## Solution

```go
if err == nil {
	return
}
visit(err)
if joined, ok := err.(interface{ Unwrap() []error }); ok {
	for _, e := range joined.Unwrap() {
		Walk(e, visit)
	}
	return
}
if wrapped, ok := err.(interface{ Unwrap() error }); ok {
	Walk(wrapped.Unwrap(), visit)
}
```

## Walkthrough

A join of two leaves produces three visits: the join's own combined message, then each branch in order.

## Pitfalls

- Visiting children before the node, inverting the order.
- Skipping the join node itself.
- Recursing into a wrapped child after already handling the joined shape, double-visiting.
