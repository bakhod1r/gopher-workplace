# Merge Sorted

**Level:** middle
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

Two already-sorted feeds are merged into one stream without re-sorting everything.

## Task

Implement the stub(s) in [mergesort.go](mergesort.go):

1. Implement `Next` and `Peek` on `*SliceFeed` (`ok` is false once drained).
2. Implement `Merge`, which merges two feeds into one ascending slice.
3. Merge in one pass — do not concatenate and sort.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Merge([1 3], [2 4])
Output: [1 2 3 4]
```

**Example 2:**

```
Input:  Merge([], [1])
Output: [1]
```

**Example 3:**

```
Input:  Merge([1 1], [1])
Output: [1 1 1]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Feed interface** | `Peek`/`Next` is the minimal contract a merge needs. |
| 2 | **Two-pointer merge** | Efficiency: O(n+m) instead of O(n log n) after concatenation. |
| 3 | **Drained sources** | Reused: the `(value, ok)` idiom for streams. |

## Hint

Peek both sides, take the smaller, and drain the remainder when one side ends.

## Validate

```bash
make verify
```
