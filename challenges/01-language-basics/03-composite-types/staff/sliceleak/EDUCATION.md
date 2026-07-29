# Sub-slices pin the whole backing array

## Intuition

A slice keeps its backing array reachable for the GC — the *entire* array, not
just the visible window. `xs[:3]` of a 1M-capacity slice keeps all 1M alive:

```go
return append([]int{}, xs[:3]...) // independent; source array can be freed
```

## Approach

1. Bug: `return xs[:3]` shares xs's backing array (cap 100000), keeping the whole array alive -> memory leak.
2. Fix: `return append([]int(nil), xs[:3]...)` copies the 3 elements into a small fresh slice (cap <= 3).

## Solution

```go
func Head3(xs []int) []int {
	return append([]int(nil), xs[:3]...)
}
```

## Walkthrough

xs cap 100000. xs[:3] has cap 100000 and pins the array. The append-copy yields a new slice holding [1 2 3] with cap 3, letting the big array be collected.

## Pitfalls

- `cap`, not `len`, reveals how much is retained.
- Same leak with sub-strings of huge strings (copy with `strings.Clone`).
- `slices.Clip` reduces capacity but still shares the array; a copy fully
  releases it.
