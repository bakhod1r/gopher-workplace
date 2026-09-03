# Clip Applied To The Wrong Slice

## Intuition

`slices.Clip(s)` returns `s[:len(s):len(s)]`. Re-slicing that to `[:n]` keeps the *capacity* of the full slice, so the result again has room to append over `s[n]`. Clipping must be the last operation, applied to the already-shortened slice.

## Approach

1. Clamp `n` into range.
2. Shorten first with `s[:n]`.
3. Clip that, so capacity is trimmed to `n`.

## Solution

```go
func Shrink[T any](s []T, n int) []T {
	if n < 0 {
		n = 0
	}
	if n > len(s) {
		n = len(s)
	}
	return slices.Clip(s[:n])
}
```

## Walkthrough

For `s = [1 2 3 4]` and `n = 2`, the buggy result has length 2 and capacity 4. `append(page, 99)` writes `99` over `s[2]`.

## Pitfalls

- Believing `Clip` makes a copy — it only narrows the header.
- Reaching for `slices.Grow` here; growing is the opposite of what is needed.
