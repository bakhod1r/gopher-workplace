# Must

## Intuition

The wrapper turns a checked API into an unchecked one at the exact points where a failure means the program is misconfigured — and keeps the value's static type intact.

## Approach

1. Panic when `ok` is false.
2. Otherwise return `v`.

## Solution

```go
func Must[T any](v T, ok bool) T {
	if !ok {
		panic("Must: value not available")
	}
	return v
}

func Lookup[K comparable, V any](m map[K]V, k K) (V, bool) {
	v, ok := m[k]
	return v, ok
}
```

## Walkthrough

`Must(Lookup(m, "missing"))` receives `ok == false` and panics rather than silently returning a zero value.

## Pitfalls

- Using `Must` on request paths, turning recoverable errors into crashes.
- Returning the zero value instead of panicking, which defeats the purpose.
- Panicking with a message that does not say which helper failed.
