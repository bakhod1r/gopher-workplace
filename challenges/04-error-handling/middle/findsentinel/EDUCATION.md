# First Matching Failure

## Intuition

Matching tells you *that* something timed out; returning the matched element tells you *which* replica did. Keeping the wrapper preserves the annotation that identifies it.

## Approach

1. Range over the slice.
2. Return the element when `errors.Is(err, target)`.
3. Return nil after the loop.

## Solution

```go
for _, err := range errs {
	if errors.Is(err, target) {
		return err
	}
}
return nil
```

## Walkthrough

The match is `"replica 2: timeout"`, not the bare `ErrTimeout` sentinel — the caller can now name the replica.

## Pitfalls

- Returning `target` on a match, discarding the context.
- Filtering nils by hand; `errors.Is(nil, target)` is already false.
- Returning the last match instead of the first.
