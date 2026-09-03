# Wrap Only Real Failures

## Intuition

Formatting functions never return nil. Any wrapper that forgets the nil check converts every success into a failure — and the caller's `if err != nil` then fires on healthy paths.

## Approach

1. Return nil when `err == nil`.
2. Otherwise format with `%w`.

## Solution

```go
if err == nil {
	return nil
}
return fmt.Errorf("%s: %w", msg, err)
```

## Walkthrough

`WrapNonNil("", ErrX)` still wraps — an empty message is not a reason to skip annotation, only a nil error is.

## Pitfalls

- Wrapping unconditionally, producing `"step: %!w(<nil>)"`.
- Guarding on `msg == ""` instead of on the error.
- Returning `err` unwrapped when the message is empty.
