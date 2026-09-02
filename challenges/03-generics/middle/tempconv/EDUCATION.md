# Unit Conversion

## Intuition

Inference reads arguments only, and `FromFloat`'s argument is `float64` for every instantiation — so the result type has to be named explicitly.

## Approach

1. `ToFloat`: return `float64(v)`.
2. `FromFloat`: return `T(f)`.
3. `Rescale`: convert out, call `f`, convert back.

## Solution

```go
func ToFloat[T ~float64](v T) float64 {
	return float64(v)
}

func FromFloat[T ~float64](f float64) T {
	return T(f)
}

func Rescale[T ~float64](v T, f func(float64) float64) T {
	return T(f(float64(v)))
}
```

## Walkthrough

`Rescale(Celsius(20), double)` computes `40.0` in plain float and re-wraps it as `Celsius`.

## Pitfalls

- Calling `FromFloat(20)` without the type argument.
- Returning `float64` from `Rescale` and dropping the unit.
- Assuming the conversion rounds — `T(f)` on `~float64` is exact.
