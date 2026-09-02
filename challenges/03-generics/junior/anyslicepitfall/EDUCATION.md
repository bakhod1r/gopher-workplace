# The []any Trap

## Intuition

An `any` holds a type word plus a data word, so a `[]any` is physically wider than a `[]int`. That is why the conversion has to be element by element — and why generic code that avoids it is faster.

## Approach

1. Allocate `[]any` with capacity `len(s)`.
2. Append each element, which boxes it.
3. Return the result.

## Solution

```go
func ToAny[T any](s []T) []any {
	out := make([]any, 0, len(s))
	for _, v := range s {
		out = append(out, v)
	}
	return out
}

func SumInts(s []int) int {
	total := 0
	for _, v := range s {
		total += v
	}
	return total
}
```

## Walkthrough

`ToAny([]int{1, 2})` allocates a new slice of interface values; the original `[]int` is untouched and still packed.

## Pitfalls

- Trying `[]any(s)` or `any(s).([]any)`, neither of which compiles or works.
- Reaching for `[]any` in new APIs where a type parameter would do.
- Assuming the conversion is free — it is O(n) with boxing.
