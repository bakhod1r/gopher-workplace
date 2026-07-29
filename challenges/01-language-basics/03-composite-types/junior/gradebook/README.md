# Gradebook Averages

**Level:** junior
**Topic:** 01-language-basics → 03-composite-types

## Context

A map from student to their scores; produce each average. Students with no
scores are skipped.

## Task

Implement `Averages(book)` (integer average).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  {"ann":[90,80,100],"bob":[70,75],"cid":[]}
Output: {"ann":90,"bob":72}
```

_Explanation:_ cid has no scores and is omitted; 145/2=72 via integer division.

**Example 2:**

```
Input:  {"x":[10,20]}
Output: {"x":15}
```

**Example 3:**

```
Input:  {}
Output: {}
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Map of slices** | Value type is `[]int`. |
| 2 | **Guard empty** | Skip zero-length slices. |
| 3 | **Build a map** | Accumulate results. |

## Hint

Range the map; for non-empty slices, sum and divide, store in the result.

## Validate

```bash
make verify
```
