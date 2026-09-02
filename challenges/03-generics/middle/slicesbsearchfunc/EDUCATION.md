# Binary Search With Comparison

## Intuition

The asymmetric signature is deliberate: it lets you search a `[]Row` by an `int` without constructing a dummy element.

## Approach

1. Return `slices.BinarySearchFunc` with `cmp.Compare(r.ID, target)`.

## Solution

```go
func FindByID(rows []Row, id int) (int, bool) {
	return slices.BinarySearchFunc(rows, id, func(r Row, target int) int {
		return cmp.Compare(r.ID, target)
	})
}
```

## Walkthrough

Searching for a missing ID returns the position where a row with that ID would belong.

## Pitfalls

- Writing a comparison that takes two `Row` values.
- Returning the comparison the wrong way round, which reverses the search.
- Trusting the result on unsorted input.
