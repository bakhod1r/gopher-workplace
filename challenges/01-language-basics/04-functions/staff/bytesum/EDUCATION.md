# Choosing an accumulator type

## Intuition

Summation can exceed the element type's range; the accumulator must be sized for the aggregate, or unsigned/overflow wraps the result silently.

## Approach

1. A `uint8` accumulator overflows past 255.
2. Accumulate into an `int`, converting each byte: `total += int(b)`.

## Solution

```go
func Sum(bs []byte) int {
	var total int
	for _, b := range bs {
		total += int(b)
	}
	return total
}
```

## Walkthrough

`uint8` wraps: 200+100 already overflows. Summing into an `int` keeps the full total 350.

## Pitfalls

- `uint8 += ...` wraps at 256; the sum of small bytes can still overflow.
- Accumulate in `int` (or `uint64`) and convert the element, not the total.
