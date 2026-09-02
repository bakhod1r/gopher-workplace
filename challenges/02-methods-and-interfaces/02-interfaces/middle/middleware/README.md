# Middleware

**Level:** middle
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A service wraps its core handler in layers that log, count, and short-circuit requests.

## Task

Implement the stub(s) in [middleware.go](middleware.go):

1. Implement `Handle` on `HandlerFunc`.
2. Implement `WithCount`, which wraps a handler and increments a counter on every call.
3. Implement `WithPrefix`, which prefixes the result.
4. Implement `Apply`, which wraps a handler in the given middlewares so the first listed runs outermost.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  WithPrefix("p:")(base).Handle("x") where base echoes
Output: "p:x"
```

**Example 2:**

```
Input:  counted := WithCount(&n)(base); counted.Handle("x"); n
Output: 1
```

**Example 3:**

```
Input:  Apply(base, WithPrefix("a:"), WithPrefix("b:")).Handle("x")
Output: "a:b:x"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Middleware type** | `func(Handler) Handler` composes behaviour around a handler. |
| 2 | **Closures over state** | The counter lives outside the handler and is captured by the wrapper. |
| 3 | **Application order** | Wrapping in reverse makes the first listed middleware outermost. |

## Hint

To make the first middleware outermost, apply the list back to front.

## Validate

```bash
make verify
```
