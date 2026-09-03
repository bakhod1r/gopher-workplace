# Combine Work And Cleanup

## Intuition

The classic bug in cleanup code is assignment: `err = cleanup()` overwrites the real failure, or the cleanup error is discarded entirely. Joining keeps both, and neither one hides the other.

## Approach

1. Call `work` and keep its result.
2. Call `cleanup` and keep its result.
3. Return `errors.Join` of the two.

## Solution

```go
workErr := work()
cleanupErr := cleanup()
return errors.Join(workErr, cleanupErr)
```

## Walkthrough

When both fail, `errors.Join` builds a two-branch error and `errors.Is` matches each of them independently.

## Pitfalls

- Returning early on the work failure, skipping the cleanup.
- Assigning `err = cleanup()`, which discards the work failure.
- Ignoring the cleanup result entirely.
