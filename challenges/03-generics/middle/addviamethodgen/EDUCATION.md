# No Operator Overloading

## Intuition

Because operators come only from type sets, any user-defined arithmetic must arrive as methods, and the constraint is what makes them callable from generic code.

## Approach

1. Declare `var out T`.
2. Fold with `out.Plus(v)`.
3. Return the accumulator.

## Solution

```go
func SumAll[T Adder[T]](s []T) T {
	var out T
	for _, v := range s {
		out = out.Plus(v)
	}
	return out
}
```

## Walkthrough

`SumAll([]Money{{2},{3}})` starts from `Money{0}` and folds to `Money{5}`.

## Pitfalls

- Trying `out += v` on a struct type parameter.
- Assuming the zero value is a valid identity without checking the domain.
- Constraining to `~int` and giving up the currency safety the type provides.
