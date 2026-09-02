# Flat Map

## Intuition

Fusing the map and the flatten avoids building the intermediate slice of slices, which is the whole reason this helper exists next to `Map`.

## Approach

1. Allocate `out` as `[]U`.
2. Splice each `f(v)` result in.
3. Return `out`.

## Solution

```go
func FlatMap[T, U any](s []T, f func(T) []U) []U {
	out := make([]U, 0, len(s))
	for _, v := range s {
		out = append(out, f(v)...)
	}
	return out
}
```

## Walkthrough

`FlatMap([]int{1,2}, dup)` splices `[1 1]` then `[2 2]`.

## Pitfalls

- Appending `f(v)` without `...`, which does not compile.
- Declaring one type parameter, forcing input and output element types to match.
- Skipping elements whose result is empty — they simply add nothing.
