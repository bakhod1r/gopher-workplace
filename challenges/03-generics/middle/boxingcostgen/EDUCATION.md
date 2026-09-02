# The Cost Of any

## Intuition

The `bool` in `SumAny` exists only because the type system was bypassed. Generics remove the failure mode entirely rather than handling it better.

## Approach

1. `SumTyped`: accumulate in `T`.
2. `SumAny`: assert each element with comma-ok, failing on the first mismatch.

## Solution

```go
func SumTyped[T Number](s []T) T {
	var out T
	for _, v := range s {
		out += v
	}
	return out
}

func SumAny(s []any) (int, bool) {
	total := 0
	for _, v := range s {
		n, ok := v.(int)
		if !ok {
			return 0, false
		}
		total += n
	}
	return total, true
}
```

## Walkthrough

`SumAny([]any{1,"x"})` fails on the string; `SumTyped` could never have been given it.

## Pitfalls

- Panicking with an unchecked assertion `v.(int)`.
- Reaching for `[]any` in new code where a type parameter fits.
- Assuming the two versions have identical performance.
