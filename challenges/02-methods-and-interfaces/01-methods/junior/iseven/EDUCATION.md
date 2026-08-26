# Boolean Methods on Defined Types

## Intuition

Combining a defined type with a boolean method creates a self-describing API:
`taskID.IsEven()` is clearer than `taskID % 2 == 0` at call sites.

## Approach

1. Use modulo: `n % 2 == 0`.

## Solution

```go
func (n MyInt) IsEven() bool {
	return n%2 == 0
}
```

## Walkthrough

For `MyInt(4)`:
- `4 % 2` = 0.
- `0 == 0` is `true`.

## Pitfalls

- In Go, `%` preserves the sign: `(-3) % 2 == -1`, not `1`. But `(-3) % 2 == 0`
  is still `false`, so the check works correctly for negatives.
- Using bitwise `n&1 == 0` also works and is marginally faster.
