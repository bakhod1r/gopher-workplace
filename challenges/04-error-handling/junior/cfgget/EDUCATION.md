# Config Value

## Intuition

Configuration has three states, not two: set to a value, set to empty, and absent. Only the comma-ok form separates the last two.

## Approach

1. Read with `v, ok := cfg[key]`.
2. Return the error when `ok` is false.
3. Return `v, nil` otherwise.

## Solution

```go
v, ok := cfg[key]
if !ok {
	return "", ErrMissingKey
}
return v, nil
```

## Walkthrough

Key `"blank"` holds `""`. `ok` is true, so the result is `"", nil` — configured, deliberately empty.

## Pitfalls

- Checking `if cfg[key] == ""`, which rejects a deliberately empty setting.
- Returning a default value instead of an error.
- Guarding a nil map by hand — reads from nil maps are safe.
