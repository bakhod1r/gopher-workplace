# Typed Panic Payload

## Intuition

`recover()` hands back `any`, so the boundary decides how much structure survives. Passing an error payload through untouched keeps `errors.Is` working across the panic.

## Approach

1. Defer a closure that recovers.
2. Type-assert the payload to `error` and assign it directly.
3. Otherwise format it with `%v`.

## Solution

```go
defer func() {
	r := recover()
	if r == nil {
		return
	}
	if e, ok := r.(error); ok {
		err = e
		return
	}
	err = fmt.Errorf("panic: %v", r)
}()
f()
return nil
```

## Walkthrough

Panicking with `ErrStop` returns the identical value, so both `==` and `errors.Is` succeed; panicking with `42` produces a fresh formatted error.

## Pitfalls

- Formatting every payload, which destroys error identity.
- Asserting to a concrete type instead of the `error` interface.
- Returning a non-nil error when nothing panicked.
