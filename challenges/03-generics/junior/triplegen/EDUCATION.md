# Triple

## Intuition

The compiler checks the rotation for you: assigning `t.Second` (type `B`) into a `First` field of type `B` only lines up if the return type is written `Triple[B, C, A]`.

## Approach

1. `MakeTriple`: set all three fields.
2. `Rotated`: return `Triple[B, C, A]` with second, third, first in that order.

## Solution

```go
func MakeTriple[A, B, C any](a A, b B, c C) Triple[A, B, C] {
	return Triple[A, B, C]{First: a, Second: b, Third: c}
}

func (t Triple[A, B, C]) Rotated() Triple[B, C, A] {
	return Triple[B, C, A]{First: t.Second, Second: t.Third, Third: t.First}
}
```

## Walkthrough

`MakeTriple(1, "a", true).Rotated()` yields `Triple[string, bool, int]{"a", true, 1}`.

## Pitfalls

- Rotating the values but leaving the return type as `Triple[A, B, C]`.
- Rotating right instead of left.
- Trying to give `Rotated` its own type parameter.
