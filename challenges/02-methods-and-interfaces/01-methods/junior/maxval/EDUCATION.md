# Finding the Maximum

## Intuition

Max mirrors Min: start with `-Inf` and replace whenever a larger value appears.

## Approach

1. Start `result = math.Inf(-1)`.
2. Range; if `v > result`, update.
3. Return.

## Solution

```go
func (s Stats) Max() float64 {
	result := math.Inf(-1)
	for _, v := range s.Values {
		if v > result {
			result = v
		}
	}
	return result
}
```

## Walkthrough

For `[]float64{3, 1, 2}`:
- result = -Inf → 3 > -Inf → result = 3.
- 1 > 3 → false.
- 2 > 3 → false.
- Returns 3.

## Pitfalls

- Using `math.Inf(1)` instead of `math.Inf(-1)` — starts at +Inf, nothing is
  larger, always returns +Inf.
