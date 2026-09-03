# Root Cause

## Intuition

An error chain is a linked list. Finding the root is walking to the end and remembering the last node, not the nil that follows it.

## Approach

1. Loop while `errors.Unwrap(err)` is non-nil.
2. Move `err` down one link each pass.
3. Return `err`.

## Solution

```go
for {
	next := errors.Unwrap(err)
	if next == nil {
		return err
	}
	err = next
}
```

## Walkthrough

For three layers the loop steps a, b, c, reaches `ErrBase`, sees `Unwrap` return nil, and returns `ErrBase`.

## Pitfalls

- Returning `errors.Unwrap(err)` after the loop, which is always nil.
- Special-casing nil unnecessarily — `errors.Unwrap(nil)` is nil, so the loop returns nil on the first pass.
- Recursing without a base case for non-wrapping errors.
