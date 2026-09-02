# Required Field

**Level:** junior
**Topic:** 04-error-handling

## Context

A support form rejects blank subject lines. Whitespace-only input is blank too — users paste spaces all the time.

## Task

Implement `Require` in [requiretext.go](requiretext.go):

1. Return `ErrRequired` when `s` is empty or contains only whitespace.
2. Return nil for any string with at least one non-space character.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Require("hello")
Output: nil
```

**Example 2:**

```
Input:  Require("")
Output: ErrRequired
```

**Example 3:**

```
Input:  Require("   ")
Output: ErrRequired
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **strings.TrimSpace** | Trimming reduces the blank cases to one check. |
| 2 | **Error-only return** | Validators report nil for success. |
| 3 | **Whitespace input** | Tabs and newlines count as blank. |

## Hint

Trim first, then compare the result against the empty string.

## Validate

```bash
make verify
```
