# To Floats

## Intuition

Conversion of a type-parameter value is allowed when the conversion is valid for every type in the set. All integer types convert to `float64`, so one conversion covers all instantiations.

## Approach

1. Allocate `out` with capacity `len(s)`.
2. Append `float64(v)` for each element.
3. Return `out`.

## Solution

```go
func ToFloats[T Integer](s []T) []float64 {
	out := make([]float64, 0, len(s))
	for _, v := range s {
		out = append(out, float64(v))
	}
	return out
}
```

## Walkthrough

`ToFloats([]int64{7})` instantiates `T = int64` and converts `7` to `7.0`.

## Pitfalls

- Declaring `out` as `[]T` — the result must be `[]float64`.
- Adding `~string` to the constraint, which breaks the conversion.
- Ignoring precision loss for very large `int64` values — real, but out of scope here.
