# Error With A Trace

**Level:** senior
**Topic:** 04-error-handling

## Context

A service keeps error messages short for users while retaining a longer trace for the log pipeline.

## Task

Implement `TracedError` in [tracederr.go](tracederr.go):

1. Give `*TracedError` an `Error() string` returning only the wrapped message.
2. Give it an `Unwrap() error` returning the cause.
3. Give it a `Trace() string` returning `"<Op> -> <cause message>"`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  (&TracedError{Op: "load", Cause: ErrDisk}).Error()
Output: "disk offline"
```

**Example 2:**

```
Input:  …Trace()
Output: "load -> disk offline"
```

**Example 3:**

```
Input:  errors.Is(traced, ErrDisk)
Output: true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Separating audiences** | Short message, detailed trace. |
| 2 | **Unwrap participation** | Custom types join the chain. |
| 3 | **Extra methods** | Callers reach detail through type assertion. |

## Hint

Three methods, one struct. Only one of them is what `errors.Is` uses.

## Validate

```bash
make verify
```
