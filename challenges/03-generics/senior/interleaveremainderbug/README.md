# Interleave That Drops The Tail

**Level:** senior  
**Topic:** 03-generics

## Context

A round-robin scheduler zips two queues together. Whenever the queues differ in length the surplus jobs vanish.

## Task

Fix the single planted bug in [interleaveremainderbug.go](interleaveremainderbug.go):

1. Find and fix the single bug so the longer slice's tail is kept.
2. The interleaved prefix must not change.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Interleave([1,2,3],[9])
Output: [1 9 2 3]
```

**Example 2:**

```
Input:  Interleave([1],[8,9])
Output: [1 8 9]
```

**Example 3:**

```
Input:  Interleave([],[])
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Loops bounded by the shorter input** | Everything past `min(len(a), len(b))` needs explicit handling. |
| 2 | **Both tails, one is empty** | Appending both remainders is safe: at most one is non-empty. |

## Hint

What happens after the shared prefix?

## Validate

```bash
make verify
```
