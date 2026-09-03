# Emptying Without Freeing

## Intuition

A slice is a length, a capacity, and a pointer. Setting the length to zero throws away the *view*, not the memory.

## Approach

1. `Reset` returns `s[:0]`, with a nil guard for the non-nil contract.
2. `FillEvens` resets first, then appends `n` values.

## Solution

```go
func Reset(s []int) []int {
	if s == nil {
		return []int{}
	}
	return s[:0]
}

func FillEvens(buf []int, n int) []int {
	out := Reset(buf)
	for i := 0; i < n; i++ {
		out = append(out, 2*i)
	}
	return out
}
```

## Walkthrough

Appending into `buf[:0]` writes over the old elements one by one, so the caller sees fresh data and the runtime allocates nothing.

## Pitfalls

- `s = nil` as a reset, which drops the array and forces a fresh allocation.
- Keeping a long-lived buffer full of pointers; the elements past `len` stay reachable and the objects they point at cannot be collected. `clear(s)` before reslicing fixes that.
- Handing the same reused buffer to two callers, who then overwrite each other.
