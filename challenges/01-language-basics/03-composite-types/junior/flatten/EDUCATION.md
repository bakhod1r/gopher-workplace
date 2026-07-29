# Row-major flattening

## Intuition

Flattening a slice of slices concatenates the rows in order:

```go
out := []int{}
for _, row := range grid { out = append(out, row...) }
```

## Approach

1. Start with an empty result slice.
2. Range the rows in order (row-major).
3. append(result, row...) spreads each row's elements.
4. Return result.

## Solution

```go
func Flatten(grid [][]int) []int {
	result := []int{}
	for _, row := range grid {
		result = append(result, row...)
	}
	return result
}
```

## Walkthrough

Flatten({{1,2},{3},{},{4,5}}): [1,2] -> +3 -> +nothing -> +4,5 = [1,2,3,4,5].

## Pitfalls

- `append(out, row...)` spreads the row; without `...` it's a type error.
- Empty and nil rows contribute nothing.
- A single `[]int` indexed by `r*w+c` is more cache-friendly for fixed width.
