# Aggregate Error Type

## Intuition

`errors.Join` returns an opaque type. Defining your own aggregate gives you control of the message while `Unwrap() []error` keeps the standard matching machinery working.

## Approach

1. Collect each member's message into a `[]string`.
2. Join them with `"; "`.
3. Return the slice itself from `Unwrap`.

## Solution

```go
// Error:
parts := make([]string, 0, len(e))
for _, err := range e {
	parts = append(parts, err.Error())
}
return strings.Join(parts, "; ")

// Unwrap:
return e
```

## Walkthrough

`errors.Is(Errors{ErrA}, ErrB)` fans out over one member, fails to match, and correctly returns false.

## Pitfalls

- Implementing only `Error()`, so `errors.Is` never sees the members.
- Writing `Unwrap() error` instead of `Unwrap() []error` — the wrong shape is ignored.
- Building the message with `+` and leaving a trailing separator.
