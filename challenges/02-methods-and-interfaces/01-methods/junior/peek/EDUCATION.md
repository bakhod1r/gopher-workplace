# Read-Only Methods

## Intuition

`Peek` only looks at data — it never changes the stack. A **value receiver** is
perfect: the method works on a copy, guaranteeing the original is untouched.

## Approach

1. If empty, return `0, false`.
2. Return the last element and `true`.

## Solution

```go
func (s Stack) Peek() (int, bool) {
	if len(s.Items) == 0 {
		return 0, false
	}
	return s.Items[len(s.Items)-1], true
}
```

## Walkthrough

For `Stack{Items: []int{1, 2, 3}}`:
- `len` = 3, not empty.
- `s.Items[2]` = 3.
- Returns `(3, true)`.

## Pitfalls

- Using a pointer receiver works but signals "may mutate" to the reader.
- Forgetting the empty check → panic on index out of range.
