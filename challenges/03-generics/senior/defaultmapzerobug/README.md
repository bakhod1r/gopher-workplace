# The Lookup That Grows The Map

**Level:** senior  
**Topic:** 03-generics

## Context

A feature-flag store with a default reports thousands of configured flags after a day of traffic, though nobody set any.

## Task

Fix the single planted bug in [defaultmapzerobug.go](defaultmapzerobug.go):

1. Find and fix the single bug so reading an absent key does not insert it.
2. The default must still be returned.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Get(missing)
Output: the default
```

**Example 2:**

```
Input:  Get(missing); Len()
Output: 0
```

**Example 3:**

```
Input:  Set(k,1); Get(k)
Output: 1
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Reads should not write** | An inserting read turns a lookup into unbounded memory growth. |
| 2 | **Zero values are ambiguous** | `m[k]` yields a zero for a missing key, so presence needs the comma-ok form. |

## Hint

Count the keys after a read.

## Validate

```bash
make verify
```
