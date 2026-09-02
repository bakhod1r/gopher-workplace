# The min And max Builtins

## Intuition

The builtins work on type parameters because they require exactly what `cmp.Ordered` promises — no import and no instantiation cost.

## Approach

1. `Middle`: raise `v` to `lo` with `max`, then cap it with `min`.
2. `Spread`: subtract the three-argument `min` from the three-argument `max`.

## Solution

```go
func Middle[T cmp.Ordered](v, lo, hi T) T {
	return min(max(v, lo), hi)
}

func Spread(a, b, c int) int {
	return max(a, b, c) - min(a, b, c)
}
```

## Walkthrough

`Middle(-1, 0, 3)` computes `max(-1, 0) = 0`, then `min(0, 3) = 0`.

## Pitfalls

- Writing `max(min(v, hi), lo)` — equivalent here, but reversed bounds behave differently.
- Reaching for `slices.Max` with a temporary slice.
- Shadowing `min` or `max` with a local variable of that name.
