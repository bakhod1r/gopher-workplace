# Error With A Trace

## Intuition

Users and log pipelines want different amounts of an error. Keeping the interface method minimal and exposing detail through extra methods lets one value serve both.

## Approach

1. Return the cause's message from `Error`.
2. Return the cause from `Unwrap`.
3. Format the operation and cause in `Trace`.

## Solution

```go
// Error:
return e.Cause.Error()

// Unwrap:
return e.Cause

// Trace:
return fmt.Sprintf("%s -> %s", e.Op, e.Cause.Error())
```

## Walkthrough

`errors.Is` reaches `ErrDisk` through `Unwrap`, which is independent of what `Error` chooses to print.

## Pitfalls

- Putting the trace into `Error()`, leaking internals to users.
- Omitting `Unwrap`, so the cause is unmatchable.
- Dereferencing a nil `Cause` in any of the three methods.
