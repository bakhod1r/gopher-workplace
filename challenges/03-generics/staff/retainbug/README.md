# The Small Result That Pins A Huge Array

**Level:** staff  
**Topic:** 03-generics

## Context

An ingest job reads multi-megabyte payloads, keeps a four-element header from each, and drops the payload. Heap usage climbs with the number of payloads processed until the process is killed, even though only a handful of ints is retained per payload.

## Task

Fix the single planted bug in [retainbug.go](retainbug.go):

1. Find and fix the single bug so the result does not retain the input's storage.
2. The clamping behaviour and the returned contents must not change.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  h := Head(s, 2); s[0] = 99
Output: h[0] is still the original value
```

**Example 2:**

```
Input:  Head([]int{1,2}, 9)
Output: [1 2]
```

**Example 3:**

```
Input:  60 headers taken from 8 MB payloads
Output: heap stays small after the payloads are dropped
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Backing-array retention** | A tiny slice of a huge array keeps the whole allocation alive. |
| 2 | **Three-index slicing is not a copy** | `s[:n:n]` fixes the capacity but still points into the original array. |
| 3 | **The GC frees allocations, not ranges** | One live byte anywhere in an allocation keeps the whole allocation alive. |

## Hint

The capacity is already correct. What is the returned slice's pointer aimed at?

## Validate

```bash
make verify
```
