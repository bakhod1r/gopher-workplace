# Failure To Status

## Intuition

A boundary translates domain failures into the vocabulary of its transport. The default branch is the important one: an unclassified failure is a server error, never an accidental success.

## Approach

1. Return 200 for nil.
2. Test each sentinel with `errors.Is`.
3. Return 500 as the default.

## Solution

```go
switch {
case err == nil:
	return 200
case errors.Is(err, ErrNotFound):
	return 404
case errors.Is(err, ErrDenied):
	return 403
case errors.Is(err, ErrConflict):
	return 409
default:
	return 500
}
```

## Walkthrough

A wrapped `ErrNotFound` reaches the second case because `errors.Is` searches the chain, so the annotation does not change the status.

## Pitfalls

- Omitting the nil case, so success maps to 500.
- Using `==` instead of `errors.Is`, which misses wrapped sentinels.
- Defaulting to 200 or 400 for unknown failures.
