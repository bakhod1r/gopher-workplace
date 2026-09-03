# Detect Whether Append Reallocated

**Level:** senior
**Topic:** 10-advanced-topics / 04-unsafe-package

## Context

A benchmark blames `append` for a copy nobody can see. Proving which appends actually reallocate needs a way to compare two slices' storage, which the language does not offer.

## Task

Implement [growcheck.go](growcheck.go):

1. Report whether `after` sits in different storage from `before`.
2. A slice with no capacity has no storage; treat that consistently.
3. Reslicing without appending must report false.

Replace the stub body in [growcheck.go](growcheck.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  s := make([]int,0,4); Grew(s, append(s,1))
Output: false
```

_Explanation:_ The capacity was enough.

**Example 2:**

```
Input:  s := make([]int,1,1); Grew(s, append(s,2))
Output: true
```

_Explanation:_ append reallocated.

**Example 3:**

```
Input:  Grew(s, s[:1])
Output: false
```

_Explanation:_ Reslicing keeps the array.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **append's growth policy** | It reallocates only when the capacity is exhausted. |
| 2 | **SliceData identity** | The data pointer is the only observable identity of the backing array. |
| 3 | **Zero-capacity slices** | There is no array to point at, so the pointer comparison is not meaningful. |

## Hint

Compare the data pointers — after deciding what a zero-capacity slice means.

## Validate

```bash
make verify
```
