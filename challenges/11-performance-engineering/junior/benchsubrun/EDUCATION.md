# Naming The Sub-Benchmarks

## Intuition

A sub-benchmark's identity is its full path. Build it the same way `b.Run` does and your tooling can find it again.

## Approach

1. Allocate the result with a capacity hint.
2. Append one formatted name per size.

## Solution

```go
func Names(base string, sizes []int) []string {
	out := make([]string, 0, len(sizes))
	for _, size := range sizes {
		out = append(out, fmt.Sprintf("%s/size=%d", base, size))
	}
	return out
}
```

## Walkthrough

`make([]string, 0, len(sizes))` returns a non-nil zero-length slice when `sizes` is nil, so the empty case needs no special branch.

## Pitfalls

- `var out []string`, which returns nil for empty input.
- `make([]string, len(sizes))` and then appending, which leaves empty strings in front.
- Sorting the sizes; the caller's order is the reporting order.
