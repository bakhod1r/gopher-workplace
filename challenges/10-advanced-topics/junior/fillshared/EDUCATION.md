# Write Through The Caller's Array

## Intuition

Passing a slice copies three words but shares one array. That is why mutating elements reaches the caller and reassigning the parameter does not.

## Approach

1. Range over the indices of `s`.
2. Assign `v` to each index.

## Solution

```go
// Fill sets every element of s to v.
//
// The parameter is a view onto the caller's array, so the writes must be
// visible to the caller. Nothing is allocated and nothing is returned.
//
// Examples:
//
// 	s := []int{1, 2}; Fill(s, 7) => s is [7 7]
func Fill(s []int, v int) {
	for i := range s {
		s[i] = v
	}
}
```

## Walkthrough

`Fill(s[1:3], 0)` receives a header with pointer `&s[1]` and length 2, so it writes indices 1 and 2 of the caller's array and leaves 0 and 3 alone.

## Pitfalls

- `s = make([]int, len(s))` — allocates and is invisible outside.
- Using `range s` with the value variable and assigning to it; that writes to the loop copy.
