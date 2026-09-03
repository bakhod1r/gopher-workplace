# Cap The Chain

## Intuition

Every layer wrapping unconditionally produces messages where the same failure is restated six times. Capping the depth keeps the most specific context — the innermost layers — and drops the repetition.

## Approach

1. Return nil for nil.
2. Count the chain with repeated `errors.Unwrap`.
3. Return `err` unchanged at or above the limit, otherwise wrap it.

## Solution

```go
if err == nil {
	return nil
}
depth := 0
for e := err; e != nil; e = errors.Unwrap(e) {
	depth++
}
if depth >= max {
	return err
}
return fmt.Errorf("%s: %w", msg, err)
```

## Walkthrough

A two-link chain hits the limit of 2, so the third annotation is refused and the identical error value is returned.

## Pitfalls

- Returning a copy or re-wrap when refusing, breaking identity comparisons.
- Counting after wrapping, which lets the chain grow one link past the limit.
- Treating `max` of 0 as unlimited.
