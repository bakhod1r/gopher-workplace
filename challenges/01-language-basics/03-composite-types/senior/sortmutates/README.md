# Sort Mutates Caller

**Level:** senior
**Topic:** 01-language-basics → 03-composite-types

## Context

`out := xs` copies the slice **header**, not the data, so `sort.Ints(out)` sorts
the caller's backing array. "SortedCopy" corrupts its input.

## Task

Fix the copy between the markers in [sortmutates.go](sortmutates.go) to duplicate
the data first.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  [3,1,2]
Output: [1,2,3]; input stays [3,1,2]
```

**Example 2:**

```
Input:  [1]
Output: [1]
```

**Example 3:**

```
Input:  [5,4,3,2,1]
Output: [1,2,3,4,5]
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Header vs data** | `out := xs` shares the array. |
| 2 | **sort in place** | `sort.Ints` mutates. |
| 3 | **Copy first** | `append([]int{}, xs...)`. |

## Hint

`out := append([]int{}, xs...)`.

## Validate

```bash
make verify
```
