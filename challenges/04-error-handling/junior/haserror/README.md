# Error Present

**Level:** junior
**Topic:** 04-error-handling

## Context

A dashboard shows a red badge whenever a step reported a problem. It needs one small helper that answers a single question.

## Task

Implement `HasError` in [haserror.go](haserror.go):

1. Return `true` when `err` is non-nil.
2. Return `false` when `err` is nil.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  HasError(nil)
Output: false
```

**Example 2:**

```
Input:  HasError(errors.New("boom"))
Output: true
```

**Example 3:**

```
Input:  HasError(ErrSample)
Output: true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Nil error** | A nil `error` means success — nothing else does. |
| 2 | **Interface values** | `error` is an interface; its zero value is nil. |
| 3 | **Boolean expressions** | `err != nil` is already the answer. |

## Hint

The comparison you write in every Go function is the whole body here.

## Validate

```bash
make verify
```
