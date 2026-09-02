# Sort By Field

## Intuition

The three-way comparison is what lets the sort distinguish "equal" from "less", which is what stability is built on.

## Approach

1. Call `slices.SortStableFunc` with a comparison returning `cmp.Compare(a.Age, b.Age)`.

## Solution

```go
func SortByAge(people []Person) {
	slices.SortStableFunc(people, func(a, b Person) int {
		return cmp.Compare(a.Age, b.Age)
	})
}
```

## Walkthrough

Two people aged 20 compare as `0`, so the stable sort leaves them in their original relative order.

## Pitfalls

- Returning a `bool` from the comparison, as `sort.Slice` wanted.
- Using `slices.SortFunc`, which may reorder equal elements.
- Subtracting ages (`a.Age - b.Age`), which overflows for extreme values.
