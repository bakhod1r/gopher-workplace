# NaN Poisons The Minimum

**Level:** senior  
**Topic:** 03-generics

## Context

A sensor dashboard shows NaN for the minimum whenever the very first reading of a window was dropped.

## Task

Fix the single planted bug in [nanbug.go](nanbug.go):

1. Find and fix the single bug so a leading NaN cannot poison the result.
2. An all-NaN or empty slice must report `false`.
3. Do not import `math` — NaN detection must stay inline.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  MinIgnoringNaN([NaN, 2])
Output: 2, true
```

**Example 2:**

```
Input:  MinIgnoringNaN([3, NaN, 1])
Output: 1, true
```

**Example 3:**

```
Input:  MinIgnoringNaN([NaN])
Output: 0, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **NaN loses every comparison** | `v < best` is false when `best` is NaN, so the seed is never replaced. |
| 2 | **The `v != v` idiom** | Only NaN is unequal to itself. |
| 3 | **Seeding needs a flag** | When any element may be invalid, `s[0]` is not a safe seed. |

## Hint

The loop skips NaN. What about the seed?

## Validate

```bash
make verify
```
