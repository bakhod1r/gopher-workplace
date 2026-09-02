# Scaling A Unit

## Intuition

Encoding dimensional rules in the signature is the cheapest form of unit safety: adding metres to metres is allowed, multiplying metres by metres is not expressible.

## Approach

1. `Times`: convert out, multiply, convert back.
2. `SumUnits`: accumulate in `T`.

## Solution

```go
func Times[T ~float64](v T, factor float64) T {
	return T(float64(v) * factor)
}

func SumUnits[T ~float64](s []T) T {
	var out T
	for _, v := range s {
		out += v
	}
	return out
}
```

## Walkthrough

`Times(Meters(2), 3)` returns `Meters(6)`; there is no overload that would accept two `Meters`.

## Pitfalls

- Typing the factor as `T`, which permits meaningless products.
- Returning `float64` and dropping the unit.
- Mixing `Meters` and `Seconds` in one `SumUnits` call — the single `T` already forbids it.
