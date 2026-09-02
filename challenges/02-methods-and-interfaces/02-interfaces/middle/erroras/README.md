# errors.As

**Level:** middle
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

An HTTP client wraps transport failures, and the retry logic needs the status code buried inside.

## Task

Implement the stub(s) in [erroras.go](erroras.go):

1. Implement `Error` on `*HTTPError`.
2. Implement `Call`, which returns a wrapped `*HTTPError` for any non-200 status.
3. Implement `StatusOf`, which digs the status code out of an error chain (0 when there is no `*HTTPError`).
4. Implement `Retryable`, which reports whether the status is 500 or above.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Call(500)
Output: error "call: http 500"
```

**Example 2:**

```
Input:  StatusOf(that error)
Output: 500
```

**Example 3:**

```
Input:  Retryable(Call(404))
Output: false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **errors.As** | Extracts the first error in the chain matching a target type. |
| 2 | **Custom error types** | Reused: structured fields instead of parsed strings. |
| 3 | **Wrapping with %w** | Reused: context without losing the concrete type. |

## Hint

`var he *HTTPError; if errors.As(err, &he) { ... }` — pass a pointer to the target variable.

## Validate

```bash
make verify
```
