# Runtime Panic Boundary

**Level:** senior
**Topic:** 04-error-handling

## Context

A worker processes records supplied by users. An out-of-range index in one record should fail that record and continue with the next.

## Task

Implement `Run` in [runsafe.go](runsafe.go):

1. Return `f`'s error when it returns normally.
2. Convert a recovered runtime panic into an error wrapping `ErrRuntime`.
3. Preserve the panic's own message in the result.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Run(okFn)
Output: nil
```

**Example 2:**

```
Input:  Run(indexPanic)
Output: an error matching ErrRuntime
```

**Example 3:**

```
Input:  Run(errFn)
Output: the returned error
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **runtime.Error** | Runtime panics carry a typed error value. |
| 2 | **Wrapping recovered values** | Context plus a matchable sentinel. |
| 3 | **Distinguishing failure kinds** | A returned error is not a panic. |

## Hint

A runtime panic's payload already implements `error` — its message is worth keeping in the wrapper.

## Validate

```bash
make verify
```
