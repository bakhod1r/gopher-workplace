# Join Failures

## Intuition

Some failures are not a chain but a set: three rules failed independently and none caused the others. `errors.Join` builds a tree that `errors.Is` searches across every branch.

## Approach

1. Pass the slice to `errors.Join` with `...`.
2. Return the result — it is already nil when nothing failed.

## Solution

```go
return errors.Join(errs...)
```

## Walkthrough

For `[]error{nil, ErrA, nil}` the nils are dropped and a single-branch join remains, so `errors.Is` matches `ErrA` but not `ErrB`.

## Pitfalls

- Filtering nils first and returning an empty join — the extra work changes nothing.
- Building the message by hand, losing `errors.Is` support.
- Returning a non-nil wrapper when every entry was nil.
