# Two-Value Result

## Intuition

A generic type does not force every field to be generic. Here only the payload varies; the failure reason is a plain string, which is why `Fail` cannot infer `T`.

## Approach

1. `Ok`: set the value and the flag.
2. `Fail`: set only the reason.
3. `Unwrap`: return both fields.
4. `Reason`: empty string when successful.

## Solution

```go
func Ok[T any](v T) Result[T] {
	return Result[T]{value: v, ok: true}
}

func Fail[T any](reason string) Result[T] {
	return Result[T]{reason: reason}
}

func (r Result[T]) Unwrap() (T, bool) {
	return r.value, r.ok
}

func (r Result[T]) Reason() string {
	if r.ok {
		return ""
	}
	return r.reason
}
```

## Walkthrough

`Fail[int]("bad").Unwrap()` returns `0, false` — the zero value comes from the unset field.

## Pitfalls

- Calling `Fail("bad")` and expecting `T` to be inferred.
- Returning the stored reason even on success.
- Making `Unwrap` panic on failure instead of reporting the flag.
