# Length Of

## Intuition

You could not index or append here: `append` is undefined for `string`. `len` survives because it is defined for both members, which is the rule for every operation on a type parameter.

## Approach

1. Return `len(v)`.

## Solution

```go
func Length[T Sized](v T) int {
	return len(v)
}
```

## Walkthrough

`Length([]byte{1, 2})` instantiates `T = []byte` and returns `2`; `Length("abc")` instantiates `T = string` and returns `3`.

## Pitfalls

- Trying `append(v, ...)` inside — undefined for `string`.
- Converting to `string` first, which allocates for `[]byte` inputs.
- Counting runes: `len` on a string counts bytes.
