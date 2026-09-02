# Product

## Intuition

An untyped constant like `1` is assignable to a type parameter as long as it is representable in every type of the set. That is what makes the identity element expressible here.

## Approach

1. Start `out` at `1`.
2. Multiply each element in.
3. Return `out`.

## Solution

```go
func Product[T Number](s []T) T {
	var out T = 1
	for _, v := range s {
		out *= v
	}
	return out
}
```

## Walkthrough

`Product([]int{})` never enters the loop and returns the initial `1`.

## Pitfalls

- Starting at `0`, which makes every product zero.
- Using `var out T` and hoping the zero value works — it does not for multiplication.
- Returning `0` for an empty slice.
