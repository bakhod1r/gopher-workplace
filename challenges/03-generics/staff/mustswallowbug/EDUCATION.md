# Must That Never Panics

## Intuition

Swallowing the error and handing back the zero value converts a fatal start-up problem into a plausible-looking empty value that propagates through the whole program.

## Approach

1. Check the error first.
2. Panic with it when it is non-nil.
3. Return the value otherwise.

## Solution

```go
func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
```

## Walkthrough

`Must(Load("missing.yaml"))` returns an empty config instead of stopping the process, so the failure surfaces hours later as "no upstreams configured".

## Pitfalls

- Panicking with a fresh string instead of the error, which discards the cause.
- Using `Must` anywhere the caller could reasonably recover — it belongs in start-up and tests only.
