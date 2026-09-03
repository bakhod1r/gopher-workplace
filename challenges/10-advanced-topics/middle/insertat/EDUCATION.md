# Make Room In The Middle

## Intuition

Insertion is a shift, and a shift needs somewhere to shift into. Appending a placeholder first turns the problem into one `copy` and one assignment.

## Approach

1. Clamp `i`.
2. `append` a zero to extend the length.
3. `copy(s[i+1:], s[i:])` to shift the tail up.
4. Write `v` at `i` and return.

## Solution

```go
// InsertAt inserts v at index i, shifting the rest up, and returns the
// extended slice.
//
// i is clamped into [0, len(s)]. The shift must not lose the element it
// overwrites.
//
// Examples:
//
// 	InsertAt([]int{1, 3}, 1, 2) => []int{1, 2, 3}
func InsertAt(s []int, i, v int) []int {
	if i < 0 {
		i = 0
	}
	if i > len(s) {
		i = len(s)
	}
	s = append(s, 0)
	copy(s[i+1:], s[i:])
	s[i] = v
	return s
}
```

## Walkthrough

For [1 3] and i = 1: appending gives [1 3 0]; the copy shifts 3 up to index 2 giving [1 3 3]; writing 2 at index 1 gives [1 2 3].

## Pitfalls

- Shifting before extending, which writes past the length.
- `append(s[:i], append([]int{v}, s[i:]...)...)` — correct, and it allocates twice.
