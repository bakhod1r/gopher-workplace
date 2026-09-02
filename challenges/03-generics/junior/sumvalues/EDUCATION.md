# Sum Map Values

## Intuition

The key never takes part in arithmetic, so it only needs to be a legal map key. Constraining it more tightly would reject valid callers for no benefit.

## Approach

1. Declare `var total V`.
2. Add each value.
3. Return the total.

## Solution

```go
func SumValues[K comparable, V Number](m map[K]V) V {
	var total V
	for _, v := range m {
		total += v
	}
	return total
}
```

## Walkthrough

`SumValues(map[string]int{"a": 1, "b": 2})` yields `3` regardless of which key it visits first.

## Pitfalls

- Constraining `K` to `Number` too, which rejects string-keyed maps.
- Summing the keys instead of the values.
- Returning `int` when the values are `float64`.
