# Appending twice to one base

## Intuition

If `base` has spare capacity, `append(base, x)` writes into that capacity without
reallocating — and so does a second `append(base, y)`, into the **same** slot.
The two results share memory and stomp each other. Clip first:

```go
base := a[:len(a):len(a)] // cap == len -> every append reallocates
```

## Approach

1. Bug: `base := a` keeps a's spare capacity, so append(base,x) and append(base,y) write to the SAME backing slot index 2; c's write clobbers b, giving b=[1 2 4].
2. Fix: `base := a[:len(a):len(a)]` clips capacity to length, forcing each append to allocate its own array.

## Solution

```go
func Branch(a []int, x, y int) ([]int, []int) {
	base := a[:len(a):len(a)]
	b := append(base, x)
	c := append(base, y)
	return b, c
}
```

## Walkthrough

a has len 2, cap 10. Unclipped, append(base,3) writes 3 at index 2 in place (b shares array); append(base,4) writes 4 at the same index 2, so b now reads [1 2 4]. Clipping cap to 2 makes both appends grow into fresh arrays -> b=[1 2 3], c=[1 2 4].

## Pitfalls

- Only clipping (or copying) guarantees independence.
- `append` reuses spare capacity; that's the whole hazard.
- `slices.Clip` expresses `a[:len(a):len(a)]`.
