# Commit Or Roll Back

**Level:** senior
**Topic:** 04-error-handling

## Context

A two-step write applies a change and then confirms it. A failure to confirm must trigger a rollback, and a failing rollback must not hide the original failure.

## Task

Implement `Do` in [twophase.go](twophase.go):

1. Run `apply`, and return its error without calling anything else when it fails.
2. Run `confirm` after a successful apply, returning nil when it succeeds.
3. Run `rollback` when `confirm` fails, returning both failures combined.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Do(okApply, okConfirm, rb)
Output: nil, rollback not called
```

**Example 2:**

```
Input:  Do(failApply, c, rb)
Output: the apply error only
```

**Example 3:**

```
Input:  Do(okApply, failConfirm, failRb)
Output: both failures
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Compensating actions** | Rollback runs only when there is something to undo. |
| 2 | **Error precedence** | The original failure is never lost. |
| 3 | **errors.Join for two causes** | Neither failure hides the other. |

## Hint

Rollback belongs on exactly one path — after a successful apply whose confirm failed.

## Validate

```bash
make verify
```
