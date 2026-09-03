# Fallback Chain

## Intuition

A fallback chain is optimistic on the way in and pessimistic on the way out: stop at the first success, but when nothing works the operator needs every reason, not just the last.

## Approach

1. Return `ErrNoSources` for an empty list.
2. Return on the first success, collecting failures otherwise.
3. Join the collected failures.

## Solution

```go
if len(sources) == 0 {
	return 0, ErrNoSources
}
var errs []error
for _, src := range sources {
	v, err := src()
	if err == nil {
		return v, nil
	}
	errs = append(errs, err)
}
return 0, errors.Join(errs...)
```

## Walkthrough

When the first source succeeds the loop returns immediately, so the second is never called and no errors are collected.

## Pitfalls

- Calling every source regardless, wasting work and side effects.
- Returning only the last failure, hiding earlier reasons.
- Treating an empty list as success.
