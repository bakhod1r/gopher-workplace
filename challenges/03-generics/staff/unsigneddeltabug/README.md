# Counter Deltas That Wrap Around

**Level:** staff  
**Topic:** 03-generics

## Context

A metrics agent turns cumulative counters into per-scrape rates. Every process restart produces one sample of roughly eighteen quintillion requests per second, and the dashboards are unreadable for the rest of the hour.

## Task

Fix the single planted bug in [unsigneddeltabug.go](unsigneddeltabug.go):

1. Find and fix the single bug so a counter reset is handled instead of subtracted.
2. Ordinary increases must keep their exact value.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Deltas([]uint8{10, 250, 5})
Output: []uint8{10, 240, 5}
```

**Example 2:**

```
Input:  Deltas([]uint64{5, 9})
Output: []uint64{5, 4}
```

**Example 3:**

```
Input:  Deltas([]uint32{})
Output: []uint32{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Unsigned has no negatives** | `a - b` for `b > a` wraps to a huge value rather than going below zero. |
| 2 | **Type sets** | A constraint names a *set* of types; the body must be correct for every member of it. |
| 3 | **Scale is a requirement** | A graded test asserts the result on millions of elements, so a defect that only shows past a threshold is caught. |

## Hint

What is `250 - 5` supposed to be, and what is `5 - 250`?

## Validate

```bash
make verify
```
