# First Failure

## Intuition

Errors travel in slices as often as they travel alone. Reducing a slice of errors to one answer is the same shape as any other search loop.

## Approach

1. Range over the slice.
2. Return the element as soon as it is non-nil.
3. Return nil after the loop when nothing failed.

## Solution

```go
for _, err := range errs {
	if err != nil {
		return err
	}
}
return nil
```

## Walkthrough

For `[]error{nil, ErrB, ErrC}` the first iteration skips nil, the second returns `ErrB`, and `ErrC` is never examined.

## Pitfalls

- Overwriting a variable in the loop and returning the *last* error.
- Special-casing a nil slice — `for range nil` already does nothing.
- Returning a new error instead of the one found.
