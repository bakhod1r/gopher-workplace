# Bounds-Checked Index

## Intuition

Go panics on an out-of-range index. Wrapping the access in a bounds check converts an unrecoverable crash into an ordinary error the caller can handle.

## Approach

1. Reject `i < 0`.
2. Reject `i >= len(s)`.
3. Return `s[i], nil` otherwise.

## Solution

```go
if i < 0 || i >= len(s) {
	return 0, ErrOutOfRange
}
return s[i], nil
```

## Walkthrough

For a nil slice `len(s)` is 0, so every index fails the second condition and the error is returned without touching the slice.

## Pitfalls

- Checking only the upper bound, letting `-1` panic.
- Using `i > len(s)` — index `len(s)` is already past the end.
- Recovering from the panic instead of preventing it.
