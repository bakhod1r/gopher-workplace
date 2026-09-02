# Longest Text

## Intuition

Ordering by length is different from ordering by value, so this needs `len`, not `<`. The constraint must therefore promise a type set where `len` is defined.

## Approach

1. Return `zero, false` for an empty slice.
2. Seed `best` from `s[0]`.
3. Replace `best` only when a strictly longer element appears.

## Solution

```go
func Longest[T Text](s []T) (T, bool) {
	if len(s) == 0 {
		var zero T
		return zero, false
	}
	best := s[0]
	for _, v := range s[1:] {
		if len(v) > len(best) {
			best = v
		}
	}
	return best, true
}
```

## Walkthrough

`Longest([]Label{"xx", "yy"})` never replaces `best`, because `len("yy") > len("xx")` is false — the tie keeps the earlier element.

## Pitfalls

- Using `>=`, which lets a later tie win.
- Constraining to `cmp.Ordered` and comparing values instead of lengths.
- Writing `string` instead of `~string`, which rejects `Label`.
