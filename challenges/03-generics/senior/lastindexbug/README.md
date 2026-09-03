# Last Index That Finds The First

**Level:** senior  
**Topic:** 03-generics

## Context

A log parser splits a line on the last separator to isolate the message. Lines containing more than one separator are being cut in the wrong place.

## Task

Fix the single planted bug in [lastindexbug.go](lastindexbug.go):

1. Find and fix the single bug so the *last* match is returned.
2. A missing value must still return -1.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  LastIndex([1,2,1], 1)
Output: 2
```

**Example 2:**

```
Input:  LastIndex([1,2,3], 2)
Output: 1
```

**Example 3:**

```
Input:  LastIndex([1], 9)
Output: -1
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Scan direction** | Returning on the first hit means the scan direction is the whole answer. |
| 2 | **Sentinel results** | -1 is the documented "not found", not a zero index. |

## Hint

Which end does the loop start from?

## Validate

```bash
make verify
```
