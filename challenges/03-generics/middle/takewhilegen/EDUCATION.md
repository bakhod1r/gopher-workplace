# Take While

## Intuition

The difference from `Filter` is one keyword, and it changes the meaning entirely: this returns a prefix, not a subset.

## Approach

1. Range over `s`.
2. Break at the first element `pred` rejects.
3. Append the rest.

## Solution

```go
func TakeWhile[T any](s []T, pred func(T) bool) []T {
	out := make([]T, 0, len(s))
	for _, v := range s {
		if !pred(v) {
			break
		}
		out = append(out, v)
	}
	return out
}
```

## Walkthrough

`TakeWhile([]int{2,4,5,6}, isEven)` stops at `5`, so the trailing `6` is not collected.

## Pitfalls

- Using `continue`, which turns the function into `Filter`.
- Returning nil when the first element already fails.
- Returning a sub-slice of the input instead of a fresh slice.
