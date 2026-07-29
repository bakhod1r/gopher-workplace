# Panic and recover

## Intuition

A panic unwinds the stack running deferred calls; `recover` in one of them captures the panic value and resumes normal return.

## Approach

1. A deferred `recover()` catches a panic during unwinding.
2. Set `recovered = true` when `recover()` is non-nil.

## Solution

```go
func SafeInvoke(f func()) (recovered bool) {
	defer func() {
		if r := recover(); r != nil {
			recovered = true
		}
	}()
	f()
	return
}
```

## Walkthrough

A panicking `f` triggers the deferred recover, setting `recovered = true`; a normal `f` leaves it false.

## Pitfalls

- `recover` only works when called directly inside a deferred function.
- Don't use panic/recover for ordinary control flow.
