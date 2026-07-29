# Clamp the Sub-slice End

**Level:** senior
**Topic:** 01-language-basics → 03-composite-types

## Context

`xs[:n]` panics when `n > len(xs)`. A "take up to n" helper must clamp the end to
the length.

## Task

Fix the return between the markers in
[subsliceclamp.go](subsliceclamp.go) to clamp `n`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  xs=[1,2,3], n=2
Output: [1,2]
```

**Example 2:**

```
Input:  xs=[1,2,3], n=10
Output: [1,2,3]
```

**Example 3:**

```
Input:  xs=[1,2,3], n=0
Output: []
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Slice bounds** | High index must be ≤ len. |
| 2 | **Clamp** | `min(n, len(xs))`. |
| 3 | **Panic safety** | Out-of-range slicing panics. |

## Hint

`if n > len(xs) { n = len(xs) }; return xs[:n]`.

## Validate

```bash
make verify
```
