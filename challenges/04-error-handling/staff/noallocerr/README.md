# Allocation-Free Failure

**Level:** staff
**Topic:** 04-error-handling

## Context

A hot path returns a failure millions of times a second. Formatting an error there dominates the profile, so the common case must not allocate.

## Task

Implement `Check` in [noallocerr.go](noallocerr.go):

1. Return `ErrEmpty` unchanged when `s` is empty.
2. Return nil when `s` is non-empty.
3. Perform no allocation on either path.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Check("")
Output: ErrEmpty
```

**Example 2:**

```
Input:  Check("a")
Output: nil
```

**Example 3:**

```
Input:  testing.AllocsPerRun(…)
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Sentinel reuse** | A package-level error costs nothing to return. |
| 2 | **fmt.Errorf allocates** | Formatting builds a new value every call. |
| 3 | **Measuring allocations** | `testing.AllocsPerRun` proves the claim. |

## Hint

The test measures allocations — anything built per call will fail it.

## Validate

```bash
make verify
```
