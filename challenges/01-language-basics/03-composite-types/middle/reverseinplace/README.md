# Reverse In Place

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

Reversing without allocating: swap the ends inward. Because slices share their
backing array, this mutates the caller's data.

## Task

Implement `Reverse(xs)` in place (also return it).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  [1,2,3,4]
Output: [4,3,2,1]
```

**Example 2:**

```
Input:  [1,2,3]
Output: [3,2,1]
```

**Example 3:**

```
Input:  nil
Output: []
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Two-pointer swap** | `i` from front, `j` from back. |
| 2 | **In-place mutation** | Writes through the shared array. |
| 3 | **Stop at middle** | Loop while `i < j`. |

## Hint

`for i, j := 0, len(xs)-1; i < j; i, j = i+1, j-1 { xs[i], xs[j] = xs[j], xs[i] }`.

## Validate

```bash
make verify
```
