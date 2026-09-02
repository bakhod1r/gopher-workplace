# Fatal Assertion Helper

## Intuition

Choosing between the two is about what comes next: keep going when more information helps, stop when the following lines would panic.

## Approach

1. Mark as a helper.
2. Call `t.Fatalf` when the length differs.

## Solution

```go
func RequireLen[T any](t testing.TB, s []T, want int) {
	t.Helper()
	if len(s) != want {
		t.Fatalf("length %d, want %d (%v)", len(s), want, s)
	}
}
```

## Walkthrough

`RequireLen(t, []int{}, 1)` stops the test, so the `s[0]` on the next line never runs.

## Pitfalls

- Using `Errorf` and letting the test panic two lines later.
- Constraining `T` unnecessarily.
- Calling `Fatalf` from a helper goroutine, where it does not stop the test.
