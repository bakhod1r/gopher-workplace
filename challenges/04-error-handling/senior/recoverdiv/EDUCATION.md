# Panic To Error

## Intuition

`recover` stops a panic only while a deferred function is running, and the recovered function still returns. Named results are what let the deferred code decide what it returns.

## Approach

1. Defer a closure that calls `recover`.
2. Set `err` and zero `result` when the recovered value is non-nil.
3. Perform the division normally.

## Solution

```go
defer func() {
	if r := recover(); r != nil {
		result = 0
		err = ErrPanic
	}
}()
return a / b, nil
```

## Walkthrough

`SafeDivide(1, 0)` panics inside the division; the deferred closure recovers, overwrites both named results, and the function returns `0, ErrPanic`.

## Pitfalls

- Calling `recover()` directly in the function body, where it always returns nil.
- Using unnamed results, so the recovered path returns the zero values with a nil error.
- Recovering without setting an error, silently reporting success.
