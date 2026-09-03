# Emptying A Map Without Rebuilding It

**Level:** junior  
**Topic:** 11-performance-engineering

## Context

`m = make(map[K]V)` throws away every bucket the map had earned and makes the garbage collector deal with them — and it only rebinds *your* variable, leaving every other reference pointing at the old contents. `clear(m)` empties the map itself and keeps the buckets.

## Task

Implement both functions in [mapclearreuse.go](mapclearreuse.go):

1. `Reset` empties `m` in place so every reference to it sees an empty map; a nil map is left alone.
2. `Tally` empties `m`, counts the words into it, and returns it.
3. Reusing a large map must not re-allocate its buckets.

## Examples

**Example 1:**

```
Input:  Reset({a:1 b:2})
Output: the same map, now empty, seen through every reference
```

**Example 2:**

```
Input:  Tally({stale:7}, [a a b])
Output: {a:2 b:1}
```

**Example 3:**

```
Input:  Reset(nil)
Output: no panic, still nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`clear` on a map** | Deletes every key while keeping the allocated buckets. |
| 2 | **Reassignment is local** | `m = make(...)` inside a function changes nothing for the caller. |
| 3 | **Maps never shrink** | A map that once held a million keys keeps that bucket array — which is what makes reuse cheap and hoarding expensive. |

## Topics used again

Maps as reference types, the `clear` builtin, `range`.

## Hint

One builtin does the whole of `Reset`, and `clear(nil map)` is already a no-op.

## Validate

```bash
make verify
```
