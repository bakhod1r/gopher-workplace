# Sort In Place

## Intuition

`slices.Sort` replaced `sort.Slice` for ordered element types: no closure, no reflection, and the compiler checks the element type is sortable.

## Approach

1. Call `slices.Sort(names)`.

## Solution

```go
func SortNames(names []string) {
	slices.Sort(names)
}
```

## Walkthrough

Sorting works through the shared backing array, so the caller's `s` reads `["a", "b"]` after the call returns.

## Pitfalls

- Assigning the result: `slices.Sort` returns nothing.
- Cloning first, which leaves the caller's slice unsorted.
- Reaching for `sort.Slice` with a hand-written `less`.
