# Shape Area

**Level:** junior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A CAD export totals the surface of every shape on a drawing.

## Task

Implement the stub(s) in [shapearea.go](shapearea.go):

1. Implement `Area` on `Rect` and `Square`.
2. Implement `TotalArea`, which sums the areas of every shape.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Rect{W: 3, H: 4}.Area()
Output: 12
```

**Example 2:**

```
Input:  Square{Side: 5}.Area()
Output: 25
```

**Example 3:**

```
Input:  TotalArea([]Shape{Rect{W: 2, H: 2}, Square{Side: 3}})
Output: 13
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Classic shape interface** | The canonical example of polymorphism in Go. |
| 2 | **Value receivers** | Shapes are immutable measurements. |
| 3 | **Float accumulation** | Reused from data types: summing `float64`. |

## Hint

`TotalArea` calls `s.Area()` — it never mentions `Rect` or `Square`.

## Validate

```bash
make verify
```
