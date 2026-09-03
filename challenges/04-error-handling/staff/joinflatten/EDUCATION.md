# Flatten The Tree

## Intuition

Once joins exist, an error is a tree: `Unwrap() error` is a single child, `Unwrap() []error` is many. Flattening means recursing on whichever shape a node implements.

## Approach

1. Return nil for nil.
2. Recurse over `Unwrap() []error` when present.
3. Recurse into `Unwrap() error` when present.
4. Otherwise return the error as a leaf.

## Solution

```go
if err == nil {
	return nil
}
if joined, ok := err.(interface{ Unwrap() []error }); ok {
	var out []error
	for _, e := range joined.Unwrap() {
		out = append(out, Leaves(e)...)
	}
	return out
}
if wrapped, ok := err.(interface{ Unwrap() error }); ok {
	return Leaves(wrapped.Unwrap())
}
return []error{err}
```

## Walkthrough

`errors.Join(ErrA, errors.Join(ErrB, ErrC))` descends into the inner join, so the result is three leaves in left-to-right order.

## Pitfalls

- Handling only the joined shape, so wrapped errors are reported instead of their causes.
- Returning wrappers as leaves.
- Using `errors.Unwrap`, which returns nil for a joined error and silently truncates the walk.
