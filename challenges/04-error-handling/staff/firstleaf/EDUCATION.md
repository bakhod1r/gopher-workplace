# The Original Failure

## Intuition

In a chain the root cause is unambiguous. In a tree it is not — a join has several roots, so the contract must pick one, and "leftmost" matches the order the failures were collected in.

## Approach

1. Return nil for nil.
2. Descend into the first joined branch when present.
3. Descend into a wrapped child when present.
4. Return the error when neither applies.

## Solution

```go
if err == nil {
	return nil
}
if joined, ok := err.(interface{ Unwrap() []error }); ok {
	members := joined.Unwrap()
	if len(members) == 0 {
		return err
	}
	return Origin(members[0])
}
if wrapped, ok := err.(interface{ Unwrap() error }); ok {
	return Origin(wrapped.Unwrap())
}
return err
```

## Walkthrough

`errors.Join(fmt.Errorf("x: %w", ErrB), ErrA)` descends the first branch and returns `ErrB`, never reaching `ErrA`.

## Pitfalls

- Returning the last leaf, which is a different contract.
- Assuming a joined error always has members.
- Using `errors.Unwrap`, which returns nil for joins and stops early.
