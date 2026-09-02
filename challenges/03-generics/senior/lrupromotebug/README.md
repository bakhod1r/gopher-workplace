# LRU That Is Really FIFO

**Level:** senior  
**Topic:** 03-generics

## Context

A read-through cache keeps evicting the hottest key. Its hit rate is far below what the size should give.

## Task

Fix the single planted bug in [lrupromotebug.go](lrupromotebug.go):

1. Find and fix the single bug so a successful read marks the entry as recently used.
2. Everything except `Get` is provided and correct.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  size 2: Put a,b; Get(a); Put(c)
Output: b evicted, a kept
```

**Example 2:**

```
Input:  Get(missing)
Output: zero, false
```

**Example 3:**

```
Input:  size 0
Output: nothing stored
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Recency versus insertion** | Without promotion on read the policy degrades to first-in-first-out. |
| 2 | **Silent degradation** | The cache still works — it just evicts the wrong entries. |
| 3 | **Where the policy lives** | One call is the entire difference between LRU and FIFO. |

## Hint

Compare what `Put` does after storing with what `Get` does after reading.

## Validate

```bash
make verify
```
