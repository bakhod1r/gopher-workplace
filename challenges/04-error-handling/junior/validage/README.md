# Age Validation

**Level:** junior
**Topic:** 04-error-handling

## Context

A signup form stores ages. Anything outside a plausible human range is a data-entry mistake and must be rejected at the boundary.

## Task

Implement `ValidAge` in [validage.go](validage.go):

1. Return `ErrTooYoung` when `age` is below 0.
2. Return `ErrTooOld` when `age` is above 130.
3. Return nil for any age in `[0, 130]`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  ValidAge(30)
Output: nil
```

**Example 2:**

```
Input:  ValidAge(-1)
Output: ErrTooYoung
```

**Example 3:**

```
Input:  ValidAge(200)
Output: ErrTooOld
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Error-only return** | A validator returns just `error`; nil means valid. |
| 2 | **Distinct sentinels** | Different failures deserve different error values. |
| 3 | **Inclusive boundaries** | 0 and 130 are both valid. |

## Hint

Two guards, two different sentinels, then `nil`. Watch the boundary values.

## Validate

```bash
make verify
```
