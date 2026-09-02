# Result Or Zero

**Level:** junior
**Topic:** 04-error-handling

## Context

A metrics collector sums per-host readings. A host that failed contributes nothing instead of a garbage number.

## Task

Implement `OrZero` in [orzero.go](orzero.go):

1. Return `v` when `err` is nil.
2. Return `0` when `err` is non-nil, whatever `v` holds.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  OrZero(42, nil)
Output: 42
```

**Example 2:**

```
Input:  OrZero(42, ErrHost)
Output: 0
```

**Example 3:**

```
Input:  OrZero(0, nil)
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Never trust a value past an error** | A failed call's result is undefined. |
| 2 | **Fallback values** | Zero is a safe neutral for a sum. |
| 3 | **Error-first checking** | The error decides which branch runs. |

## Hint

The value argument is deliberately non-zero in the failing test case — do not pass it through.

## Validate

```bash
make verify
```
