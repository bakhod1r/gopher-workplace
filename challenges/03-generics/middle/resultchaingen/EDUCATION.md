# Chaining Results

## Intuition

This is the same restriction that forced `Map` to be a function: the moment the result type changes, the operation moves out of the method set.

## Approach

1. `Then`: rebuild the failure as `Result[U]` when `r` failed, otherwise call `f`.
2. The constructors and `Unwrap` are straightforward field work.

## Solution

```go
func Ok[T any](v T) Result[T] {
	return Result[T]{value: v, ok: true}
}

func Fail[T any](reason string) Result[T] {
	return Result[T]{reason: reason}
}

func Then[T, U any](r Result[T], f func(T) Result[U]) Result[U] {
	if !r.ok {
		return Fail[U](r.reason)
	}
	return f(r.value)
}

func (r Result[T]) Unwrap() (T, bool) {
	return r.value, r.ok
}
```

## Walkthrough

`Then(Fail[int]("bad"), double)` never calls `double` and returns a `Result[string]` carrying `"bad"`.

## Pitfalls

- Trying to declare `func (r Result[T]) Then[U any](...)`.
- Calling `f` before checking the flag.
- Dropping the reason when converting the failure to `Result[U]`.
