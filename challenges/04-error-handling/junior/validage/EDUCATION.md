# Age Validation

## Intuition

A validator has nothing to return but a verdict. `error` alone carries it: nil for valid, a specific sentinel for each way the input can be wrong.

## Approach

1. Guard the low bound first.
2. Guard the high bound second.
3. Return nil when neither guard fires.

## Solution

```go
if age < 0 {
	return ErrTooYoung
}
if age > 130 {
	return ErrTooOld
}
return nil
```

## Walkthrough

`ValidAge(130)` passes both guards — `130 > 130` is false — and returns nil.

## Pitfalls

- Using `>=` on the upper bound, rejecting a valid 130.
- Returning one shared error for both failures, so callers cannot tell them apart.
- Returning a bool instead of an error.
