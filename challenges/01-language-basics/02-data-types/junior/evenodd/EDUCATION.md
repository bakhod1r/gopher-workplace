# Remainder and its sign

## The idea

`%` gives the remainder of integer division. For parity, `n % 2` is 0 when `n`
is even. The reliable test is `n%2 == 0`.

## Why it matters

In Go, `%` takes the **sign of the dividend**: `-7 % 2 == -1`, not `1`. So the
naive "odd means `n%2 == 1`" is wrong for negatives — `-7 % 2` is `-1`. Testing
against 0 sidesteps the sign entirely.

## Watch out

- `a % b` sign follows `a` (the dividend), unlike Python where it follows `b`.
- `%` is only for integers; use `math.Mod` for floats.
- Division by zero panics; `%` by zero panics too.

## Try it yourself

```go
-7 % 2  // -1
7 % -2  // 1
-7 % 2 == 0 // false -> odd, correct
```
