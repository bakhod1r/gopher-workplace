# Ordering A Report Without Wrecking The Data

**Level:** junior  
**Topic:** 11-performance-engineering

## Context

A profile viewer sorts the same rows several ways — by cum, by flat, by name — often while another part of the program still holds the original list. A sort that reorders the caller's slice in place turns "show me another view" into a data race waiting to happen.

## Task

Implement `SortByCum` in [profilesort.go](profilesort.go):

1. Order by `Cum` descending, then `Flat` descending, then `Func` ascending.
2. Return a new slice; leave the input untouched.
3. Empty input returns an empty, non-nil slice.

## Examples

**Example 1:**

```
Input:  [{a 1 5} {b 2 9} {c 0 7}]
Output: [{b 2 9} {c 0 7} {a 1 5}]
```

**Example 2:**

```
Input:  [{z 1 5} {a 1 5} {m 4 5}]
Output: [{m 4 5} {a 1 5} {z 1 5}]
```

**Example 3:**

```
Input:  nil
Output: [] (non-nil)
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Sorting mutates** | `slices.SortFunc` reorders the backing array, which the caller shares. |
| 2 | **Copy before sorting** | `slices.Clone` is the cheap way to give the caller their order back. |
| 3 | **Cascading comparison keys** | Fall through to the next key only when the previous compares equal. |

## Topics used again

`slices.Clone`, `slices.SortFunc`, `cmp.Compare`.

## Hint

Clone first, sort the clone.

## Validate

```bash
make verify
```
