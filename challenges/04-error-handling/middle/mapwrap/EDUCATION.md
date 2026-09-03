# Contextual Lookup Failure

## Intuition

Sentinels are for callers; messages are for humans. Wrapping the sentinel with the key satisfies both without forcing every caller to define its own error type.

## Approach

1. Read with the comma-ok idiom.
2. On a miss, wrap `ErrNotFound` with the key.
3. Otherwise return the value and nil.

## Solution

```go
v, ok := m[key]
if !ok {
	return 0, fmt.Errorf("key %s: %w", key, ErrNotFound)
}
return v, nil
```

## Walkthrough

The key `"zero"` stores 0 and is present, so no error is produced — the comma-ok check, not the value, decides.

## Pitfalls

- Returning the bare sentinel, losing the key.
- Using `%v` for the sentinel and breaking `errors.Is`.
- Reporting a miss for a legitimately stored zero.
