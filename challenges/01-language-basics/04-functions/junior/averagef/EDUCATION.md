# Guarding aggregation

## Intuition

An average needs at least one sample; returning a validity flag lets callers avoid a division-by-zero or a misleading 0.

## Approach

1. Guard the empty case (`ok == false`).
2. Sum then divide by `float64(len(nums))`.

## Solution

```go
func Average(nums ...float64) (float64, bool) {
	if len(nums) == 0 {
		return 0, false
	}
	sum := 0.0
	for _, v := range nums {
		sum += v
	}
	return sum / float64(len(nums)), true
}
```

## Walkthrough

`Average(2,4,6)`: sum 12 over 3 → 4.0, true; no args returns 0, false.

## Pitfalls

- Dividing by `len(nums)` as an int truncates; convert to `float64`.
- Guard the empty slice before the division.
