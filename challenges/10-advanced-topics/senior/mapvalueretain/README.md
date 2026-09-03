# The Map Entry That Pinned The Whole Buffer

**Level:** senior
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

An indexer stores a few bytes per record out of a megabyte read buffer. The index is long-lived, the buffers are not, and resident memory grows by a megabyte per indexed record.

## Task

Fix the single planted bug in [mapvalueretain.go](mapvalueretain.go):

1. Store the `n` bytes of `batch` starting at `off` under `key`.
2. The stored value must own its bytes — the batch is reused and then dropped.
3. Ignore a nil map or an out-of-range range.
4. Fix the single bug.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Index(m, "a", batch, 0, 5)
Output: m["a"] is a 5-byte copy
```

**Example 2:**

```
Input:  the batch is overwritten afterwards
Output: the entry is unchanged
```

**Example 3:**

```
Input:  cap of the stored value
Output: the copy's size, not the batch's
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Allocation-granular collection** | One live view pins the whole array. |
| 2 | **Three-index slicing is not a copy** | It caps the capacity and still points at the batch. |
| 3 | **Lifetime mismatch** | A short-lived buffer stored in a long-lived map is the whole bug. |

## Hint

The capacity cap makes appends safe. What does it do about the pointer?

## Validate

```bash
make verify
```
