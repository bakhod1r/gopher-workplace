# Event Bus

**Level:** middle  
**Topic:** 03-generics

## Context

A UI dispatches typed events to whoever is listening, and listeners come and go as views are opened and closed.

## Task

Implement the stub(s) in [eventbusgen.go](eventbusgen.go):

1. Implement `Subscribe`, `Unsubscribe`, and `Publish`.
2. Ids must be unique and must not be reused after unsubscribing.
3. `Publish` returns the number of handlers it called; delivery order is unspecified.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Subscribe(h); Publish(1)
Output: 1 handler called
```

**Example 2:**

```
Input:  Unsubscribe(id); Publish(1)
Output: 0 handlers
```

**Example 3:**

```
Input:  Unsubscribe(unknown)
Output: false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Typed events** | `Bus[T]` fixes the payload type at compile time — no `any` and no assertions. |
| 2 | **Lazy allocation** | Allocating the map on first subscribe keeps the zero value usable. |
| 3 | **Ids from a counter** | A monotonically increasing counter never reuses an id, unlike a slice index. |

## Hint

Increment the counter before storing, so ids start at 1 and never repeat.

## Validate

```bash
make verify
```
