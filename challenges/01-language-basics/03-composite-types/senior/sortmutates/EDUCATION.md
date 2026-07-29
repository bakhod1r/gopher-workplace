# "Copy" that isn't

## Intuition

`out := xs` copies the slice header (pointer/len/cap) — both names point at the
same backing array. `sort.Ints(out)` therefore sorts `xs` too. Duplicate the data
to isolate:

```go
out := append([]int{}, xs...) // independent copy
sort.Ints(out)
```

## Approach

1. Bug: out := xs aliases the caller's slice; sort.Ints(out) sorts in place, corrupting xs. 2. Fix: copy first: out := append([]int(nil), xs...). 3. Sorting the independent copy leaves the input untouched.

## Solution

```go
import "sort"

func SortedCopy(xs []int) []int {
	out := append([]int(nil), xs...)
	sort.Ints(out)
	return out
}
```

## Walkthrough

out=xs shares the array; sort.Ints sorts it so xs is reordered too. A copied out is sorted independently, xs preserved.

## Pitfalls

- Assigning a slice never copies its elements.
- `slices.Sorted`/`slices.Clone` avoid the footgun.
- `sort.Ints`, `sort.Slice`, and `slices.Sort` all mutate in place.
