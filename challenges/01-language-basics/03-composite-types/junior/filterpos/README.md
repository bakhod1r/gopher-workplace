# Filter Positives

**Level:** junior
**Topic:** 01-language-basics → 03-composite-types

## Context

Filtering builds a new slice with `append`, keeping only elements that pass a
test.

## Task

Implement `Positives(xs)` returning a new slice of the positive elements. Empty
result must be non-nil.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  []int{1,-2,3,0,4}
Output: []int{1,3,4}
```

_Explanation:_ 0 is not positive.

**Example 2:**

```
Input:  []int{-1,-2}
Output: []int{}
```

**Example 3:**

```
Input:  []int{}
Output: []int{}
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **append** | Grows a slice, returning the new header. |
| 2 | **Non-nil empty** | Start with `[]int{}` (or make) to avoid nil. |
| 3 | **Predicate** | Keep when `x > 0`. |

## Hint

`out := []int{}; for _, x := range xs { if x > 0 { out = append(out, x) } }`.

## Validate

```bash
make verify
```
