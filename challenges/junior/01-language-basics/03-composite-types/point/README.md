# Struct Point

**Level:** junior
**Topic:** 01-language-basics → 03-composite-types
**Estimated time:** 10 min

## Context

Geometry code models a 2-D coordinate as a `struct`. Like arrays, a struct is a
value: passing a `Point` copies its fields, so a translate helper returns a new
`Point` rather than mutating the caller's.

## Task

Implement `Translate` in [point.go](point.go) so it returns a new `Point`
shifted by `(dx, dy)`, without changing the caller's `p`.

Do **not** change the function signature, the `Point` type, or the tests.

## Examples

```go
Translate(Point{1, 2}, 3, 4)   // => Point{4, 6}
Translate(Point{0, 0}, -1, -1) // => Point{-1, -1}
Translate(Point{5, 5}, 0, 0)   // => Point{5, 5}
```

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
make verify   # fmt-check + vet + test
```

Green tests + clean `vet`/`gofmt` = challenge passed.
