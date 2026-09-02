# Equal With Custom Equality

## Intuition

The cross-type signature is what makes this useful in tests: expected and actual rarely share a type.

## Approach

1. Return `slices.EqualFunc(a, b, eq)`.

## Solution

```go
func SameRows[T, U any](a []T, b []U, eq func(T, U) bool) bool {
	return slices.EqualFunc(a, b, eq)
}
```

## Walkthrough

`SameRows(nil, []string{}, matches)` compares two zero lengths and reports `true`.

## Pitfalls

- Hand-rolling the loop you already wrote once.
- Using `reflect.DeepEqual`, which cannot express fuzzy equality.
- Assuming the predicate is called for mismatched lengths — it is not.
