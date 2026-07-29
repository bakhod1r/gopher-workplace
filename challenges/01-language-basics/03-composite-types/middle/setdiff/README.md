# Set Difference

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

"In A but not B" — revoked items, missing keys.

## Task

Implement `Diff(a, b)` — sorted, unique, in a not b.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  a=[1,2,3,3,4], b=[2,4]
Output: [1,3]
```

_Explanation:_ in a not in b, deduped+sorted

**Example 2:**

```
Input:  a=[1], b=[1]
Output: []
```

**Example 3:**

```
Input:  a=[5,5], b=[]
Output: [5]
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Exclusion set** | Membership test against b. |
| 2 | **Dedup** | Emit each survivor once. |
| 3 | **Sort** | Deterministic order. |

## Hint

Set of b; for each a not in b and not already emitted, add; then sort.

## Validate

```bash
make verify
```
