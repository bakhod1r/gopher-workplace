# Equality That Reads Too Many Fields

## Intuition

`Line` is `comparable`, so `slices.Equal` type-checks and quietly compares all three fields — including the one the spec says to ignore. The type system cannot know which fields carry identity; only `EqualFunc` lets you say so.

## Approach

1. Use `slices.EqualFunc` so equality is yours to define.
2. Compare `SKU` and `Qty`; say nothing about `Note`.
3. Let `EqualFunc` handle the length check.

## Solution

```go
func SameLines(a, b []Line) bool {
	return slices.EqualFunc(a, b, func(x, y Line) bool {
		return x.SKU == y.SKU && x.Qty == y.Qty
	})
}
```

## Walkthrough

`[{x 1 "picked"}]` versus `[{x 1 ""}]`: identical order content, but `==` sees two different structs and reports a change.

## Pitfalls

- Zeroing `Note` on a copy before comparing — it works and it allocates, on every comparison.
- Adding the ignored field later and forgetting the comparator, which silently re-introduces the bug.
