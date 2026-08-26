# Multi Check

**Level:** middle
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A registration form validates user input. Instead of failing on the first
error, collect all validation failures.

## Task

Implement `Validate` on `User` in [multicheck.go](multicheck.go):

1. Check: Name not empty → `"name is required"`.
2. Check: Email contains "@" → `"invalid email"`.
3. Check: Age >= 0 → `"age must be non-negative"`.
4. Return nil if all pass.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  User{"Go", "go@go.dev", 10}.Validate()
Output: nil
```

**Example 2:**

```
Input:  User{"", "bad", -1}.Validate()
Output: ["name is required", "invalid email", "age must be non-negative"]
```

**Example 3:**

```
Input:  User{"X", "nope", 0}.Validate()
Output: ["invalid email"]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Value receiver** | Read-only validation. |
| 2 | **Collecting errors** | Append to a slice; return nil if empty. |
| 3 | **strings.Contains** | Simple "@" check. |

## Hint

Build a `var errs []string`, append for each failing check, return `errs` (nil
if nothing was appended).

## Validate

```bash
make verify
```
