# An Append That Cannot Reach The Caller's Tail

## Intuition

`append` is allowed to write into the slice's spare capacity in place. When the slice is a prefix view of a longer array, that spare capacity is another view's live data.

## Approach

1. Cap the slice's capacity at its own length with a three-index slice.
2. Append to that — the reallocation is now forced whenever room would have been borrowed.

## Solution

```go
// Add returns s with v appended.
//
// The caller may be holding a longer slice over the same array. Appending
// must never overwrite elements past len(s); the result must get its own
// storage whenever that would happen.
//
// Examples:
//
// 	Add([]int{1, 2}, 3) => []int{1, 2, 3}
func Add(s []int, v int) []int {
	return append(s[:len(s):len(s)], v)
}
```

## Walkthrough

`b[:2]` has len 2, cap 4. Plain `append` writes 99 into `b[2]`. `b[:2:2]` has cap 2, so `append` allocates a new array and `b` is untouched.

## Pitfalls

- Copying unconditionally — correct, but it allocates even when the slice already owns its capacity.
- `s[:len(s)]` is not the fix; the third index is what matters.
