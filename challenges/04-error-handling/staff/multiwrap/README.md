# Two Causes In One Error

**Level:** staff
**Topic:** 04-error-handling

## Context

A request failed both its primary and its fallback path. The single error returned must keep both causes matchable.

## Task

Implement `Both` in [multiwrap.go](multiwrap.go):

1. Return nil when both errors are nil.
2. Return an error whose message is `"<primary>; <fallback>"` when both fail.
3. Keep both matchable by `errors.Is`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Both(ErrA, ErrB).Error()
Output: "a; b"
```

**Example 2:**

```
Input:  errors.Is(Both(ErrA, ErrB), ErrB)
Output: true
```

**Example 3:**

```
Input:  Both(nil, nil)
Output: nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Multiple %w verbs** | `fmt.Errorf` accepts more than one since Go 1.20. |
| 2 | **Unwrap() []error** | The result exposes both causes. |
| 3 | **Message control** | Unlike Join, the separator is yours. |

## Hint

One `fmt.Errorf` call can wrap two errors at once — the result implements the multi-error unwrap shape.

## Validate

```bash
make verify
```
