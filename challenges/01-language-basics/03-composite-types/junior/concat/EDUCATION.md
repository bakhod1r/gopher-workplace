# Variadic parameters and append spread

## Intuition

A variadic `...[]int` parameter arrives as a `[][]int`. Flatten with the spread
operator:

```go
out := []int{}
for _, s := range slices { out = append(out, s...) }
```

## Approach

1. Start with an empty result slice.
2. Range over the variadic slices.
3. append(result, s...) spreads each slice's elements in order (nil spreads to nothing).
4. Return result.

## Solution

```go
func Concat(slices ...[]int) []int {
	result := []int{}
	for _, s := range slices {
		result = append(result, s...)
	}
	return result
}
```

## Walkthrough

Concat([1,2],[3],nil,[4,5]): result grows [1,2] -> [1,2,3] -> unchanged by nil -> [1,2,3,4,5].

## Pitfalls

- `append(out, s...)` spreads `s`; without `...` it's a type error.
- Appending a `nil` slice is a safe no-op.
- Pre-size with the total length to avoid regrowth.
