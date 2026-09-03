# Custom Error Type

## Intuition

A sentinel says *which* failure happened. A custom type says which failure happened *and* carries the details — the caller reads fields instead of parsing a message.

## Approach

1. Format the message from the two fields in `Error`.
2. Return `&ValidationError{...}` from the constructor.
3. The pointer satisfies `error` because the method has a pointer receiver.

## Solution

```go
// Error:
return fmt.Sprintf("%s: %s", e.Field, e.Reason)

// NewValidation:
return &ValidationError{Field: field, Reason: reason}
```

## Walkthrough

`errors.As` matches `*ValidationError` even under a `fmt.Errorf("save: %w", …)` wrapper, so `Field` is still readable at the top layer.

## Pitfalls

- Defining `Error()` on the value receiver while returning a pointer, or vice versa.
- Returning `ValidationError` (a value) where the method set requires a pointer.
- Building the message in the constructor and dropping the fields.
