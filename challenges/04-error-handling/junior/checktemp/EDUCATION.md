# Temperature Range

## Intuition

A range check is two guards. Naming each side separately lets the operator tell a frozen probe from a burnt one without reading the value.

## Approach

1. Guard the lower bound.
2. Guard the upper bound.
3. Return nil in between.

## Solution

```go
if c < -40 {
	return ErrBelowRange
}
if c > 85 {
	return ErrAboveRange
}
return nil
```

## Walkthrough

`-40.1 < -40` is true, so the lower guard fires; `-40` itself passes both guards.

## Pitfalls

- Using `<=` or `>=`, rejecting the valid bounds.
- Returning one error for both directions.
- Comparing floats with a tolerance where exact bounds are specified.
