# Area

**Level:** junior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A floor-plan tool needs to compute the area of rectangular rooms. Each room is
stored as a `Rect` with `Width` and `Height`.

## Task

Implement the `Area` method on `Rect` in [area.go](area.go) so that:

1. It returns `Width × Height`.
2. Zero dimensions produce `0`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Rect{3, 4}.Area()
Output: 12
```

**Example 2:**

```
Input:  Rect{0, 5}.Area()
Output: 0
```

**Example 3:**

```
Input:  Rect{2.5, 4}.Area()
Output: 10
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Value receiver** | `(r Rect)` receives a copy — safe for read-only computation. |
| 2 | **Struct fields** | Access `r.Width` and `r.Height` directly. |
| 3 | **Methods vs functions** | `r.Area()` reads as "the area *of* this rect". |

## Hint

Multiply the two fields and return. A value receiver is correct here because
`Area` only reads the struct.

## Validate

```bash
make verify
```
