# Set Union

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

Combining two tag sets into one sorted, unique list.

## Task

Implement `Union(a, b)` — sorted, de-duplicated.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  a=[3,1,2], b=[2,4,1]
Output: [1,2,3,4]
```

**Example 2:**

```
Input:  nil,nil
Output: []
```

**Example 3:**

```
Input:  a=[1,1], b=[1]
Output: [1]
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Set membership** | `map[int]struct{}`. |
| 2 | **Collect + sort** | Keys to slice, then `sort.Ints`. |
| 3 | **Dedup** | Set removes duplicates automatically. |

## Hint

Insert all of a and b into a set, collect keys, `sort.Ints`.

## Validate

```bash
make verify
```
