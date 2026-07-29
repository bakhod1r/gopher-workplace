# Slice expressions

## Intuition

`xs[a:b]` creates a slice sharing the same backing array, covering indices
`[a, b)`. A tail of `n` is `xs[len(xs)-n:]` — but only after clamping `n`:

```go
if n > len(xs) { n = len(xs) }
if n < 0 { n = 0 }
return xs[len(xs)-n:]
```

## Approach

1. If n <= 0, return an empty slice.
2. If n >= len(xs), return all of xs.
3. Otherwise return the slice expression xs[len(xs)-n:].

## Solution

```go
func Last(xs []int, n int) []int {
	if n <= 0 {
		return []int{}
	}
	if n >= len(xs) {
		return xs
	}
	return xs[len(xs)-n:]
}
```

## Walkthrough

Last([1,2,3,4,5],2): n=2 is between 0 and 5, so return xs[5-2:] = xs[3:] = [4,5].

## Pitfalls

- `xs[len(xs)-n:]` panics if `n > len(xs)` — clamp first.
- The result **shares** memory with `xs`; appending to it may overwrite.
- `xs[i:]` is shorthand for `xs[i:len(xs)]`.
