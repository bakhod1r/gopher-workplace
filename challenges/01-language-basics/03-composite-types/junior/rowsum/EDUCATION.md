# 2-D data as slices of slices

## Intuition

Go's 2-D structures are slices of slices, and they can be **ragged**. Nested
ranging handles that:

```go
for _, row := range grid {
	s := 0
	for _, v := range row { s += v }
	out = append(out, s)
}
```

## Approach

1. Start with an empty result slice.
2. Range each row; inner-range to accumulate its sum.
3. Append each row's sum (0 for empty rows).
4. Return result.

## Solution

```go
func RowSums(grid [][]int) []int {
	result := []int{}
	for _, row := range grid {
		sum := 0
		for _, v := range row {
			sum += v
		}
		result = append(result, sum)
	}
	return result
}
```

## Walkthrough

RowSums({{1,2,3},{4,5},{}}): 1+2+3=6; 4+5=9; empty=0 -> [6,9,0].

## Pitfalls

- Rows can differ in length; don't assume a fixed width.
- An empty row sums to 0.
- Rows are separate allocations and may be shared.
