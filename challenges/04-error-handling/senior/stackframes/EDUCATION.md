# Capture Frames

## Intuition

A stack is captured as program counters first and resolved to names second. Resolving is the expensive half, which is why the capture step takes a bounded slice.

## Approach

1. Return an empty slice for a non-positive `max`.
2. Fill a `[]uintptr` with `runtime.Callers(2, pc)`.
3. Iterate `runtime.CallersFrames` collecting `frame.Function`.

## Solution

```go
if max <= 0 {
	return []string{}
}
pc := make([]uintptr, max)
n := runtime.Callers(2, pc)
frames := runtime.CallersFrames(pc[:n])
out := make([]string, 0, n)
for {
	frame, more := frames.Next()
	out = append(out, frame.Function)
	if !more || len(out) == max {
		break
	}
}
return out
```

## Walkthrough

Skip 2 discards `runtime.Callers` itself and `Frames`, so the first name is the test function that called it.

## Pitfalls

- Using skip 0, which reports runtime internals.
- Sizing the PC slice to a fixed constant instead of `max`.
- Ignoring `More()` and reading past the end of the frames.
