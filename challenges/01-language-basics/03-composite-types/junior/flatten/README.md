# Flatten a Grid

**Level:** junior
**Topic:** 01-language-basics → 03-composite-types

## Context

Turning a 2-D grid into a flat list in row-major order.

## Task

Implement `Flatten(grid)`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  [][]int{{1,2},{3},{},{4,5}}
Output: []int{1,2,3,4,5}
```

_Explanation:_ Empty rows add nothing.

**Example 2:**

```
Input:  nil
Output: []int{}
```

**Example 3:**

```
Input:  [][]int{{9}}
Output: []int{9}
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Row-major** | Iterate rows, then columns. |
| 2 | **append spread** | `append(out, row...)`. |
| 3 | **Ragged safe** | Handles rows of any length. |

## Hint

`for _, row := range grid { out = append(out, row...) }`.

## Validate

```bash
make verify
```
