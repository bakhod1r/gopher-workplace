# Run Every Step

**Level:** middle
**Topic:** 04-error-handling

## Context

A startup sequence runs independent initialisers. All of them must run, and every failure must be reported together.

## Task

Implement `RunAll` in [runall.go](runall.go):

1. Call every function in order, even after one fails.
2. Return a single error combining all failures.
3. Return nil when every step succeeded or the slice is empty.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  RunAll(okFn, okFn)
Output: nil
```

**Example 2:**

```
Input:  RunAll(failA, okFn, failB)
Output: an error matching both failures
```

**Example 3:**

```
Input:  RunAll()
Output: nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Collect, do not abort** | Independent steps all deserve a chance to run. |
| 2 | **errors.Join** | Combines the collected failures. |
| 3 | **Variadic functions** | `fs ...func() error` takes any number of steps. |

## Hint

Collect into a slice as you go, then join once at the end — the tests check that every function ran.

## Validate

```bash
make verify
```
