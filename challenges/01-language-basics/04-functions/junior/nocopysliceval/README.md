# Sum Without Mutating

**Level:** junior
**Topic:** 01-language-basics → 04-functions · _call-by-value_

## Context

A slice header is passed by value, but it points at a shared backing array —
so reading is safe, yet writing to elements would leak to the caller.

## Task

Implement `SumKeep` in [nocopysliceval.go](nocopysliceval.go): total the elements, changing none.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  SumKeep([1 2 3])
Output: 6
```

**Example 2:**

```
Input:  SumKeep(nil)
Output: 0
```

**Example 3:**

```
Input:  SumKeep([-2 2])
Output: 0
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Slice header copy** | The header is copied; the array is shared. |
| 2 | **Read-only traversal** | Ranging values never mutates. |
| 3 | **Named result** | `sum` is pre-declared as 0. |

## Hint

Range over `xs` adding each value to `sum`; return.

## Validate

```bash
make verify
```
