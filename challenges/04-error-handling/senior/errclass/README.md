# Count By Class

**Level:** senior
**Topic:** 04-error-handling

## Context

A dashboard breaks a batch's failures down by class so operators see whether a spike is one bad dependency or many.

## Task

Implement `Classify` in [errclass.go](errclass.go):

1. Return a map from class name to count.
2. Classify by `ErrTimeout` → `"timeout"`, `ErrDenied` → `"denied"`, anything else → `"other"`.
3. Skip nil entries, and return an empty map when nothing failed.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Classify([]error{ErrTimeout, ErrTimeout})
Output: {"timeout": 2}
```

**Example 2:**

```
Input:  Classify([]error{nil})
Output: {}
```

**Example 3:**

```
Input:  Classify([]error{errors.New("x")})
Output: {"other": 1}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Classification with errors.Is** | Wrapped failures still classify. |
| 2 | **Map accumulation** | Counting into a map by key. |
| 3 | **Empty result** | An initialised empty map, not nil. |

## Hint

The tests compare against an empty non-nil map — build it up front.

## Validate

```bash
make verify
```
