# Independent Slice Copy

**Level:** junior
**Topic:** 01-language-basics → 03-composite-types

## Context

Assigning a slice shares its backing array. A real copy needs `make` + `copy`.

## Task

Implement `Clone(xs)` — an independent copy; nil → non-nil empty.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  xs=[]int{1,2,3}
Output: []int{1,2,3} (independent copy)
```

**Example 2:**

```
Input:  clone then c[0]=99
Output: xs[0] still 1
```

_Explanation:_ No shared backing array.

**Example 3:**

```
Input:  xs=nil
Output: []int{} (non-nil, len 0)
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Shared backing** | `b := a` shares memory. |
| 2 | **make + copy** | Allocate then `copy(dst, src)`. |
| 3 | **Non-nil empty** | Return a length-0 non-nil slice. |

## Hint

`out := make([]int, len(xs)); copy(out, xs); return out`.

## Validate

```bash
make verify
```
