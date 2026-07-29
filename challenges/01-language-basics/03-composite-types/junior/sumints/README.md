# Sum a Slice

**Level:** junior
**Topic:** 01-language-basics → 03-composite-types

## Context

The most basic slice operation: walk it and accumulate.

## Task

Implement `Sum(xs)` returning the total; empty/nil → 0.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  []int{1,2,3}
Output: 6
```

**Example 2:**

```
Input:  nil
Output: 0
```

**Example 3:**

```
Input:  []int{-5,5}
Output: 0
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **for range** | `for _, x := range xs` visits each element. |
| 2 | **Nil vs empty** | Both have length 0; ranging is safe. |
| 3 | **Accumulator** | Start at 0, add each element. |

## Hint

`total := 0; for _, x := range xs { total += x }`.

## Validate

```bash
make verify
```
