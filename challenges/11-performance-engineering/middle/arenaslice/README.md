# One Block, Many Objects

**Level:** middle  
**Topic:** 11-performance-engineering

## Context

When a batch of objects all die at the same moment — everything allocated while handling one request — per-object allocation and per-object collection are pure overhead. A bump allocator carves them from one block by moving a cursor, and frees the entire batch by moving it back. That is the entire idea behind arenas, and behind `sync.Pool`-backed request scratch space.

## Task

Implement the five pieces in [arenaslice.go](arenaslice.go):

1. `NewArena(n)` allocates one block of `n` bytes; `Alloc(n)` carves `n` zeroed bytes from it and reports whether it fit.
2. A failed allocation must consume nothing, and a non-positive `n` gives an empty non-nil slice and `true`.
3. `Used` and `Free` report occupancy; `Reset` frees everything at once, and memory handed out after a reset must be zeroed.

## Examples

**Example 1:**

```
Input:  a := NewArena(64); a.Alloc(16)
Output: 16 bytes, true; Used 16, Free 48
```

**Example 2:**

```
Input:  NewArena(16).Alloc(32)
Output: false, Used still 0
```

**Example 3:**

```
Input:  fill an allocation with 0xFF, Reset, Alloc again
Output: zeroed bytes
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Bump allocation** | Allocation is a bounds check and an addition; that is all. |
| 2 | **Lifetime is collective** | Everything in the arena dies at `Reset`, so pointers into it must not outlive it. |
| 3 | **Reuse means re-zeroing** | Go's allocator hands out zeroed memory; an arena that does not is a source of very confusing bugs. |

## Topics used again

Slice windows over one array, `clear`, methods on pointer receivers, bounds checks.

## Hint

`a.block[a.used : a.used+n : a.used+n]` — the third index caps the capacity so a caller's `append` cannot scribble over the next allocation.

## Validate

```bash
make verify
```
