# Underlying Types

## Intuition

Without the tilde, the constraint's type set holds exactly `int`, and `Celsius` — a different type — is rejected. `~int` widens the set to every type whose underlying type is `int`.

## Approach

1. Declare `var total T`.
2. Add each element.
3. Return the total.

## Solution

```go
func SumTemps[T IntLike](s []T) T {
	var total T
	for _, v := range s {
		total += v
	}
	return total
}
```

## Walkthrough

`SumTemps([]Celsius{1, 2})` instantiates `T = Celsius`; the result keeps the named type rather than decaying to `int`.

## Pitfalls

- Writing `int` instead of `~int` in the constraint, which rejects `Celsius`.
- Converting to `int` inside the loop and losing the named return type.
- Assuming `~int` also covers `int64` — it does not.
