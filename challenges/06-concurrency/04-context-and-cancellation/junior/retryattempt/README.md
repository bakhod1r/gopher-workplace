# Overriding a Value per Attempt

**Level:** junior
**Topic:** 06-concurrency → 04-context-and-cancellation

## Context

The outbound HTTP client retries idempotent calls to the shipping provider. Each retry derives a fresh context tagged with the attempt number, so the request carries an `X-Attempt` header for the provider's idempotency logic and every log line says which try it belongs to. The original request context is left untouched.

## Task

Implement the exported function(s) in [retryattempt.go](retryattempt.go) so that:

1. `WithAttempt` stores `n` under the unexported key and returns the new context.
2. `Attempt` returns the stored `int`, or `0` when the key is absent or holds another type.
3. Wrapping again overrides the value for the new context only.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  Attempt(WithAttempt(bg, 2))
Output: 2
```

**Example 2:**

```
Input:  Attempt(context.Background())
Output: 0
```

**Example 3:**

```
Input:  Attempt(WithAttempt(WithAttempt(bg, 1), 2))
Output: 2
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Shadowing** | A newer `WithValue` for the same key hides the older one for that branch. |
| 2 | **Immutable contexts** | Wrapping never mutates the parent — older contexts still see the old value. |
| 3 | **Typed lookups** | The value is stored and asserted as `int`; a string under the key is a miss. |

## Hint

`Attempt` uses `n, _ := ctx.Value(attemptKey{}).(int)` — the zero value of `int` is the wanted default.

## Validate

```bash
make verify
```
