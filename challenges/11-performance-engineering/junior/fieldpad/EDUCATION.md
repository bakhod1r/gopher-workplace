# Where The Padding Comes From

## Intuition

Lay the fields down left to right. Before writing each one, skip forward to the next legal offset for its alignment. At the end, skip forward once more so the next copy of the struct would also start legally.

## Approach

1. `AlignUp` guards `a <= 1` and rounds otherwise.
2. `StructSize` walks the sizes, aligning the offset and adding the size, tracking the maximum.
3. Align the final offset to that maximum.

## Solution

```go
func AlignUp(n, a int) int {
	if a <= 1 {
		return n
	}
	return (n + a - 1) / a * a
}

func StructSize(sizes []int) int {
	offset, widest := 0, 0
	for _, s := range sizes {
		if s <= 0 {
			continue
		}
		offset = AlignUp(offset, s) + s
		widest = max(widest, s)
	}
	return AlignUp(offset, widest)
}
```

## Walkthrough

For `[1 8 2 4]`: the bool sits at 0, the `int64` must wait until 8, the `int16` lands at 16, the `int32` needs a multiple of 4 so it goes to 20 — total 24, versus 16 for the same fields ordered widest first.

## Pitfalls

- Forgetting the trailing padding, which underestimates every struct with mixed widths.
- Aligning to the *previous* field instead of the one being placed.
- Dividing by a zero alignment when a field size slips through the guard.
