# Merge Maps

**Level:** junior
**Topic:** 01-language-basics → 03-composite-types

## Context

Merging config layers: a base map plus overrides, where later values win.

## Task

Implement `Merge(a, b)` — a new map, `b` overriding `a` on collisions, inputs
untouched.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  a={"x":1,"y":2}, b={"y":20,"z":3}
Output: {"x":1,"y":20,"z":3}
```

_Explanation:_ b overrides a on key y.

**Example 2:**

```
Input:  inputs after merge
Output: a still {"x":1,"y":2}
```

_Explanation:_ Fresh map, inputs untouched.

**Example 3:**

```
Input:  a={}, b={"k":9}
Output: {"k":9}
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **make(map)** | Create a fresh map to return. |
| 2 | **Range a map** | `for k, v := range m`. |
| 3 | **Override order** | Copy a first, then b. |

## Hint

`out := make(map[string]int)`; copy a, then copy b.

## Validate

```bash
make verify
```
