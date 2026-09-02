# Error Interface

## Intuition

`error` is not special machinery — it is an ordinary one-method interface. Defining your own error type means callers can inspect fields instead of parsing strings.

## Approach

1. Format the message with `fmt.Sprintf("%s: %s", e.Field, e.Message)`.
2. In `Validate`, return `&ValidationError{...}` for an empty name.
3. Return the untyped `nil` otherwise.

## Solution

```go
func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

func Validate(name string) error {
	if name == "" {
		return &ValidationError{Field: "name", Message: "required"}
	}
	return nil
}
```

## Walkthrough

`Validate("")` builds a `*ValidationError`. `errors.As` unwraps the interface back to that concrete pointer, so the test can read `Field` and `Message`.

## Pitfalls

- Declaring `var e *ValidationError` and returning it on the success path — a non-nil `error` holding a nil pointer.
- Value receiver on `Error` while returning `&ValidationError{}` — still works, but mixing receivers confuses the method set.
- Treating `" "` as empty when the spec says only `""` is invalid.
