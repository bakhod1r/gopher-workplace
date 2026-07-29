# Dedupe Consecutive

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

Collapsing runs of equal values (e.g. deduping a sorted list) in one pass.

## Task

Implement `Dedupe(xs)` collapsing consecutive duplicates.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  [1,1,2,3,3,3,4]
Output: [1,2,3,4]
```

**Example 2:**

```
Input:  [5,5,5]
Output: [5]
```

**Example 3:**

```
Input:  []
Output: []
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Track previous** | Keep the last emitted value. |
| 2 | **Compare & skip** | Skip when equal to previous. |
| 3 | **First element** | Always keep the first. |

## Hint

Keep the first; append `xs[i]` only when it differs from `xs[i-1]`.

## Validate

```bash
make verify
```
