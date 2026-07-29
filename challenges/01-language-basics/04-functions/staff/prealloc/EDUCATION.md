# make length vs capacity

## Intuition

`make([]T, n)` fills n zero elements; `make([]T, 0, n)` reserves capacity but stays empty — mixing the two with append yields leading zeros.

## Approach

1. `make([]int, n)` pre-fills `n` zeros; appending then adds after them.
2. Use `make([]int, 0, n)` to preallocate capacity without length.

## Solution

```go
func Doubles(n int) []int {
	out := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		out = append(out, i*2)
	}
	return out
}
```

## Walkthrough

With length `n`, appending yields `[0 0 0 2 4 6]`. Starting at length 0 with capacity `n` appends into the reserved space, giving `[2 4 6]`.

## Pitfalls

- `make([]T, n)` + append = n zeros then your data.
- Use `make([]T, 0, n)` to reserve, or index into a length-n slice.
