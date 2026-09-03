# Join Failures

**Level:** middle
**Topic:** 04-error-handling

## Context

A form validator runs independent rules. The caller wants one error that reports every rule that failed.

## Task

Implement `JoinAll` in [joinerrs.go](joinerrs.go):

1. Return a single error combining all non-nil entries.
2. Return nil when every entry is nil or the slice is empty.
3. Keep each original error matchable with `errors.Is`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  JoinAll([]error{ErrA, ErrB})
Output: an error matching both
```

**Example 2:**

```
Input:  JoinAll([]error{nil, ErrA})
Output: an error matching ErrA
```

**Example 3:**

```
Input:  JoinAll(nil)
Output: nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **errors.Join** | Combines several errors into one, skipping nils. |
| 2 | **Multi-error matching** | `errors.Is` searches every joined branch. |
| 3 | **All-nil input** | `errors.Join` of nothing is nil. |

## Hint

`errors.Join` already drops nil entries and returns nil when nothing is left — you do not need to filter first.

## Validate

```bash
make verify
```
