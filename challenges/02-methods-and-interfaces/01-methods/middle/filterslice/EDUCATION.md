# Filtering Slices

## Intuition

Functional-style methods on slices return new slices rather than mutating.
`var res IntList` creates a nil slice, which `append` handles gracefully.
`reflect.DeepEqual` treats nil and empty interchangeably for the test here.

## Approach

1. `var res IntList` (or `res := make(IntList, 0)`).
2. Loop over `l`, append if `v%2 == 0`.
3. Return `res`.

## Solution

```go
func (l IntList) FilterEvens() IntList {
	res := make(IntList, 0)
	for _, v := range l {
		if v%2 == 0 {
			res = append(res, v)
		}
	}
	return res
}
```

## Walkthrough

For `{1, 2, 3, 4}`:
- 1 is odd.
- 2 is even → append.
- 3 is odd.
- 4 is even → append.
- Return `{2, 4}`.

## Pitfalls

- Doing it in-place by removing elements is harder and breaks the test's
  immutability check.
- `append` to a nil `IntList` correctly produces an `IntList`.
