# Arrays Are Values, Slices Are Not

## Intuition

`[4]int` is a value like an `int` is. The function receives its own copy, so mutating it is private, and returning it copies again. Nothing is shared, and nothing is allocated.

## Approach

1. Range over the parameter and increment each element.
2. Return the parameter.

## Solution

```go
// Bump returns a copy of a with every element increased by one.
//
// The caller's array must not change: an array parameter is a value, and
// the returned array is a separate value again.
//
// Examples:
//
// 	Bump([4]int{1, 2, 3, 4}) => [4]int{2, 3, 4, 5}
func Bump(a [4]int) [4]int {
	for i := range a {
		a[i]++
	}
	return a
}
```

## Walkthrough

`Bump(a)` copies four ints onto the stack, increments them, and copies four ints back out. `a` in the caller is untouched throughout.

## Pitfalls

- Expecting a mutation to reach the caller — take `*[4]int` or a slice for that.
- Allocating a fresh array to fill; the parameter already is one.
