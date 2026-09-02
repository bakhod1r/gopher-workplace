# Min Of Two

## Intuition

Both directions rely on the same promise from `cmp.Ordered`: the `<` and `>` operators are defined for every type in the set.

## Approach

1. Return `a` when `a < b`, else `b`.

## Solution

```go
func Min[T cmp.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}
```

## Walkthrough

`Min(3.5, 1.5)` compares `3.5 < 1.5` (false) and returns `b`, which is `1.5`.

## Pitfalls

- Copying `Max` and forgetting to flip the operator.
- Using `<=` and changing which argument wins on ties.
- Constraining to `Number`, which would reject strings.
