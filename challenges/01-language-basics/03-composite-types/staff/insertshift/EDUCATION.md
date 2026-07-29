# Opening a gap for insertion

## Intuition

In-place insertion at `i` first grows the slice by one, then shifts the tail one
step **right** to open a gap:

```go
xs = append(xs, 0)
copy(xs[i+1:], xs[i:]) // dst starts one past src -> shift right
xs[i] = v
```

`copy` is overlap-safe (memmove semantics), so the right-shift doesn't corrupt.

## Approach

1. Bug: `copy(xs[i:], xs[i+1:])` shifts the tail LEFT, dropping elements instead of making room.
2. Fix: `copy(xs[i+1:], xs[i:])` shifts the tail RIGHT by one so index i is free for v.

## Solution

```go
func InsertAt(xs []int, i, v int) []int {
	xs = append(xs, 0) // make room at the end
	copy(xs[i+1:], xs[i:])
	xs[i] = v
	return xs
}
```

## Walkthrough

xs=[1 2 3], append 0 -> [1 2 3 0]. i=1: copy(xs[2:],xs[1:]) shifts [2 3] to positions 2,3 -> [1 2 2 3], then xs[1]=9 -> [1 9 2 3].

## Pitfalls

- Insert shifts right; delete shifts left.
- Grow the slice before shifting, or you write past the end.
- `copy` copies `min(len(dst),len(src))` — the counts line up here.
