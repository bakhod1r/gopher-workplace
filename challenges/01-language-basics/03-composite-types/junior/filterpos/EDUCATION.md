# Building slices with append

## Intuition

`append` adds to a slice and returns a (possibly relocated) new header — always
assign the result back:

```go
out := []int{}
for _, x := range xs {
	if x > 0 { out = append(out, x) }
}
```

## Approach

1. Start with an empty (non-nil) result slice.
2. Range xs; append only elements with x > 0.
3. Return result, preserving original order.

## Solution

```go
func Positives(xs []int) []int {
	result := []int{}
	for _, x := range xs {
		if x > 0 {
			result = append(result, x)
		}
	}
	return result
}
```

## Walkthrough

Positives([1,-2,3,0,4]): 1>0 keep; -2 drop; 3 keep; 0 drop; 4 keep -> [1,3,4].

## Pitfalls

- `append(s, x)` may reallocate; the returned slice is authoritative.
- `[]int{}` is non-nil length 0; `var out []int` is nil — pick per the test/spec.
- Pre-sizing with `make([]int, 0, len(xs))` avoids regrowth in hot paths.
