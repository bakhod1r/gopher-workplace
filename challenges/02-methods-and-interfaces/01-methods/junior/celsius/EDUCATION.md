# Methods on Defined Types

## Intuition

In Go, you can define a **new named type** from any existing type: `type Celsius
float64`. The new type has no methods by default, but you can attach methods to
it — something you cannot do with plain `float64`.

This is Go's way of giving meaning to raw types: `Celsius(37)` carries intent
that `float64(37)` does not.

## Approach

1. Convert `c` to `float64`.
2. Apply the formula `C × 9/5 + 32`.
3. Return the result.

## Solution

```go
func (c Celsius) ToFahrenheit() float64 {
	return float64(c)*9/5 + 32
}
```

## Walkthrough

For `Celsius(100)`:
- `float64(100) * 9 / 5` = 180.
- `180 + 32` = 212.

## Pitfalls

- Integer division trap: `9/5` in Go integer math is `1`. Use `float64`
  arithmetic or write `9.0/5.0`.
- You cannot add methods to `float64` itself — only to your own defined type.
- `Celsius` and `float64` are distinct types; explicit conversion is needed.
