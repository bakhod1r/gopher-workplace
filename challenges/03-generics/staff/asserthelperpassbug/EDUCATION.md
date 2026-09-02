# The Assertion That Always Passes

## Intuition

The mismatch is logged, but control falls out of the `if` and hits the unconditional `return true`. Every caller that branches on the result therefore keeps going into state the helper just proved is wrong.

## Approach

1. Compare the two values.
2. On a mismatch, report it and return `false`.
3. Otherwise return `true`.

## Solution

```go
func AssertEqual[T comparable](t Reporter, got, want T) bool {
	if got != want {
		t.Errorf("got %v, want %v", got, want)
		return false
	}
	return true
}
```

## Walkthrough

`AssertEqual(t, 1, 2)` logs "got 1, want 2" and then answers `true`, so the caller proceeds as though the value were valid.

## Pitfalls

- Returning the comparison but forgetting to report it — the opposite failure, a silent pass.
- Reporting with a plain `Errorf` where the caller genuinely cannot continue; a guard result only helps if it is honest.
