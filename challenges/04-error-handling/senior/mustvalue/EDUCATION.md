# Must Helper

## Intuition

`Must` marks a failure the program cannot meaningfully continue past. Panicking with the error value keeps the failure inspectable by whatever recovers it.

## Approach

1. Panic when `err != nil`.
2. Return `v` otherwise.

## Solution

```go
if err != nil {
	panic(err)
}
return v
```

## Walkthrough

The test's deferred `recover` receives the `ErrLoad` value itself, so it compares equal to the sentinel.

## Pitfalls

- Panicking with `err.Error()`, turning a rich error into a bare string.
- Returning an error instead of panicking, defeating the point of `Must`.
- Panicking when `v` is zero rather than when `err` is non-nil.
