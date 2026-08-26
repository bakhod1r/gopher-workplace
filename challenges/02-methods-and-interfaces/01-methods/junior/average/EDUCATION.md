# Average and Edge Cases

## Intuition

Average = total / count. The tricky part is avoiding division by zero when the
collection is empty.

## Approach

1. If `len(s.Values) == 0`, return 0.
2. Sum all values.
3. Divide by count.

## Solution

```go
func (s Stats) Average() float64 {
	n := len(s.Values)
	if n == 0 {
		return 0
	}
	total := 0.0
	for _, v := range s.Values {
		total += v
	}
	return total / float64(n)
}
```

## Walkthrough

For `[]float64{2, 4, 6}`:
- `n` = 3.
- total = 12.
- `12 / 3` = 4.

## Pitfalls

- Integer division: `len` returns `int`; convert to `float64` before dividing.
- Missing the empty guard → runtime panic: divide by zero with int conversion.
