# Fallback Chain

**Level:** senior
**Topic:** 04-error-handling

## Context

A configuration loader tries several sources in order. It uses the first one that works and reports every failure when none do.

## Task

Implement `First` in [firstsuccess.go](firstsuccess.go):

1. Return the first successful source's value and nil.
2. Return `0` and all failures combined when every source fails.
3. Return `ErrNoSources` when the list is empty.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  First(okFn)
Output: 7, nil
```

**Example 2:**

```
Input:  First(failFn, okFn)
Output: 7, nil
```

**Example 3:**

```
Input:  First()
Output: 0, ErrNoSources
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Fallback ordering** | Later sources only run when earlier ones fail. |
| 2 | **Reporting all attempts** | A total failure needs every reason. |
| 3 | **Empty input** | No sources is a distinct failure. |

## Hint

Stop calling sources as soon as one succeeds — the tests count invocations.

## Validate

```bash
make verify
```
