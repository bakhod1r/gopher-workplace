# Bounded Value

## Intuition

When limits come from outside, they are input too. A reversed range is a caller bug and deserves its own error rather than being reported as a value that failed to fit.

## Approach

1. Reject `lo > hi` first.
2. Reject `n < lo || n > hi`.
3. Return nil otherwise.

## Solution

```go
if lo > hi {
	return ErrBadBounds
}
if n < lo || n > hi {
	return ErrOutOfRange
}
return nil
```

## Walkthrough

`InRange(5, 10, 1)` never reaches the value check — the bounds guard fires first and names the real problem.

## Pitfalls

- Skipping the bounds check, so every value looks out of range.
- Treating the range as exclusive.
- Checking the value first and reporting `ErrOutOfRange` for a caller's reversed bounds.
