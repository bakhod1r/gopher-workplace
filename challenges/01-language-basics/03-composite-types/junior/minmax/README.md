# Min and Max

**Level:** junior
**Topic:** 01-language-basics → 03-composite-types

## Context

A single pass finds both extremes — seed from the first element, then compare.

## Task

Implement `MinMax(xs)` returning min, max, and `ok=false` for empty.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  []int{3,1,4,1,5}
Output: 1, 5, true
```

**Example 2:**

```
Input:  []int{7}
Output: 7, 7, true
```

**Example 3:**

```
Input:  nil
Output: 0, 0, false
```

_Explanation:_ Empty input -> ok is false.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Seed from first** | Initialize both to `xs[0]`. |
| 2 | **Single pass** | Update min/max together. |
| 3 | **Empty guard** | Return ok=false, avoid indexing xs[0]. |

## Hint

Guard empty; set `min=max=xs[0]`; loop the rest updating both.

## Validate

```bash
make verify
```
