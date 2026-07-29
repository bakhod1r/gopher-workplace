# Reversing a slice

## Intuition

Swap symmetric elements from the ends inward, stopping at the middle:

```go
for i, j := 0, len(xs)-1; i < j; i, j = i+1, j-1 { xs[i], xs[j] = xs[j], xs[i] }
```

## Approach

1. Set two indices i=0 and j=len(nums)-1.
2. While i < j, swap nums[i] and nums[j], then move i up and j down.
3. Mutate the shared backing array in place; return nothing.

## Solution

```go
func Reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
```

## Walkthrough

Reverse([1,2,3]): i=0,j=2 swap -> [3,2,1]; i=1,j=1 loop stops. Caller sees [3,2,1].

## Pitfalls

- In-place reversal mutates the caller's slice (shared backing array).
- Multiple assignment swaps without a temp.
- `slices.Reverse` does this generically (Go 1.21+).
