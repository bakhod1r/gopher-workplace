# Ring Buffer Index Underflow

**Level:** senior  
**Topic:** 03-generics

## Context

A crash reporter panics with an index-out-of-range while dumping its last log lines — but only when the buffer is partly filled.

## Task

Fix the single planted bug in [ringstartbug.go](ringstartbug.go):

1. Find and fix the single bug in `Items`.
2. Partly filled and fully wrapped buffers must both work.
3. `Add` and `NewRing` are provided and correct.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  cap 3, add 1
Output: [1]
```

**Example 2:**

```
Input:  cap 2, add 1,2,3
Output: [2 3]
```

**Example 3:**

```
Input:  cap 2, no adds
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Go's `%` keeps the sign** | `-1 % 3` is `-1`, so a bare modulus can produce a negative index. |
| 2 | **Normalising a modulus** | `((x % n) + n) % n` maps any integer into `[0, n)`. |
| 3 | **Partial fills expose it** | Once the ring has wrapped, `head - n` happens to be non-negative and hides the bug. |

## Hint

What is `head - n` before the ring has wrapped?

## Validate

```bash
make verify
```
