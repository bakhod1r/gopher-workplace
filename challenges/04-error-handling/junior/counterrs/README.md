# Count Failures

**Level:** junior
**Topic:** 04-error-handling

## Context

A nightly batch job records one result per record. The summary line reports how many of them failed.

## Task

Implement `CountErrors` in [counterrs.go](counterrs.go):

1. Return the number of non-nil entries in the slice.
2. Return `0` for an empty or nil slice.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  CountErrors([]error{nil, ErrX, ErrX})
Output: 2
```

**Example 2:**

```
Input:  CountErrors([]error{nil})
Output: 0
```

**Example 3:**

```
Input:  CountErrors(nil)
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Accumulator loop** | A counter declared before the loop, returned after. |
| 2 | **Nil check** | Only non-nil entries count. |
| 3 | **Zero value** | `var n int` starts at 0 — the right answer for an empty slice. |

## Hint

Unlike a search, this loop must visit every element before returning.

## Validate

```bash
make verify
```
