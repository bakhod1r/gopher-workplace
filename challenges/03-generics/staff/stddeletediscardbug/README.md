# The Deletion That Was Never Kept

**Level:** staff  
**Topic:** 03-generics

## Context

A retention job trims expired records from a batch. The batch it hands on is always full length, and the tail is a smear of records that were supposedly deleted — some of them duplicated.

## Task

Fix the single planted bug in [stddeletediscardbug.go](stddeletediscardbug.go):

1. Find and fix the single bug so the returned slice ends after the kept items.
2. The kept items must stay in their original order.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Purge([1,2,3,4], even)
Output: [1 3]
```

**Example 2:**

```
Input:  Purge([2,4], even)
Output: []
```

**Example 3:**

```
Input:  Purge([1,3], even)
Output: [1 3]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Helpers return, they do not mutate the header** | A function that reslices must hand the new header back, and you must keep it. |
| 2 | **Compaction leaves debris** | Everything past the new length is whatever the shift left behind, not padding. |

## Hint

`kept` is computed, consulted once, and then abandoned. Which slice is actually returned?

## Validate

```bash
make verify
```
