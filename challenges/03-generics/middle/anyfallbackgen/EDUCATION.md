# When Generics Are Not Enough

## Intuition

Knowing when to stop reaching for generics is part of using them well: unbounded nesting is a genuine limit, not a puzzle to work around.

## Approach

1. Switch on the dynamic type.
2. Recurse into every element of a `[]any`.
3. Count `nil` as zero and anything else as one.

## Solution

```go
func DeepCount(v any) int {
	switch x := v.(type) {
	case []any:
		total := 0
		for _, e := range x {
			total += DeepCount(e)
		}
		return total
	case nil:
		return 0
	default:
		return 1
	}
}
```

## Walkthrough

`DeepCount([]any{1, []any{2, 3}})` counts `1`, then recurses to count `2` and `3`.

## Pitfalls

- Trying to write `DeepCount[T any](v []T)`, which handles exactly one depth.
- Forgetting the `nil` case and counting it as a leaf.
- Using reflection where a type switch is enough.
