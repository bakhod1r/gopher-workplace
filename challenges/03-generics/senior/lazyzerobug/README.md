# The Lazy Value Recomputed Forever

**Level:** senior  
**Topic:** 03-generics

## Context

A config loader is supposed to hit the disk once. Profiling shows it reads the file on every request whenever the parsed value happens to be the type's zero.

## Task

Fix the single planted bug in [lazyzerobug.go](lazyzerobug.go):

1. Find and fix the single bug so `Fn` runs at most once.
2. A computed zero value must still count as computed.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  first Get()
Output: runs Fn once
```

**Example 2:**

```
Input:  second Get()
Output: Fn not run again
```

**Example 3:**

```
Input:  Fn returning 0
Output: still only one call
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Zero values are ambiguous** | `m[k]` yields a zero for a missing key, so presence needs the comma-ok form. |
| 2 | **A flag, not a sentinel** | Only a separate boolean can distinguish "unset" from "set to the zero value". |

## Hint

How does the code decide whether the value was computed?

## Validate

```bash
make verify
```
