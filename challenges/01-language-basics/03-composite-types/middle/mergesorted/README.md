# Merge Sorted Slices

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

The merge step of merge sort: combine two sorted runs into one, in O(n).

## Task

Implement `Merge(a, b)` (ascending, duplicates kept).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  [1,3,5],[2,4,6]
Output: [1,2,3,4,5,6]
```

**Example 2:**

```
Input:  [1,1,2],[1,3]
Output: [1,1,1,2,3]
```

_Explanation:_ duplicates kept

**Example 3:**

```
Input:  nil,nil
Output: []
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Two pointers** | Advance the smaller side. |
| 2 | **Drain remainder** | Append the leftover tail. |
| 3 | **Stable dupes** | Keep equal elements. |

## Hint

Walk `i,j`; append the smaller; then append whichever slice remains.

## Validate

```bash
make verify
```
