# Distance

**Level:** junior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A mapping application needs to compute straight-line distances between two
points on a 2D plane.

## Task

Implement `DistanceTo` on `Point` in [distance.go](distance.go):

1. Return the Euclidean distance: `√((x₂−x₁)² + (y₂−y₁)²)`.
2. Same point → distance is `0`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Point{0,0}.DistanceTo(Point{3,4})
Output: 5
```

**Example 2:**

```
Input:  Point{1,1}.DistanceTo(Point{1,1})
Output: 0
```

**Example 3:**

```
Input:  Point{0,0}.DistanceTo(Point{7,0})
Output: 7
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Value receiver** | Both points are read-only. |
| 2 | **Method with parameters** | Methods can take additional arguments beyond the receiver. |
| 3 | **math package** | `math.Sqrt` for square roots. |

## Hint

`dx := other.X - p.X`, `dy := other.Y - p.Y`, then `math.Sqrt(dx*dx + dy*dy)`.

## Validate

```bash
make verify
```
