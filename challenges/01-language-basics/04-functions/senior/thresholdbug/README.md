# Boundary Comparison

**Level:** senior
**Topic:** 01-language-basics → 04-functions · _conditionals_

## Context

"Strictly greater" excludes the boundary. Using `>=` admits values equal to the
threshold — an off-by-one in the comparison operator, not the index.

## Task

Fix the comparison in [thresholdbug.go](thresholdbug.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  AboveThreshold([1 5 5 8 3], 5)
Output: [8]
```

**Example 2:**

```
Input:  AboveThreshold([1 2 3], 0)
Output: [1 2 3]
```

**Example 3:**

```
Input:  AboveThreshold([1], 5)
Output: []
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Strict vs inclusive comparison** | `>` excludes equality; `>=` includes it. |
| 2 | **Specification wording** | "greater than" means strict. |
| 3 | **Boundary values** | Equal-to-threshold must be dropped. |

## Hint

Use `v > t`, not `v >= t`.

## Validate

```bash
make verify
```
