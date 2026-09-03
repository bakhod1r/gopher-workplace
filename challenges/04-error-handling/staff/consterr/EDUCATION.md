# Constant Errors

## Intuition

`errors.New` returns a pointer, so its sentinels are unique but mutable package variables. A defined string type gives constants that satisfy `error`, cannot be reassigned, and compare by value.

## Approach

1. Convert the receiver back to `string`.
2. Return it.

## Solution

```go
return string(e)
```

## Walkthrough

`errors.Is` compares the unwrapped value against `ErrClosed`; both are `Error("closed")`, which compares equal by value.

## Pitfalls

- Using a pointer receiver, which a constant cannot satisfy.
- Returning a formatted string, so equal-valued constants stop comparing equal in messages.
- Assuming value-typed sentinels with the same text stay distinct — they do not.
