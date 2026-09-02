# Sort Interface

## Intuition

This is how `sort.Sort` works: the algorithm is written once against `Len/Less/Swap`, and any collection that supplies those three becomes sortable.

## Approach

1. `Len` returns `len(s)`; `Less` returns `s[i] < s[j]`; `Swap` exchanges with a tuple assignment.
2. In `BubbleSort`, read `data.Len()` once.
3. For each pass, compare neighbours with `Less(j+1, j)` and `Swap` when out of order.

## Solution

```go
func (s IntSlice) Len() int { return len(s) }

func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }

func (s IntSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func BubbleSort(data Sortable) {
	n := data.Len()
	for i := 0; i < n; i++ {
		for j := 0; j < n-1-i; j++ {
			if data.Less(j+1, j) {
				data.Swap(j, j+1)
			}
		}
	}
}
```

## Walkthrough

`IntSlice{3,1,2}`: first pass swaps 3 and 1, then 3 and 2, giving `[1 2 3]`; later passes find nothing to swap.

## Pitfalls

- Indexing `data` directly — `Sortable` has no elements, only methods.
- `Less(j, j+1)` in the condition, which sorts in the wrong direction.
- Looping `j` to `n-1` and reading index `n`, which panics.
