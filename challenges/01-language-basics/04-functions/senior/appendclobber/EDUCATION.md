# Append and shared backing arrays

## Intuition

Appending to a slice with spare capacity mutates shared memory; a full-slice expression `s[:len:len]` sets cap=len so the next append copies.

## Approach

1. Two appends onto the same base with spare capacity write into the **same** backing array.
2. Clip base capacity first: `base[:len(base):len(base)]`, so each append reallocates.

## Solution

```go
func TwoTails(base []int, x, y int) (a, b []int) {
	clipped := base[:len(base):len(base)]
	a = append(clipped, x)
	b = append(clipped, y)
	return a, b
}
```

## Walkthrough

Without clipping, `b`'s append overwrites the slot `a` just wrote, so `a[2]` becomes 200. The full-slice clip forces separate arrays.

## Pitfalls

- `append` reuses capacity when available — two appends to the same base collide.
- `base[:n:n]` forces the next append to allocate fresh.
