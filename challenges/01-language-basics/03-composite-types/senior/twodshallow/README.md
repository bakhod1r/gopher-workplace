# Deep-Copy a Grid

**Level:** senior
**Topic:** 01-language-basics → 03-composite-types

## Context

`copy(out, grid)` copies the **row slice headers**, so both grids share the same
rows. Writing a cell in the clone changes the original.

## Task

Fix the copy between the markers in [twodshallow.go](twodshallow.go) to clone each
row.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  [[1,2],[3,4]]
Output: independent copy; mutating out[0][0] doesn't affect input
```

**Example 2:**

```
Input:  [[9]]
Output: [[9]]
```

**Example 3:**

```
Input:  []
Output: []
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Shallow copy** | `copy` duplicates headers, not arrays. |
| 2 | **Per-row clone** | Copy each inner slice. |
| 3 | **Nested references** | Deep copy recurses one level. |

## Hint

`for i := range grid { out[i] = append([]int{}, grid[i]...) }`.

## Validate

```bash
make verify
```
