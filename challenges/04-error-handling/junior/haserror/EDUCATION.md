# Error Present

## Intuition

An `error` is an interface value. Success is represented by exactly one value — nil. Any non-nil error, even one with an empty message, is a failure.

## Approach

1. Compare `err` against nil.
2. Return the result of that comparison directly.

## Solution

```go
return err != nil
```

## Walkthrough

For `errors.New("")` the message is empty but the interface value is not nil, so the answer is `true`.

## Pitfalls

- Checking `err.Error() != ""` — an empty message is still an error.
- Calling `err.Error()` on a nil error — that panics.
- Writing an `if/else` that returns `true`/`false` instead of the expression.
