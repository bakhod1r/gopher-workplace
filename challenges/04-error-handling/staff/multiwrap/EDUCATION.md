# Two Causes In One Error

## Intuition

A single `fmt.Errorf` with two `%w` verbs produces an error with two causes and a message you control — the middle ground between one wrap and an opaque `errors.Join`.

## Approach

1. Return nil when both inputs are nil.
2. Format both with `%w` and the chosen separator.

## Solution

```go
if primary == nil && fallback == nil {
	return nil
}
return fmt.Errorf("%w; %w", primary, fallback)
```

## Walkthrough

The returned error implements `Unwrap() []error` with two members, which is why `errors.Is` matches each of them.

## Pitfalls

- Wrapping one and formatting the other with `%v`, losing half the matchability.
- Assuming multiple `%w` verbs need Go 1.19 or earlier syntax — they do not exist before 1.20.
- Returning a non-nil error when both causes are nil.
