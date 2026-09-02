# Transpose

**Level:** middle  
**Topic:** 03-generics

## Context

A CSV export needs the table pivoted, and malformed uploads with uneven rows must be rejected rather than crashing.

## Task

Implement the stub(s) in [transposegen.go](transposegen.go):

1. Implement `Transpose`, turning rows into columns.
2. Return an empty result when the input is empty or its rows differ in length.
3. Never index out of range.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Transpose([][]int{{1,2},{3,4}})
Output: [][]int{{1,3},{2,4}}
```

**Example 2:**

```
Input:  Transpose([][]int{{1,2},{3}})
Output: [][]int{}
```

**Example 3:**

```
Input:  Transpose([][]int{})
Output: [][]int{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Validating before working** | Checking rectangularity up front removes every bounds check later. |
| 2 | **Nested index order** | The output's `[c][r]` reads the input's `[r][c]`. |
| 3 | **Rows of zero width** | A row of length 0 makes the result empty, not a panic. |

## Hint

Validate rectangularity first; then the nested loops need no guards.

## Validate

```bash
make verify
```
