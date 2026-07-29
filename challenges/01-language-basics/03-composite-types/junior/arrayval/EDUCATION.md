# Arrays are values

## Intuition

A fixed-size array (`[N]T`) is a value type: its length is part of its type, and
assigning or passing it **copies** every element. This differs from a slice,
which shares a backing array.

## Approach

1. The parameter a is already a copy — Go passes arrays by value.
2. Assign a[0] = v to mutate the local copy.
3. Return a. The caller's original is unaffected.

## Solution

```go
func SetFirst(a [3]int, v int) [3]int {
	a[0] = v
	return a
}
```

## Walkthrough

SetFirst([3]int{1,2,3}, 9): the function receives its own copy {1,2,3}, sets index 0 to 9 giving {9,2,3}, and returns it; the caller still holds {1,2,3}.

## Pitfalls

- `[3]int` and `[4]int` are different types.
- Arrays are comparable with `==` (element-wise); slices are not.
- Large arrays copied by value are expensive.
