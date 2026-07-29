# Sorted Map Keys

**Level:** junior
**Topic:** 01-language-basics → 03-composite-types

## Context

Map iteration order is random, so to print or compare deterministically you
collect the keys and sort them.

## Task

Implement `Sorted(m)` returning keys in ascending order.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  {"banana":1,"apple":2,"cherry":3}
Output: ["apple","banana","cherry"]
```

_Explanation:_ Sorted ascending regardless of map order.

**Example 2:**

```
Input:  {}
Output: [] (empty)
```

**Example 3:**

```
Input:  {"b":1,"a":2}
Output: ["a","b"]
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Collect keys** | Range the map, append `k`. |
| 2 | **sort.Strings** | Sorts a []string in place. |
| 3 | **Random order** | Map ranging is intentionally unordered. |

## Hint

Append each key, then `sort.Strings(out)`.

## Validate

```bash
make verify
```
