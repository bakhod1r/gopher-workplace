# Isolate A Callback

## Intuition

A host that calls foreign code owns the blast radius. Recovering at the call site turns "the process died" into "that one plugin failed".

## Approach

1. Guard against a nil `f`.
2. Defer a closure that recovers and formats the payload.
3. Return `f()` normally.

## Solution

```go
if f == nil {
	return ErrNilFunc
}
defer func() {
	if r := recover(); r != nil {
		err = fmt.Errorf("panic: %v", r)
	}
}()
return f()
```

## Walkthrough

A panic with the int `7` renders as `"panic: 7"` — `%v` handles any payload type, so the host does not need to know what plugins throw.

## Pitfalls

- Omitting the nil guard, so `SafeCall(nil)` panics inside your own code.
- Using an unnamed result, which discards the error set during recovery.
- Recovering panics that should crash the process, such as programmer bugs in the host itself.
