# Split A Joined Error

## Intuition

Go has two unwrapping shapes. `Unwrap() error` walks a chain of causes; `Unwrap() []error` fans out into siblings. `errors.Is` understands both, but code that inspects a chain must ask for the right one.

## Approach

1. Return nil for a nil error.
2. Assert to `interface{ Unwrap() []error }`.
3. Return the slice on success, otherwise `[]error{err}`.

## Solution

```go
if err == nil {
	return nil
}
if joined, ok := err.(interface{ Unwrap() []error }); ok {
	return joined.Unwrap()
}
return []error{err}
```

## Walkthrough

`errors.Join(nil, ErrB)` already dropped the nil at construction time, so the split yields a single element.

## Pitfalls

- Calling `errors.Unwrap`, which returns nil for a joined error.
- Returning nil instead of a one-element slice for a plain error.
- Assuming the join preserved nil entries.
