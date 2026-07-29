# recover requires a deferred context

## Intuition

A direct `recover()` call returns nil; only a `recover()` inside a function invoked by `defer` during a panic captures the panic value.

## Approach

1. `recover` only works inside a **deferred** function during unwinding.
2. The bug calls `recover()` before `f()` runs, so it always sees nil.
3. Wrap it in `defer func(){ ... }()` before calling `f`.

## Solution

```go
func Guard(f func()) (ok bool) {
	defer func() {
		if r := recover(); r != nil {
			ok = true
		}
	}()
	f()
	return ok
}
```

## Walkthrough

Calling `recover()` up front returns nil and the panic from `f()` is never caught. A deferred recover fires while the stack unwinds, setting `ok = true`.

## Pitfalls

- `recover` outside a deferred call always returns nil.
- Put `f()` after scheduling the deferred recover, not before it.
