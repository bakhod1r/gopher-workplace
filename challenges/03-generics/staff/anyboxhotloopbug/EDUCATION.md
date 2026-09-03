# The Sum That Allocates Per Element

## Intuition

The generic body already knows `T` at compile time, so it can add straight into an accumulator. Detouring every element through an `any` slice forces each value onto the heap and then pays for a type assertion to get it back.

## Approach

1. Declare an accumulator of type `T`.
2. Add each element to it directly.
3. Return the accumulator.

## Solution

```go
func Sum[T Number](s []T) T {
	var total T
	for _, v := range s {
		total += v
	}
	return total
}
```

## Walkthrough

For three million `int` samples the boxed version performs three million heap allocations and copies; the direct version performs none.

## Pitfalls

- Assuming the escape analyser will optimise the boxing away — it cannot, because the values outlive the loop iteration.
- Reaching for `any` to "simplify" a generic function, which reintroduces exactly the cost generics exist to remove.
