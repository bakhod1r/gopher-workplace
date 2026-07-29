# Collapsing consecutive duplicates

## Intuition

Emit an element only when it differs from the previously emitted one:

```go
out := []int{}
for i, v := range xs {
	if i == 0 || v != xs[i-1] { out = append(out, v) }
}
```

## Approach

1. Start with empty result.
2. Walk each element with its index.
3. Keep it if it is the first element or differs from the previous element.
4. Return result.

## Solution

```go
func Dedupe(xs []int) []int {
	out := []int{}
	for i, v := range xs {
		if i == 0 || v != xs[i-1] {
			out = append(out, v)
		}
	}
	return out
}
```

## Walkthrough

[1,1,2,...]: i0 keep 1; i1 v1==xs[0] skip; i2 v2!=1 keep; ... yields [1,2,3,4].

## Pitfalls

- It only removes **consecutive** duplicates; unsorted input keeps distant repeats.
- `slices.Compact` (Go 1.21+) does this in place.
- Comparing to `xs[i-1]` (source) vs the last kept value is equivalent here since
  runs are contiguous.
