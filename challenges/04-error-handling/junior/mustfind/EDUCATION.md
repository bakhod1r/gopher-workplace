# Find Value

## Intuition

Index 0 is a perfectly good answer, so it cannot double as "not found". The error carries that distinction and the index stays a pure position.

## Approach

1. Range with the index.
2. Return `i, nil` on the first match.
3. Return `-1, ErrNotFound` after the loop.

## Solution

```go
for i, v := range s {
	if v == target {
		return i, nil
	}
}
return -1, ErrNotFound
```

## Walkthrough

For `[]int{5, 5, 5}` the first iteration already matches, so the answer is `0, nil` and later duplicates are ignored.

## Pitfalls

- Returning `0` for a missing value, which collides with a real match at index 0.
- Continuing the loop after a match and reporting the last index.
- Returning an index without an error, forcing callers to test for `-1`.
