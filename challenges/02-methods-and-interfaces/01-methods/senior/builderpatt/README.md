# Builder Pattern

**Level:** senior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

Constructing a `Request` with many optional fields via one giant constructor is
unreadable. A builder gives each field its own method and returns the builder
itself, so calls chain into a sentence: `NewBuilder().Method("GET").URL("/api")`.

## Task

Implement `URL` and `Auth` on `*RequestBuilder` in [builderpatt.go](builderpatt.go):

1. `URL(u)` sets `b.req.URL` and returns `b`.
2. `Auth(t)` sets `b.req.Auth` and returns `b`.
3. `Method` already shows the shape — follow it.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  NewBuilder().URL("/api").Build()
Output: Request{URL: "/api"}
```

**Example 2:**

```
Input:  NewBuilder().Method("GET").URL("/api").Auth("token123").Build()
Output: Request{Method: "GET", URL: "/api", Auth: "token123"}
```

**Example 3:**

```
Input:  NewBuilder().Build()
Output: Request{}
```

_Explanation:_ unset fields keep their zero values.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Returning the receiver** | `return b` is what makes the next `.Method(...)` legal. |
| 2 | **Pointer receiver** | Chaining on a value receiver would mutate and discard copies. |
| 3 | **Accumulator field** | `b.req` is built up in place, then handed out by `Build`. |

## Hint

Two lines each: assign the field, `return b`. If a method returns nothing, the
chain stops compiling at the next dot.

## Validate

```bash
make verify
```
