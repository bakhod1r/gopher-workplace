# Min And Max By Comparison

## Intuition

The pattern repeats across the stdlib: a fast, panicking core plus a thin wrapper that encodes your API's contract.

## Approach

1. Return zero and `false` for an empty slice.
2. Otherwise delegate with a `cmp.Compare` on the price.

## Solution

```go
func Cheapest(items []Item) (Item, bool) {
	if len(items) == 0 {
		var zero Item
		return zero, false
	}
	return slices.MinFunc(items, func(a, b Item) int {
		return cmp.Compare(a.Price, b.Price)
	}), true
}

func Priciest(items []Item) (Item, bool) {
	if len(items) == 0 {
		var zero Item
		return zero, false
	}
	return slices.MaxFunc(items, func(a, b Item) int {
		return cmp.Compare(a.Price, b.Price)
	}), true
}
```

## Walkthrough

`Cheapest(nil)` returns before touching `slices.MinFunc`, avoiding its panic.

## Pitfalls

- Delegating without the empty guard.
- Returning a `bool` from the comparison.
- Sorting the slice to find one element.
