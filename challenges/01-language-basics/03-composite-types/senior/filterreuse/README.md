# Filter Corrupts Input

**Level:** senior
**Topic:** 01-language-basics → 03-composite-types

## Context

`out := xs[:0]` reuses `xs`'s backing array. Appending the kept elements
overwrites `xs` from the front, so the input is corrupted mid-iteration.

## Task

Fix the initialization between the markers in
[filterreuse.go](filterreuse.go) to use a fresh slice.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  [1,2,3,4]
Output: [2,4]
```

_Explanation:_ input must stay [1,2,3,4].

**Example 2:**

```
Input:  [2,4,6]
Output: [2,4,6]
```

**Example 3:**

```
Input:  [1,3,5]
Output: []
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **xs[:0] aliases** | Same array, length 0. |
| 2 | **In-place filter** | Valid only when overwriting is intended. |
| 3 | **Fresh slice** | `[]int{}` or make for independence. |

## Hint

`out := []int{}` (don't reuse `xs`'s array).

## Validate

```bash
make verify
```
