# Pop From Stack

## Intuition

Removing from the end of a slice is a re-slice, not a copy. The only real hazard is the empty case, where the last index does not exist.

## Approach

1. Return the error when `len(s) == 0`.
2. Take `s[len(s)-1]` as the value.
3. Return `s[:len(s)-1]` as the remainder.

## Solution

```go
if len(s) == 0 {
	return nil, 0, ErrEmpty
}
return s[:len(s)-1], s[len(s)-1], nil
```

## Walkthrough

For `[]int{7}` the value is `s[0] == 7` and the remainder is `s[:0]`, an empty but non-nil slice.

## Pitfalls

- Indexing before the length check — `s[-1]` panics.
- Returning `s[1:]`, which pops from the front.
- Returning a partial result alongside `ErrEmpty`.
