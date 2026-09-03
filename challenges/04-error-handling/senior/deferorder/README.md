# Deferred Cleanup Order

**Level:** senior
**Topic:** 04-error-handling

## Context

A resource stack must be released in reverse order of acquisition, and every release must run even if an earlier one fails.

## Task

Implement `CloseAll` in [deferorder.go](deferorder.go):

1. Call every closer, last registered first.
2. Run all of them even when some return errors.
3. Return the failures combined, or nil when all succeeded.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  CloseAll(a, b, c)
Output: closes c, b, a
```

**Example 2:**

```
Input:  CloseAll()
Output: nil
```

**Example 3:**

```
Input:  CloseAll(failing, ok)
Output: the failure
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **LIFO release** | Reverse order mirrors acquisition. |
| 2 | **Continuing past failure** | One bad closer must not strand the rest. |
| 3 | **errors.Join** | Independent cleanup failures combine. |

## Hint

Iterate backwards explicitly — you are reproducing what `defer` would do, not using it.

## Validate

```bash
make verify
```
