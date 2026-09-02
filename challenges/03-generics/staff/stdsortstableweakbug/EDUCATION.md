# A Comparator That Cannot Decide

## Intuition

`f(a, a)` returns -1: the comparator claims every element is strictly less than itself. That is not a strict weak ordering, so `SortStableFunc`'s merge steps take the wrong branch on ties — reversing equal runs, and on larger inputs producing an order that is not sorted at all.

## Approach

1. Clone the input so the caller keeps its order.
2. Compare the two priorities with `cmp.Compare`, which returns 0 for a tie.
3. Let `SortStableFunc` do the rest.

## Solution

```go
func SortByPriority(tasks []Task) []Task {
	out := slices.Clone(tasks)
	slices.SortStableFunc(out, func(a, b Task) int {
		return cmp.Compare(a.Pri, b.Pri)
	})
	return out
}
```

## Walkthrough

Sorting `[{a 2} {c 2}]`: the merge asks `f(c, a)`, is told `-1` ("c is smaller"), and emits `c` first — reversing arrival order.

## Pitfalls

- Writing `if a.Pri < b.Pri { return -1 }; return 1` — the mirror image of the same defect.
- Subtracting the fields (`a.Pri - b.Pri`), which is a valid ordering right up until the subtraction overflows.
