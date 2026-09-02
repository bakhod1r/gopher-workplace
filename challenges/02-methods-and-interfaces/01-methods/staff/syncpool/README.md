# Typed Buffer Pool

**Level:** staff
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

`sync.Pool` recycles allocations to take pressure off the garbage collector,
but its API is untyped: `Get` returns `any`. Wrapping it in a small struct with
two methods gives callers a type-safe pool and keeps the assertion in exactly
one place.

## Task

Implement `Get` and `Put` on `*BufferPool` in [syncpool.go](syncpool.go):

1. `Get()` takes a value from `p.pool` and asserts it to `*Buffer`.
2. `Put(b)` returns `b` to `p.pool`.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  NewPool(); Get()
Output: a *Buffer with 1024 bytes (from the pool's New func)
```

**Example 2:**

```
Input:  Get(); Put(b); Get()
Output: a usable *Buffer — very likely the same one
```

**Example 3:**

```
Input:  Get(); Get() with no Put in between
Output: two distinct buffers
```

_Explanation:_ the pool hands out an object at most once until it is returned.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Type assertion at the boundary** | `p.pool.Get().(*Buffer)` — the wrapper is what keeps `any` out of the caller's code. |
| 2 | **`New` guarantees non-nil** | With a `New` func set, `Get` never returns nil, so the assertion cannot fail. |
| 3 | **Pool contents are not guaranteed** | Entries may vanish at any GC; a pool is a cache, not a free list. |

## Hint

One line each. Use the single-value assertion — with a `New` function set, the
comma-ok form's failure branch is unreachable.

## Validate

```bash
make verify
```
