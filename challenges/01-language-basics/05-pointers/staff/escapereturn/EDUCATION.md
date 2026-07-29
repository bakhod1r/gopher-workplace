# Aliasing vs escaping copies

## Intuition

`&localCopy` escapes to the heap but is independent of the slice; `&xs[i]` addresses the backing array so writes are visible in the slice.

## Approach

1. `v := xs[best]; return &v` returns the address of a **copy**.
2. Return `&xs[best]` to alias the real element.

## Solution

```go
func MaxPtr(xs []int) *int {
	best := 0
	for i := range xs {
		if xs[i] > xs[best] {
			best = i
		}
	}
	return &xs[best]
}
```

## Walkthrough

Writing through the bug's pointer changes a dead local, not the slice. `&xs[best]` lets the caller mutate the actual max element.

## Pitfalls

- `v := xs[i]; &v` gives an independent pointer.
- `&xs[i]` gives an aliasing pointer into the slice.
