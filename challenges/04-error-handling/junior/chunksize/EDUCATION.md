# Even Chunks

## Intuition

Integer division rounds toward zero, so splitting 10 into 3 parts of 3 silently drops a record. Rounding up is the arithmetic that preserves the whole.

## Approach

1. Reject `parts <= 0`.
2. Reject `total < 0`.
3. Return `(total + parts - 1) / parts`.

## Solution

```go
if parts <= 0 {
	return 0, ErrBadParts
}
if total < 0 {
	return 0, ErrNegativeTotal
}
return (total + parts - 1) / parts, nil
```

## Walkthrough

`(10 + 3 - 1) / 3` is `12/3 == 4`; `(9 + 3 - 1) / 3` is `11/3 == 3`, so exact divisions are unaffected.

## Pitfalls

- Using plain `total / parts`, losing the remainder.
- Guarding `parts == 0` only, letting a negative count through.
- Reordering the guards so a zero `parts` reaches the division.
