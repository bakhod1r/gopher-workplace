# Custom Type That Wraps

## Intuition

`Unwrap() error` is the whole contract for participating in error chains. A custom type that implements it becomes as transparent to `errors.Is` as a `fmt.Errorf` wrapper.

## Approach

1. Format the code and the cause's message in `Error`.
2. Return `e.Cause` from `Unwrap`.
3. The standard library does the rest.

## Solution

```go
// Error:
return fmt.Sprintf("[%d] %s", e.Code, e.Cause.Error())

// Unwrap:
return e.Cause
```

## Walkthrough

`errors.Is` on the `fmt.Errorf` wrapper unwraps to the `*CodeError`, calls its `Unwrap`, reaches `ErrDB`, and matches.

## Pitfalls

- Omitting `Unwrap`, so the cause is visible in the text but unmatchable.
- Returning `nil` from `Unwrap` when `Cause` is set.
- Defining the methods on the value receiver while constructing pointers.
