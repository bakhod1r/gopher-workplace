# Set Intersection

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

Finding common elements (shared permissions, mutual friends).

## Task

Implement `Intersect(a, b)` — sorted, unique, present in both.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  a=[1,2,3,4], b=[2,4,6,2]
Output: [2,4]
```

_Explanation:_ in both, deduped+sorted

**Example 2:**

```
Input:  a=[1], b=[2]
Output: []
```

**Example 3:**

```
Input:  a=[3,3], b=[3]
Output: [3]
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Set from a** | Membership test for b. |
| 2 | **Dedup result** | A value in b twice counts once. |
| 3 | **Sort output** | Deterministic order. |

## Hint

Set of a; for each b, if in a-set and not yet emitted, add; then sort.

## Validate

```bash
make verify
```
