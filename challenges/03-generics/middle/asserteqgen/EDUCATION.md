# Slice Assertion Helper

## Intuition

A helper that reports the first difference with its index is far more useful than a dumped diff, and it costs nothing extra to write.

## Approach

1. Mark as a helper.
2. Report and bail on a length mismatch.
3. Report and bail at the first differing index.
4. Return `true` otherwise.

## Solution

```go
func EqualSlice[T comparable](t testing.TB, got, want []T) bool {
	t.Helper()
	if len(got) != len(want) {
		t.Errorf("length %d, want %d (got %v, want %v)", len(got), len(want), got, want)
		return false
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("index %d: got %v, want %v", i, got[i], want[i])
			return false
		}
	}
	return true
}
```

## Walkthrough

Comparing `[1 2]` with `[1 3]` reports `index 1: got 2, want 3`.

## Pitfalls

- Taking `*testing.T` and excluding benchmarks.
- Reporting every mismatched element in a long slice.
- Returning nothing, so callers cannot skip dependent assertions.
