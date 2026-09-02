# Equal From Stdlib

## Intuition

The stdlib codifies the same length-then-elements rule you would write by hand, including the nil-versus-empty decision that trips people up.

## Approach

1. Return `slices.Equal(a, b)`.

## Solution

```go
func SameOrder[T comparable](a, b []T) bool {
	return slices.Equal(a, b)
}
```

## Walkthrough

`SameOrder(nil, []int{})` compares two zero lengths and returns `true`.

## Pitfalls

- Using `reflect.DeepEqual`, which is slower and distinguishes nil from empty.
- Writing `a == b`, which does not compile.
- Using `slices.EqualFunc` when plain equality is enough.
