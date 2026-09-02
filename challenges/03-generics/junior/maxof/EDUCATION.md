# Max Of Two

## Intuition

`comparable` and `cmp.Ordered` are different promises: equality versus ordering. Structs are comparable but not ordered, which is why the two constraints exist.

## Approach

1. Return `a` when `a > b`, else `b`.

## Solution

```go
func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
```

## Walkthrough

`Max("a", "b")` instantiates `T = string`; `>` on strings compares byte by byte, so `"b"` wins.

## Pitfalls

- Using `comparable`, which does not allow `>`.
- Writing your own union and forgetting a type the caller needs.
- Returning `a` on ties when the doc says `b` — harmless here, but read the spec.
