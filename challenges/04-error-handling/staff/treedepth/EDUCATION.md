# Depth Of An Error Tree

## Intuition

Depth is the classic tree recursion, complicated only by Go's two unwrap shapes. Measuring it is what lets a service refuse an error value that would blow the stack during formatting.

## Approach

1. Return 0 for nil.
2. For a joined error, take the maximum branch depth.
3. For a wrapped error, take the child's depth.
4. Add one for the node itself.

## Solution

```go
if err == nil {
	return 0
}
best := 0
if joined, ok := err.(interface{ Unwrap() []error }); ok {
	for _, e := range joined.Unwrap() {
		if d := Depth(e); d > best {
			best = d
		}
	}
} else if wrapped, ok := err.(interface{ Unwrap() error }); ok {
	best = Depth(wrapped.Unwrap())
}
return best + 1
```

## Walkthrough

`errors.Join(ErrA, fmt.Errorf("x: %w", ErrB))` is the join node plus its deepest branch of two, giving 3.

## Pitfalls

- Counting the join node as free, undercounting by one.
- Summing branch depths instead of taking the maximum.
- Handling only one unwrap shape.
