# One Allocation, Not One Per Element

**Level:** senior
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A graph loader allocates a node per vertex. For a million vertices that is a million small allocations, a million objects for the collector to scan, and a traversal that misses cache on every hop.

## Task

Fix the single planted bug in [ptrperitem.go](ptrperitem.go):

1. Return `n` nodes with IDs `0..n-1`, each addressed by one element of the result.
2. The nodes must be distinct and writable, including linking one to another.
3. Fix the single bug so the allocation count does not scale with `n`.
4. `n <= 0` returns nil.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Build(3)
Output: three nodes, IDs 0, 1, 2
```

**Example 2:**

```
Input:  256 nodes
Output: a handful of allocations, not 256
```

**Example 3:**

```
Input:  got[0].Next = got[1]
Output: links fine
```

_Explanation:_ The nodes are ordinary addressable values.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Block allocation** | One `[]Node` plus pointers into it replaces n separate objects. |
| 2 | **Elements of a slice are addressable** | `&block[i]` is a normal pointer, valid as long as the array lives. |
| 3 | **Locality** | Contiguous nodes make a traversal cache-friendly. |
| 4 | **Lifetime coupling** | Every node stays alive as long as any pointer into the block does. |

## Hint

The nodes do not have to be allocated one at a time to be pointed at one at a time.

## Validate

```bash
make verify
```
