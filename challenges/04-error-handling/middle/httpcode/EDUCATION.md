# Extract A Typed Error

## Intuition

`errors.Is` answers "is this the same error?"; `errors.As` answers "is there an error of this type in the chain, and can I have it?" Data attached to an error is retrieved with `As`.

## Approach

1. Declare `var httpErr *HTTPError`.
2. Call `errors.As(err, &httpErr)`.
3. Return the code and true on success, `0, false` otherwise.

## Solution

```go
var httpErr *HTTPError
if errors.As(err, &httpErr) {
	return httpErr.Code, true
}
return 0, false
```

## Walkthrough

For the double-wrapped case `errors.As` unwraps twice, matches the `*HTTPError` type, assigns it to `httpErr`, and the code `503` is read from the struct.

## Pitfalls

- Using `err.(*HTTPError)`, which fails on any wrapped error.
- Passing `httpErr` instead of `&httpErr` — `errors.As` panics on a non-pointer target.
- Returning the code without checking the boolean result.
