# Perimeter

**Level:** junior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

The floor-plan tool also needs to compute the perimeter of a room to estimate
the total skirting-board length.

## Task

Implement the `Perimeter` method on `Rect` in [perimeter.go](perimeter.go):

1. Return `2 × (Width + Height)`.
2. Zero dimensions still produce a correct perimeter.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Rect{3, 4}.Perimeter()
Output: 14
```

**Example 2:**

```
Input:  Rect{0, 5}.Perimeter()
Output: 10
```

**Example 3:**

```
Input:  Rect{2.5, 3.5}.Perimeter()
Output: 12
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Value receiver** | Read-only access to struct fields via copy. |
| 2 | **Arithmetic on fields** | Combine `Width` and `Height` with standard operators. |

## Hint

`2 * (r.Width + r.Height)` — nothing more.

## Validate

```bash
make verify
```
