# Commit Or Roll Back

## Intuition

Compensating logic is where failures get lost. The rule is simple: the failure that triggered the rollback outranks the rollback's own failure, and neither may be discarded.

## Approach

1. Return the apply error immediately when it fails.
2. Return nil when confirm succeeds.
3. Otherwise run rollback and join both errors.

## Solution

```go
if err := apply(); err != nil {
	return err
}
if err := confirm(); err != nil {
	return errors.Join(err, rollback())
}
return nil
```

## Walkthrough

`errors.Join` drops a nil rollback result, so a successful rollback leaves only the confirm failure in the returned error.

## Pitfalls

- Rolling back after a failed apply, undoing something that never happened.
- Returning only the rollback error, hiding why the rollback was needed.
- Ignoring the rollback result entirely.
