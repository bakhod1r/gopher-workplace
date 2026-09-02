# Error Message

## Intuition

`error` is an interface with a single method, `Error() string`. Reading the message means calling that method — and a nil interface has no method to call.

## Approach

1. Return `""` when `err` is nil.
2. Otherwise return `err.Error()`.

## Solution

```go
if err == nil {
	return ""
}
return err.Error()
```

## Walkthrough

For `nil` the guard returns `""` immediately. For `ErrTimeout` the guard is skipped and `Error()` yields `"timeout"`.

## Pitfalls

- Calling `err.Error()` before the nil check — nil pointer dereference panic.
- Using `fmt.Sprint(err)`, which renders nil as `"<nil>"`.
- Returning a placeholder like `"none"` instead of the empty string.
