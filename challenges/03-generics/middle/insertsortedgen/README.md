# Insert Keeping Order

**Level:** middle  
**Topic:** 03-generics

## Context

A leaderboard keeps at most a few hundred entries sorted; re-sorting on every submission is wasteful.

## Task

Implement the stub(s) in [insertsortedgen.go](insertsortedgen.go):

1. Implement `InsertSorted`, returning a sorted copy with `v` inserted.
2. Place `v` after any equal elements.
3. The input is assumed sorted and must not be modified.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  InsertSorted([]int{1,3}, 2)
Output: []int{1,2,3}
```

**Example 2:**

```
Input:  InsertSorted([]int{1,2}, 2)
Output: []int{1,2,2}
```

**Example 3:**

```
Input:  InsertSorted([]int{}, 1)
Output: []int{1}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Upper-bound search** | Advancing past equal elements gives the last valid position. |
| 2 | **`slices.Insert`** | Handles the shifting once you know the index. |
| 3 | **Insertion beats re-sorting** | One insert is O(n); a sort is O(n log n) per submission. |

## Hint

Advance while `out[i] <= v` — the `=` is what puts `v` after its equals.

## Validate

```bash
make verify
```
