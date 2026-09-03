# What `-benchmem` Prints

## Intuition

The memory columns are bookkeeping: total bytes and total allocations counted across the whole run, divided by the number of iterations.

## Approach

1. Return the zero line when `iters <= 0`.
2. Convert `iters` to `uint64` and divide both totals.
3. Format with a tab separator.

## Solution

```go
func Report(bytes, allocs uint64, iters int) string {
	if iters <= 0 {
		return "0 B/op\t0 allocs/op"
	}
	n := uint64(iters)
	return fmt.Sprintf("%d B/op\t%d allocs/op", bytes/n, allocs/n)
}
```

## Walkthrough

`Report(10, 3, 4)` gives `10/4 = 2` bytes and `3/4 = 0` allocations: an allocation that happens on some iterations but not all rounds away entirely.

## Pitfalls

- Converting `iters` before checking it is positive — `uint64(-1)` is huge.
- Rounding instead of truncating; the real tool truncates.
- A space instead of the tab, which breaks tools that parse benchmark output.
