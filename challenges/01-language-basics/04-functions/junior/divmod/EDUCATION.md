# Integer division and modulo

## Intuition

For integers `/` yields the truncated quotient and `%` the remainder; together they satisfy `a == (a/b)*b + a%b`.

## Approach

1. Integer division `a / b` gives the quotient.
2. `a % b` gives the remainder.
3. Return both.

## Solution

```go
func DivMod(a, b int) (q, r int) {
	return a / b, a % b
}
```

## Walkthrough

`DivMod(17, 5)`: `17/5` is 3, `17%5` is 2.

## Pitfalls

- Division truncates toward zero, so `-7/2 == -3` and `-7%2 == -1`.
- Dividing by zero panics; the task guarantees `b != 0`.
