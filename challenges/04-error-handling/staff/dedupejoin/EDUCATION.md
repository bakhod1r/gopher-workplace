# Join Without Repeats

## Intuition

Deduplicating before the join is what keeps the resulting tree small. Joining first and filtering later means the duplicates are already part of the value everything else has to traverse.

## Approach

1. Track seen messages in a map.
2. Skip nil and already-seen entries.
3. Join the kept errors.

## Solution

```go
seen := make(map[string]bool)
var kept []error
for _, err := range errs {
	if err == nil {
		continue
	}
	msg := err.Error()
	if seen[msg] {
		continue
	}
	seen[msg] = true
	kept = append(kept, err)
}
return errors.Join(kept...)
```

## Walkthrough

`sameAsA` is a distinct value with the same message, so it is dropped and `errors.Is` no longer matches it.

## Pitfalls

- Deduplicating by value, which keeps distinct values with identical text.
- Joining first and trying to filter the result.
- Returning a non-nil join when everything was nil.
