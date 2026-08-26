# Shape Area

**Level:** junior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

Different shapes compute area differently. An interface lets us sum areas
without knowing the concrete types.

## Task

Implement `Area` on `Circle` and `Rectangle` in [shapearea.go](shapearea.go):

1. `Circle.Area()` returns `math.Pi * r²`.
2. `Rectangle.Area()` returns `w * h`.

## Examples

**Example 1:**

```
Input:  Circle{5}.Area()
Output: 78.54...
```

**Example 2:**

```
Input:  Rectangle{3, 4}.Area()
Output: 12.0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Interface polymorphism** | `Shape` can be Circle or Rectangle. |
| 2 | **math.Pi** | Standard constant for π. |

## Validate

```bash
make verify
```
