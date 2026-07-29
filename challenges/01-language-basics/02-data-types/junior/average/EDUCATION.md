# Integer vs float average

## Intuition

The mean of integers is a *ratio*, so the division must happen in floating point.
Summing in `int` and dividing by the count as ints truncates the fraction:

```go
float64(sum) / float64(n) // 7/2 -> 3.5, not 3
```

## Approach

1. If the slice is empty or nil, return 0.
2. Sum the ints into an int accumulator.
3. Convert both sum and len to float64 before dividing so the fraction survives.

## Solution

```go
func Average(nums []int) float64 {
	if len(nums) == 0 {
		return 0
	}
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return float64(sum) / float64(len(nums))
}
```

## Walkthrough

Average([]int{1,2,4}): sum=7, len=3, float64(7)/float64(3)=2.3333333333333335.

## Pitfalls

- Convert to `float64` before dividing, not after.
- Guard `n == 0` to avoid a divide-by-zero panic.
- Summing many ints can overflow; widen the accumulator if needed.
