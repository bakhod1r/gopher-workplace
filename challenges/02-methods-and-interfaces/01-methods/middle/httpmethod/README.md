# HTTP Method

**Level:** middle
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

An HTTP server binds handlers from application methods. Instead of wrapping in
an anonymous function, pass the method value directly.

## Task

Implement `Handler` on `*App` in [httpmethod.go](httpmethod.go):

1. Return `a.HealthHandler` as an `http.HandlerFunc`.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  App{"myapp"}.Handler() → GET /health
Output: "OK: myapp"
```

**Example 2:**

```
Input:  App{"test"}.Handler() → GET /health
Output: "OK: test"
```

**Example 3:**

```
Input:  App{""}.Handler() → GET /health
Output: "OK: "
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Method values** | `a.HealthHandler` is a `func(http.ResponseWriter, *http.Request)`. |
| 2 | **http.HandlerFunc** | Type alias for handler signature — implicit conversion. |

## Hint

`return a.HealthHandler` — it's already the right signature.

## Validate

```bash
make verify
```
