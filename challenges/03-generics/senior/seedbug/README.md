# Minimum Seeded Wrong

**Level:** senior  
**Topic:** 03-generics

## Context

A latency floor gauge reports 0 ms for every window, even though no sample has ever been that fast.

## Task

Fix the single planted bug in [seedbug.go](seedbug.go):

1. Find and fix the single bug so the minimum comes from the slice.
2. All-positive and all-negative inputs must both work.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  MinOf([]int{4, 7})
Output: 4, true
```

**Example 2:**

```
Input:  MinOf([]int{-5, -1})
Output: -5, true
```

**Example 3:**

```
Input:  MinOf([]int{})
Output: 0, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Seeding an accumulator** | For an unknown type the only safe seed is an element, not the zero value. |
| 2 | **No universal maximum** | There is no `MaxValue[T]` you could seed with for an arbitrary ordered type. |
| 3 | **The empty case exists for this** | Because the seed must be an element, emptiness needs its own answer. |

## Hint

Where does `best` start?

## Validate

```bash
make verify
```
