# Named return values

## Intuition

Declaring result names lets you assign them anywhere in the body and use a bare `return`; they default to zero values.

## Approach

1. Guard division by zero: a bare `return` yields the zero-valued named returns.
2. Otherwise set `q, ok = a/b, true`.

## Solution

```go
func SafeDiv(a, b int) (q int, ok bool) {
	if b == 0 {
		return
	}
	q, ok = a/b, true
	return
}
```

## Walkthrough

`SafeDiv(1, 0)` hits the guard and bare-returns `0, false`; a valid divisor returns the quotient and true.

## Pitfalls

- A bare `return` on the zero path yields (0,false) automatically.
- Overusing named returns hurts readability; reserve for defer/guard cases.
