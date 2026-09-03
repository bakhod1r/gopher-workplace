# A Function That Touches No Heap

## Intuition

Nothing about summing needs memory that outlives the call. Keep the running total in a plain local and the compiler keeps it in a register.

## Approach

1. Declare a local `total`.
2. Range and add.
3. Return `total`.

## Solution

```go
// Sum returns the total of s.
//
// The function must not allocate: every value it needs fits in a local
// variable, and the input is only read.
//
// Examples:
//
// 	Sum([]int{1, 2, 3}) => 6
func Sum(s []int) int {
	total := 0
	for _, v := range s {
		total += v
	}
	return total
}
```

## Walkthrough

Summing 512 ints touches 4 KB of the caller's array and one stack slot. No allocation is ever requested.

## Pitfalls

- `total := new(int)` — now the pointer may escape and the count is 1.
- Building an intermediate slice or using `fmt` anywhere in a hot helper.
