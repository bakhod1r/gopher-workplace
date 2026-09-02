# GC Sweep Phase

**Level:** staff
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

Once the mark phase has coloured everything reachable, the sweep walks the heap
linearly and reclaims whatever stayed unmarked. Here the heap is a slice of mark
bits and "reclaiming" means counting.

## Task

Implement `Sweep` on `*Heap` in [gcswp.go](gcswp.go):

1. Count how many entries in `h.Objects` are `false` (unmarked).
2. Return that count.

**Constraint (staff):** one linear pass, no allocations, over a 10M-object heap.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  [true, false, true]
Output: 1
```

**Example 2:**

```
Input:  [false, false]
Output: 2
```

**Example 3:**

```
Input:  [] (empty heap)
Output: 0
```

_Explanation:_ nothing allocated, nothing to free.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Linear heap scan** | Sweeping is O(heap), independent of how much is live. |
| 2 | **Counting the complement** | Unmarked means unreachable means free. |
| 3 | **`range` over a slice** | Value form is enough — nothing is mutated. |

## Hint

One accumulator, one loop, `if !marked { freed++ }`.

## Validate

```bash
make verify
```
