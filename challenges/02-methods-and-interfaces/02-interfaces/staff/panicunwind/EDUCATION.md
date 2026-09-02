# Panic And Unwind

## Intuition

`recover` is only meaningful in a function deferred directly by the panicking frame, and it can only change the caller's result if that result is named. Everything else about panics — LIFO defers, unwinding one frame at a time — follows from that.

## Approach

1. Give `SafeRun` a named `err` result.
2. Defer a closure that calls `recover` and assigns a wrapped error on a non-nil value.
3. Return `t.Run()` normally; the defer only overwrites `err` when a panic occurred.
4. `Order` defers three appends and panics; the last defer registered runs first.

## Solution

```go
func SafeRun(t Task) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()
	return t.Run()
}

func Order() (order []string) {
	defer func() {
		recover()
		order = append(order, "third")
	}()
	defer func() { order = append(order, "second") }()
	defer func() { order = append(order, "first") }()

	panic("unwind")
}
```

## Walkthrough

In `Order` the innermost-registered defer (`"first"`) runs first, then `"second"`, then the outermost one — which also recovers and appends `"third"`. Recovery in the outermost defer is what stops the panic from escaping.

## Pitfalls

- Calling `recover()` in a helper function rather than directly in the deferred closure — it returns nil and the panic continues.
- Using an unnamed result: the deferred function cannot change what the caller sees.
- Recovering across goroutines: a panic in a goroutine can only be recovered inside that goroutine.
