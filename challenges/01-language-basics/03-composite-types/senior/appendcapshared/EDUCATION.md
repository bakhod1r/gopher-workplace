# The three-index slice expression

## Intuition

`xs[:2]` has length 2 but inherits `xs`'s **capacity**, so `append` writes into
the shared tail (`xs[2]`). The three-index form `xs[low:high:max]` caps capacity,
forcing `append` to allocate a fresh array:

```go
head := xs[:2:2] // len 2, cap 2 -> append reallocates
return append(head, extra)
```

## Approach

1. Bug: head := xs[:2] shares xs's backing array AND its spare capacity, so append(head, extra) writes into xs[2] instead of a fresh array. 2. Fix: use a full slice expression head := xs[:2:2] which caps capacity at 2. 3. Now append sees no spare capacity and allocates a new backing array, leaving xs untouched.

## Solution

```go
func FirstTwoPlus(xs []int, extra int) []int {
	head := xs[:2:2]
	return append(head, extra)
}
```

## Walkthrough

With xs=[1,2,3]: xs[:2] has len 2 cap 3. append writes extra at index 2 of the shared array -> xs becomes [1,2,9]. With xs[:2:2] cap is 2, append must grow -> new array [1,2,9], xs stays [1,2,3].

## Pitfalls

- Capacity, not length, decides whether append reuses memory.
- The bug only manifests when spare capacity exists — easy to miss in tests.
- `slices.Clip` sets cap==len for the same effect.
