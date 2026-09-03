# The Split That Keeps Both Halves

**Level:** staff  
**Topic:** 03-generics

## Context

A bucketed ordered index keeps each bucket small by splitting it in two when it overflows. After a bulk load the index reports more keys than were inserted, and a scan emits several of them twice.

## Task

Fix the single planted bug in [bucketsplitcopybug.go](bucketsplitcopybug.go):

1. Find and fix the single bug so a split *moves* the upper half out of the original bucket.
2. The scan must stay in ascending key order.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Insert 1..20 with Max 4; Keys()
Output: [1 2 ... 20], length 20
```

**Example 2:**

```
Input:  Insert 1..20; Get(13)
Output: the value stored for 13
```

**Example 3:**

```
Input:  Insert 5 keys with Max 4
Output: two buckets, no key in both
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Structural invariants** | Every operation must restore what the type promises about itself. |
| 2 | **Move versus copy** | Splitting a container means partitioning it; leaving the source intact duplicates every moved element. |
| 3 | **Shifting a slice of slices** | Making room in the middle needs an append plus a `copy` of the tail, in that order. |

## Hint

Compare what the new bucket receives with what the old bucket is left holding.

## Validate

```bash
make verify
```
