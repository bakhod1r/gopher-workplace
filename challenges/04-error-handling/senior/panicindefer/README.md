# Cleanup That Panics

**Level:** senior
**Topic:** 04-error-handling

## Context

A resource wrapper must survive a cleanup function that itself panics, reporting the failure instead of unwinding the whole request.

## Task

Implement `Guard` in [panicindefer.go](panicindefer.go):

1. Run `work`, then always run `cleanup`.
2. Convert a panic in either function into an error wrapping `ErrPanic`.
3. Return the work error when only the work fails.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Guard(okFn, okFn)
Output: nil
```

**Example 2:**

```
Input:  Guard(okFn, panicFn)
Output: an error matching ErrPanic
```

**Example 3:**

```
Input:  Guard(failFn, okFn)
Output: the work error
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Recovering around cleanup** | Cleanup is code too, and can fail. |
| 2 | **Nested defers** | Each recovery needs its own deferred scope. |
| 3 | **Named results** | Both paths write the same result variable. |

## Hint

Wrapping the cleanup call in its own function with its own defer keeps its panic from escaping.

## Validate

```bash
make verify
```
