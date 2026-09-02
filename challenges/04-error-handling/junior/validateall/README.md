# Collect All Failures

**Level:** junior
**Topic:** 04-error-handling

## Context

A bulk import validates every row before writing anything. The operator wants the full list of problems, not just the first one.

## Task

Implement `Validate` in [validateall.go](validateall.go):

1. Return one `ErrNegative` entry for each negative value, in input order.
2. Return nil for an empty or nil slice.
3. Return nil when every value is valid.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Validate([]int{1, -2, -3})
Output: [ErrNegative ErrNegative]
```

**Example 2:**

```
Input:  Validate([]int{1, 2})
Output: nil
```

**Example 3:**

```
Input:  Validate(nil)
Output: nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Collecting vs aborting** | Reporting everything needs a full pass. |
| 2 | **append to a nil slice** | `append` on a nil slice allocates on demand. |
| 3 | **Nil result** | No failures means a nil slice, not an empty one. |

## Hint

Declare the result with `var out []error` so the no-failure case is nil without extra work.

## Validate

```bash
make verify
```
