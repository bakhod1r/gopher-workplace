# The Detour Through Interface Values

## Intuition

An interface value is a type word plus a data word, and the data word must be a pointer. Putting an int behind `any` therefore means allocating somewhere for the int to live -- once per element, for a generality nobody asked for.

## Approach

1. Delete the boxing pass.
2. Range `vals` directly and accumulate into an int64.

## Solution

```go
// Total sums vals.
//
// Passing the values through []any boxes every element: an interface value
// needs a word to point at, so each int gets a heap home it never needed.
//
// Examples:
//
// 	Total([]int{1, 2, 3}) => 6
func Total(vals []int) int64 {
	var total int64
	for _, v := range vals {
		total += int64(v)
	}
	return total
}
```

## Walkthrough

For 64 values above 255, the boxed version allocates the `[]any` plus one box per element -- 65 allocations. The direct loop allocates nothing and keeps the accumulator in a register.

## Pitfalls

- Testing with small values, where the runtime's 0-255 cache hides most of the cost.
- Accumulating in `int`, which is fine on 64-bit and not on every target.
