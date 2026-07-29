# Variadic slice parameters

## Intuition

`...[]int` collects trailing slice arguments; spreading each with `g...` builds a flat concatenation.

## Approach

1. Range the variadic groups.
2. Spread each into the output with `append(out, g...)`.

## Solution

```go
func Flatten(groups ...[]int) []int {
	var out []int
	for _, g := range groups {
		out = append(out, g...)
	}
	return out
}
```

## Walkthrough

`Flatten([1 2], [3], [4 5])` concatenates all groups into `[1 2 3 4 5]`.

## Pitfalls

- A nil group contributes nothing (spreads to zero elements).
- Return empty, not nil-panic, for no groups.
