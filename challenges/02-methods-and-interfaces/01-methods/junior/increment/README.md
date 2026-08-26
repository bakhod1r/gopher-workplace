# Increment

**Level:** junior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A rate limiter tracks request counts. Each request increments the counter.

## Task

Implement `Inc` on `*Counter` in [increment.go](increment.go):

1. Add 1 to `N`.
2. Must use a pointer receiver so the caller sees the change.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  c := Counter{0}; c.Inc(); c.Value()
Output: 1
```

**Example 2:**

```
Input:  c := Counter{5}; c.Inc(); c.Value()
Output: 6
```

**Example 3:**

```
Input:  c := Counter{0}; c.Inc(); c.Inc(); c.Inc(); c.Value()
Output: 3
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Pointer receiver** | `*Counter` — mutation must persist. |
| 2 | **Value receiver** | `Value()` uses `Counter` (no `*`) — read-only. |
| 3 | **Methods vs functions** | `c.Inc()` vs `Inc(&c)`. |

## Hint

`c.N++` on a `*Counter` receiver.

## Validate

```bash
make verify
```
