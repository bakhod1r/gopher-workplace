# Combine Work And Cleanup

**Level:** middle
**Topic:** 04-error-handling

## Context

A file-processing step must always run its cleanup, and a cleanup failure must not disappear just because the work succeeded — or vice versa.

## Task

Implement `Do` in [safeclose.go](safeclose.go):

1. Call `work`, then always call `cleanup`.
2. Return both failures combined when both fail.
3. Return nil only when both succeeded.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Do(okFn, okFn)
Output: nil
```

**Example 2:**

```
Input:  Do(failWork, okFn)
Output: the work error
```

**Example 3:**

```
Input:  Do(failWork, failCleanup)
Output: an error matching both
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Cleanup always runs** | Failure of the work must not skip the cleanup. |
| 2 | **errors.Join** | Two independent failures combine without a hierarchy. |
| 3 | **Losing errors** | Overwriting one error with the other hides a failure. |

## Hint

Both functions run in every case; the only question is what you return when each of them fails.

## Validate

```bash
make verify
```
