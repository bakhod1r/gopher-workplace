# Struct Point

**Level:** junior
**Topic:** 01-language-basics → 03-composite-types

## Context

Geometry code models a 2-D coordinate as a `struct`. Like arrays, a struct is a
value: passing a `Point` copies its fields, so a translate helper returns a new
`Point` rather than mutating the caller's.

## Task

Implement `Translate` in [point.go](point.go) so it returns a new `Point`
shifted by `(dx, dy)`, without changing the caller's `p`.

Do **not** change the function signature, the `Point` type, or the tests.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  p=Point{1,2}, dx=3, dy=4
Output: Point{4,6}
```

**Example 2:**

```
Input:  p=Point{0,0}, dx=-1, dy=-1
Output: Point{-1,-1}
```

**Example 3:**

```
Input:  p=Point{5,5}, dx=0, dy=0
Output: Point{5,5}
```

_Explanation:_ Caller's p not mutated — a fresh Point is returned.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Struct types** | A `struct` groups named fields; the type is defined once and reused. |
| 2 | **Struct literals** | Build a value with `Point{X: a, Y: b}` (or positional `Point{a, b}`). |
| 3 | **Value semantics** | A struct argument is copied; mutating the parameter never affects the caller. |
| 4 | **Field access** | Read fields with `p.X`, `p.Y`. |

## Hint

Return a fresh literal built from the shifted fields:
`Point{X: p.X + dx, Y: p.Y + dy}`. Keep `dx` for X and `dy` for Y.

## Validate

```bash
make verify
```
