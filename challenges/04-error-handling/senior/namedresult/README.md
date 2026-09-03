# Annotate On The Way Out

**Level:** senior
**Topic:** 04-error-handling

## Context

Every method in a repository layer must prefix its failures with the operation name. Doing it at each return statement is where the mistakes creep in.

## Task

Implement `Do` in [namedresult.go](namedresult.go):

1. Return the value from `f` unchanged on success.
2. Annotate any non-nil error as `"<op>: <err>"` in one place.
3. Preserve the wrapped error for `errors.Is`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Do("load", okFn)
Output: 7, nil
```

**Example 2:**

```
Input:  Do("load", failFn)
Output: 0, "load: boom"
```

**Example 3:**

```
Input:  errors.Is(err, ErrBoom)
Output: true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Deferred annotation** | One defer covers every return path. |
| 2 | **Named results** | The defer reads and rewrites `err`. |
| 3 | **Single point of change** | Adding a return cannot forget the wrapper. |

## Hint

A single deferred closure that reads the named `err` and rewrites it covers every exit, including future ones.

## Validate

```bash
make verify
```
