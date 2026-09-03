# The Scratch Array That Went To The Heap

**Level:** senior
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A formatter declares a generously sized local array so it never has to grow. The allocation count looks fine and the process allocates four kilobytes for every number it prints.

## Task

Fix the single planted bug in [bigscratch.go](bigscratch.go):

1. Return `v`'s decimal digits, including the sign.
2. Fix the single bug so a call allocates on the order of the digits, not kilobytes.
3. Results from separate calls must be independent.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Format(42)
Output: "42"
```

**Example 2:**

```
Input:  Format(-9223372036854775808)
Output: the full minimum int64
```

_Explanation:_ Twenty bytes is enough for any int64.

**Example 3:**

```
Input:  2000 calls
Output: well under 128 bytes each
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Escape analysis follows the result** | The returned slice points into the scratch, so the scratch escapes. |
| 2 | **A local array is only free while it stays local** | Once it escapes, its full size is allocated. |
| 3 | **Size the buffer to the problem** | Twenty bytes covers every int64. |
| 4 | **Allocation count hides allocation size** | One 4 KiB allocation counts the same as one 20-byte one. |

## Hint

The allocation count is already 1. Look at how many bytes that one allocation is.

## Validate

```bash
make verify
```
