# Stack At The Panic

## Intuition

By the time a recovered function returns, the frames that panicked are gone. The stack has to be captured inside the deferred call, which is the last moment it still exists.

## Approach

1. Defer a closure that recovers.
2. Fill a fixed buffer with `runtime.Stack(buf, false)`.
3. Format the payload and the captured bytes into the named error.

## Solution

```go
defer func() {
	if r := recover(); r != nil {
		buf := make([]byte, 4096)
		n := runtime.Stack(buf, false)
		err = fmt.Errorf("panic: %v\n%s", r, buf[:n])
	}
}()
f()
return nil
```

## Walkthrough

The stack captured inside the deferred function still contains this package's frames, which is what the test checks for.

## Pitfalls

- Capturing the stack after the recovered function returns, where the frames are already unwound.
- Writing the whole buffer instead of `buf[:n]`, appending kilobytes of zero bytes.
- Passing `true` to `runtime.Stack`, dumping every goroutine in the process.
