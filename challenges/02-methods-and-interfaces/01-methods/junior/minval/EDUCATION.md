# Finding the Minimum

## Intuition

The minimum is found by starting with the largest possible value (`+Inf`) and
replacing it whenever a smaller element is found. For empty input, `+Inf` is
the natural identity value for the min operation.

## Approach

1. Start `result = math.Inf(1)`.
2. Range over values; if `v < result`, update.
3. Return result.

## Solution

```go
func (s Stats) Min() float64 {
	result := math.Inf(1)
	for _, v := range s.Values {
		if v < result {
			result = v
		}
	}
	return result
}
```

## Walkthrough

For `[]float64{3, 1, 2}`:
- result = +Inf → 3 < +Inf → result = 3.
- 1 < 3 → result = 1.
- 2 < 1 → false.
- Returns 1.

## Pitfalls

- Initializing with `0` instead of `+Inf` — fails for all-positive datasets
  (would incorrectly return 0).
- Initializing with `Values[0]` — panics on empty slice.
