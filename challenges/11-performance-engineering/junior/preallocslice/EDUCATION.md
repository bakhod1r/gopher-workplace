# One Allocation, Not Eleven

## Intuition

`append` doubles the backing array when it runs out of room. Telling `make` the final size up front removes every one of those doublings.

## Approach

1. Clamp a negative `n` to zero.
2. `make([]int, 0, n)`.
3. Append the squares.

## Solution

```go
func Squares(n int) []int {
	if n < 0 {
		n = 0
	}
	out := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		out = append(out, i*i)
	}
	return out
}
```

## Walkthrough

Growing from nil to 1000 `int`s takes about eleven allocations and copies roughly 2000 elements; the capacity hint makes it one allocation and 1000 writes.

## Pitfalls

- `make([]int, n)` plus `append`, which returns 2n elements, half of them zero.
- `var out []int`, which is nil for the empty case and regrows for every other.
- Preallocating a wildly overestimated capacity, which trades time for wasted memory.
