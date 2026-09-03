# Carve Aligned Blocks From One Allocation

**Level:** staff
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A parser allocates thousands of small nodes per document. They all die together when the document is finished, which is exactly the shape a bump allocator serves best.

## Task

Implement [arena.go](arena.go):

1. Return the next `n` bytes at an offset that is a multiple of `align`.
2. Report false for a negative size, an `align` that is not a power of two, or a request that does not fit.
3. A rejected request must consume nothing.
4. Each block's capacity must equal its length, so an append cannot reach the next block.

Replace the stub body in [arena.go](arena.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  NewArena(64).Alloc(8, 8)
Output: an 8-byte aligned block, true
```

**Example 2:**

```
Input:  Alloc(1,1) then Alloc(8,8)
Output: the second starts at offset 8
```

_Explanation:_ Padding is consumed.

**Example 3:**

```
Input:  Alloc(8, 3)
Output: nil, false
```

_Explanation:_ 3 is not a power of two.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Bump allocation** | One pointer moves forward; there is no per-object bookkeeping. |
| 2 | **Rounding up to an alignment** | `(x + a - 1) &^ (a - 1)` for a power-of-two `a`. |
| 3 | **Three-index slicing** | Caps each block so appends cannot spill. |
| 4 | **Bulk lifetime** | The whole arena is freed at once, or not at all. |

## Hint

Round the cursor up first, then check whether the block fits without overflowing.

## Validate

```bash
make verify
```
