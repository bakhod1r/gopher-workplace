# Percent

## Intuition

Dividing in `T` first would truncate for integer instantiations — `Percent(1, 4)` would come out as `0`. Converting up front makes one implementation right for every member of the set.

## Approach

1. Return `0` when `whole` is zero.
2. Convert both values to `float64`.
3. Divide and multiply by 100.

## Solution

```go
func Percent[T Number](part, whole T) float64 {
	if whole == 0 {
		return 0
	}
	return float64(part) / float64(whole) * 100
}
```

## Walkthrough

`Percent(1, 4)` converts to `1.0 / 4.0`, giving `0.25`, then scales to `25`.

## Pitfalls

- Computing `part / whole` in `T` and converting afterwards.
- Omitting the zero guard, which panics for integers.
- Multiplying by 100 before dividing when `T` is a small integer type — it can overflow.
