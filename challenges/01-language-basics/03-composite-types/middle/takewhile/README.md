# Take While Positive

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

Reading a leading run until a condition breaks — like parsing a prefix.

## Task

Implement `TakePositive(xs)` returning the longest all-positive prefix.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  [1,2,3,-1,4]
Output: [1,2,3]
```

_Explanation:_ stops at first <=0

**Example 2:**

```
Input:  [-1,2]
Output: []
```

**Example 3:**

```
Input:  [1,2,3]
Output: [1,2,3]
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Prefix scan** | Stop at first failure. |
| 2 | **Slice up to i** | `xs[:i]` is the taken prefix. |
| 3 | **Empty result** | Return non-nil empty. |

## Hint

Find the first index where `xs[i] <= 0`, return `xs[:i]` (or all).

## Validate

```bash
make verify
```
