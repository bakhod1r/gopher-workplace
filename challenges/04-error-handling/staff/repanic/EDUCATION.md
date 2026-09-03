# Re-panic Selectively

## Intuition

Recovery is a policy decision. Absorbing an application panic keeps a service up; absorbing a nil dereference hides memory-level corruption behind a 500 and lets the process keep serving.

## Approach

1. Defer a closure that recovers.
2. Re-panic with the same value when it satisfies `runtime.Error`.
3. Otherwise format it into the named error.

## Solution

```go
defer func() {
	r := recover()
	if r == nil {
		return
	}
	if _, ok := r.(runtime.Error); ok {
		panic(r)
	}
	err = fmt.Errorf("panic: %v", r)
}()
f()
return nil
```

## Walkthrough

The index panic's payload implements `runtime.Error`, so it is re-panicked untouched and the test's own recover sees the original value.

## Pitfalls

- Re-panicking with a new value, discarding the original payload and its type.
- Absorbing every panic, including runtime faults.
- Re-panicking outside the deferred function, where the value is no longer available.
