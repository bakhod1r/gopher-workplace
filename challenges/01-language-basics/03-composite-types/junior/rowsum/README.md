# Row Sums

**Level:** junior
**Topic:** 01-language-basics → 03-composite-types

## Context

A spreadsheet totals each row. A 2-D grid is a slice of slices, possibly ragged.

## Task

Implement `RowSums(grid)` returning one sum per row.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  [][]int{{1,2,3},{4,5},{}}
Output: []int{6,9,0}
```

_Explanation:_ Empty row sums to 0.

**Example 2:**

```
Input:  nil
Output: []int{}
```

**Example 3:**

```
Input:  [][]int{{5}}
Output: []int{5}
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Slice of slices** | Rows may differ in length. |
| 2 | **Nested range** | Outer rows, inner elements. |
| 3 | **One output per row** | Append a sum per row. |

## Hint

`for _, row := range grid { s := 0; for _, v := range row { s += v }; out = append(out, s) }`.

## Validate

```bash
make verify
```
