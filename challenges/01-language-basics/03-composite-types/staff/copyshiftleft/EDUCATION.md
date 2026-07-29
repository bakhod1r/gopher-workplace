# Overlapping copy and direction

## Intuition

`copy(dst, src)` copies `min(len)` elements and handles overlap like `memmove`.
To shift left (drop index 0), the destination is the earlier region:

```go
copy(xs, xs[1:]) // xs[i] = xs[i+1]
xs[len(xs)-1] = 0
```

`copy(xs[1:], xs)` shifts the other way, replicating `xs[0]`.

## Approach

1. Bug: `copy(xs[1:], xs)` copies forward into overlapping memory, duplicating xs[0].
2. Fix: `copy(xs, xs[1:])` shifts every element left one; then clear the last slot.

## Solution

```go
func ShiftLeft(xs []int) {
	if len(xs) == 0 {
		return
	}
	copy(xs, xs[1:])
	xs[len(xs)-1] = 0
}
```

## Walkthrough

[1 2 3]: copy(xs, xs[1:]) copies [2 3] onto positions 0,1 -> [2 3 3], then xs[2]=0 -> [2 3 0]. The buggy copy(xs[1:],xs) would give [1 1 2] before clearing -> [1 1 0].

## Pitfalls

- Left shift: `copy(xs, xs[1:])`. Right shift: `copy(xs[1:], xs)`.
- `copy` returns the count copied (`len-1` here).
- Remember to clear the vacated slot to avoid a stale/duplicated value.
