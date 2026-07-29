# Slices are references under a value header

## Intuition

Passing a slice copies its (ptr,len,cap) header but not the backing array, so element writes are visible to the caller — reads are always safe.

## Approach

1. Range and add each value into the named `sum`.

## Solution

```go
func SumKeep(xs []int) (sum int) {
	for _, v := range xs {
		sum += v
	}
	return
}
```

## Walkthrough

`SumKeep([1 2 3])` accumulates to 6.

## Pitfalls

- Don't sort or assign into `xs[i]` — that would mutate the caller's array.
- The named return `sum` starts at 0.
