# Concatenate Slices

**Level:** junior
**Topic:** 01-language-basics → 03-composite-types

## Context

Flattening several slices into one with a variadic parameter and append-spread.

## Task

Implement `Concat(slices ...[]int)`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  []int{1,2}, []int{3}, nil, []int{4,5}
Output: []int{1,2,3,4,5}
```

_Explanation:_ nil slice contributes nothing when spread with append.

**Example 2:**

```
Input:  (no args)
Output: []int{}
```

_Explanation:_ empty variadic yields empty result.

**Example 3:**

```
Input:  []int{7}, []int{8,9}
Output: []int{7,8,9}
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Variadic of slices** | `...[]int` is a slice of slices. |
| 2 | **append spread** | `append(out, s...)`. |
| 3 | **nil handling** | Appending nil adds nothing. |

## Hint

`for _, s := range slices { out = append(out, s...) }`.

## Validate

```bash
make verify
```
