# The Whole Numeric Set

## Intuition

Widening a constraint narrows the operations you may rely on. `+` survives; anything assuming a sign does not.

## Approach

1. Declare `var out T`.
2. Add each element.
3. Return the total.

## Solution

```go
func Total[T Number](s []T) T {
	var out T
	for _, v := range s {
		out += v
	}
	return out
}
```

## Walkthrough

`Total([]uint{1, 2})` instantiates `T = uint`, where the same `+=` is correct without any sign handling.

## Pitfalls

- Adding a `v < 0` branch, which is always false for unsigned instantiations.
- Starting from a literal that is not representable in every member.
- Splitting the function per kind, which is what generics exist to avoid.
