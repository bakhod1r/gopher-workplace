# Report The Known Cause

## Intuition

A boundary decides how much internal structure leaks out. Collapsing recognised causes gives users a stable vocabulary; passing the rest through unchanged keeps information that would otherwise be lost forever.

## Approach

1. Return nil for nil.
2. Test each public sentinel with `errors.Is`.
3. Return `err` unchanged as the default.

## Solution

```go
switch {
case err == nil:
	return nil
case errors.Is(err, ErrNotFound):
	return ErrNotFound
case errors.Is(err, ErrDenied):
	return ErrDenied
default:
	return err
}
```

## Walkthrough

A wrapped unknown error is returned exactly as received — including its `"layer: "` prefix — because nothing recognised it.

## Pitfalls

- Returning a generic error for unknown causes, destroying the detail.
- Using `errors.Unwrap` instead of `errors.Is`, which only reaches one level.
- Returning nil for unrecognised errors, silently swallowing failures.
