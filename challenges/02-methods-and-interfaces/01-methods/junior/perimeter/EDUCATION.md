# Computing with Value Receivers

## Intuition

Methods that only compute and return a value — never mutating the receiver —
should use a value receiver. The caller's struct stays untouched.

## Approach

1. Sum the two dimensions.
2. Multiply by 2.
3. Return.

## Solution

```go
func (r Rect) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}
```

## Walkthrough

For `Rect{3, 4}`:
- `r.Width + r.Height` = 7.
- `2 * 7` = 14.

## Pitfalls

- A common mistake is `2*r.Width + r.Height` (missing parentheses) which gives
  `2*3 + 4 = 10` instead of 14.
- The formula is correct even when one or both dimensions are zero.
