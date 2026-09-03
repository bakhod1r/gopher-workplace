# Strip Internal Detail

## Intuition

Error messages are an output channel like any other. Mapping internal causes onto a fixed public vocabulary means new internal detail can never accidentally become part of the API.

## Approach

1. Return nil for nil.
2. Match each recognised internal cause with `errors.Is`.
3. Collapse everything else to the generic error.

## Solution

```go
switch {
case err == nil:
	return nil
case errors.Is(err, errInternalMissing):
	return ErrPublicNotFound
case errors.Is(err, errInternalParse):
	return ErrPublicInvalid
default:
	return ErrPublicInternal
}
```

## Walkthrough

The unknown error's message names a path, but the default branch discards it entirely and returns the fixed generic error.

## Pitfalls

- Wrapping the internal error into the public one, which leaks the message through `Error()`.
- Passing unrecognised errors through unchanged.
- Sanitising in several places instead of once at the boundary.
