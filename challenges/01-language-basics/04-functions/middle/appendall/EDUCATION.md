# Forwarding variadic arguments

## Intuition

`extra...` spreads a slice into another variadic/append call; appending to a caller's slice can clobber it if capacity allows, so copy when purity matters.

## Approach

1. Copy `base` into a fresh slice so the caller's slice is not aliased.
2. Append the variadic `extra`.

## Solution

```go
func Concat(base []int, extra ...int) []int {
	out := append([]int(nil), base...)
	return append(out, extra...)
}
```

## Walkthrough

Copying `base` first means appending never overwrites the caller's `[1 2]`; the result is `[1 2 3 4]`.

## Pitfalls

- `append(base, extra...)` can mutate base's backing array if it has spare cap.
- Copy into a nil-based slice first for a guaranteed-fresh result.
