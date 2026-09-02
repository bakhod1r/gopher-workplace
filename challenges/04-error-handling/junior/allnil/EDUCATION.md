# All Checks Passed

## Intuition

"All succeeded" is disproved by a single counterexample, so the loop looks for a failure and the success answer waits until the loop is exhausted.

## Approach

1. Range over the slice.
2. Return `false` on the first non-nil entry.
3. Return `true` after the loop.

## Solution

```go
for _, err := range errs {
	if err != nil {
		return false
	}
}
return true
```

## Walkthrough

For an empty slice the loop body never runs and `true` is returned — nothing failed.

## Pitfalls

- Returning `true` inside the loop, so only the first entry is judged.
- Special-casing empty to `false`; a gate with no checks is open.
- Building a counter when one early return is enough.
