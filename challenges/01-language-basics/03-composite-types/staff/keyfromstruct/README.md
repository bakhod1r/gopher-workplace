# Struct Map Key Fields

**Level:** staff
**Topic:** 01-language-basics → 03-composite-types

## Context

A struct of comparable fields is a valid map key. But the key is built with `X`
and `Y` swapped (`Point{p.Y, p.X}`), so `(1,2)` and `(2,1)` collide.

## Task

Fix the key between the markers in [keyfromstruct.go](keyfromstruct.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  pts=[{1,2},{1,2},{3,4}]
Output: { {1,2}:2, {3,4}:1 }
```

**Example 2:**

```
Input:  pts=[{0,1}]
Output: { {0,1}:1 }
```

**Example 3:**

```
Input:  pts=[{2,3},{3,2}]
Output: { {2,3}:1, {3,2}:1 }
```

_Explanation:_ order-sensitive keys stay distinct.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Structs as keys** | Comparable structs are valid keys. |
| 2 | **Field identity** | Field values define the key. |
| 3 | **No slice fields** | A slice field makes it non-comparable. |

## Hint

`m[Point{p.X, p.Y}]++` (or `m[p]++`).

## Validate

```bash
make verify
```
