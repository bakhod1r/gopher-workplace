# if init statement

## Intuition

The init clause runs once and binds a variable scoped to the whole if/else chain, keeping helpers local.

## Approach

1. Compute `r := n % 3` in the if init clause.
2. Branch on `r` being 0, 1, or 2.
3. `r` is scoped to the if/else chain.

## Solution

```go
func Bucket(n int) string {
	if r := n % 3; r == 0 {
		return "zero"
	} else if r == 1 {
		return "one"
	} else {
		return "two"
	}
}
```

## Walkthrough

`Bucket(10)`: `r` is 1, so the middle branch returns "one".

## Pitfalls

- `r` is not visible after the if/else ends.
- The remainder is computed a single time.
