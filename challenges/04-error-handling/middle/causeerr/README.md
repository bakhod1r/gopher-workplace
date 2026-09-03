# Custom Type That Wraps

**Level:** middle
**Topic:** 04-error-handling

## Context

A service error carries a numeric code and the failure that triggered it. Callers must still match the underlying cause.

## Task

Implement `CodeError` in [causeerr.go](causeerr.go):

1. Give `*CodeError` an `Error() string` of the form `"[<Code>] <Cause>"`.
2. Give it an `Unwrap() error` returning `Cause`.
3. Make `errors.Is(&CodeError{Cause: ErrDB}, ErrDB)` true.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  (&CodeError{Code: 7, Cause: ErrDB}).Error()
Output: "[7] db unavailable"
```

**Example 2:**

```
Input:  errors.Unwrap(&CodeError{Cause: ErrDB})
Output: ErrDB
```

**Example 3:**

```
Input:  errors.Is(wrapped, ErrDB)
Output: true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Implementing Unwrap** | One method makes a custom type transparent to `errors.Is`. |
| 2 | **Chain participation** | Custom errors join the standard machinery. |
| 3 | **Message plus cause** | Both are rendered, only one is matched. |

## Hint

`errors.Is` never inspects your message — it calls `Unwrap`, which is the second missing method.

## Validate

```bash
make verify
```
