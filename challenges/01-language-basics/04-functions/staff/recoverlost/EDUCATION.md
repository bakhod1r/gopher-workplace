# Setting named returns inside deferred recover

## Intuition

The panic-to-error idiom must write the NAMED result variable in the deferred closure; a local assignment is invisible to the caller.

## Approach

1. The recovered error must be written to the **named return** `err`.
2. The bug assigns to a local `e` and discards it; assign to `err` directly.

## Solution

```go
import "fmt"

func Safe(f func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()
	f()
	return nil
}
```

## Walkthrough

Building `e` and dropping it leaves `err` nil despite the panic. Writing `err = fmt.Errorf(...)` surfaces the recovered failure to the caller.

## Pitfalls

- Deferred recover must set the named return, not a fresh local.
- `err = ...` in the closure body reaches the caller.
