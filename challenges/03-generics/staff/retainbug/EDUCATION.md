# The Small Result That Pins A Huge Array

## Intuition

Capping the capacity stops the caller appending over `s`, but the header still addresses `s`'s array. The garbage collector cannot free part of an allocation, so a four-element view pins the entire multi-megabyte payload for as long as the view lives.

## Approach

1. Clamp `n` into range.
2. Allocate a fresh slice of length `n`.
3. Copy the prefix into it and return the copy.

## Solution

```go
func Head[T any](s []T, n int) []T {
	if n < 0 {
		n = 0
	}
	if n > len(s) {
		n = len(s)
	}
	out := make([]T, n)
	copy(out, s)
	return out
}
```

## Walkthrough

Taking a 4-element head from each of sixty 8 MB payloads retains about 480 MB with the slice view and about a kilobyte with the copy.

## Pitfalls

- Treating `s[:n:n]` as "the safe slice" — it is safe against `append`, not against retention.
- Copying only when `n` is small; the caller's lifetime, not the size, decides.
