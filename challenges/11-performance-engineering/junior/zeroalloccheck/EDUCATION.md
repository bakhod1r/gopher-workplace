# A Regression Guard You Can Commit

## Intuition

The measurement already exists in the standard library. The helper's job is to normalise the inputs and turn a float into a pass/fail with enough context to debug.

## Approach

1. Clamp `runs` and `limit`.
2. Round the measured average.
3. Fill in the result struct.

## Solution

```go
func Check(runs, limit int, f func()) Result {
	if runs < 1 {
		runs = 1
	}
	if limit < 0 {
		limit = 0
	}
	allocs := int(math.Round(testing.AllocsPerRun(runs, f)))
	return Result{Allocs: allocs, Limit: limit, OK: allocs <= limit}
}
```

## Walkthrough

Returning the measured count alongside the verdict is what makes the failure message useful: "3 allocations, limit 1" points straight at the regression.

## Pitfalls

- A strict `<`, which fails a function that sits exactly on its documented budget.
- Truncating the average instead of rounding, hiding an allocation that happens on most runs.
- Importing `testing` into production code — this helper belongs in a test file or a test-only package.
