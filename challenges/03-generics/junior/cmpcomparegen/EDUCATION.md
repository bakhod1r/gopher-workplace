# cmp.Compare

## Intuition

Chained comparators read as a list of priorities. Because `cmp.Compare` returns `0` for equal, "not decided yet" and "equal" are the same value, which makes the chain trivial.

## Approach

1. Compare the lengths; return that result when it is non-zero.
2. Otherwise return the comparison of the strings themselves.

## Solution

```go
func ByLengthThenName(a, b string) int {
	if c := cmp.Compare(len(a), len(b)); c != 0 {
		return c
	}
	return cmp.Compare(a, b)
}
```

## Walkthrough

`ByLengthThenName("bb", "aa")` sees equal lengths, then compares `"bb"` against `"aa"` and returns a positive number.

## Pitfalls

- Comparing the strings first, which ignores the primary key.
- Returning a `bool`.
- Subtracting lengths directly, which is fine for small values but hides the intent.
